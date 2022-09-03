package models

import (
	"math/rand"

	"github.com/kenoyer130/wartgame/interfaces"
)

type Weapon struct {
	Name          string
	Range         int
	WeaponType    WeaponType
	Strength      int
	ArmorPiercing int
	Damage        int
	Abilities     []string
	Fired         bool
}

func (re *Weapon) SetArmorPiercing(value int) {
	re.ArmorPiercing = value
}

func (re *Weapon) GetWeaponAbilities() []string {
	return re.Abilities
}

type WeaponType struct {
	Type   string
	Dice   string
	Number int
}

func (re WeaponType) GetAttacks(attacks int, target interfaces.Entity, shootingWeapon ShootingWeapon) int {
	if shootingWeapon.Weapon.WeaponType.Type == "RF" {
		distance := shootingWeapon.Unit.Location.Subtract(target.GetLocation())
		rapidFireRange := shootingWeapon.Weapon.Range / 2

		if (distance.X <= rapidFireRange) && (distance.Y <= rapidFireRange) {
			return attacks * 2
		}

		return attacks
	}

	if shootingWeapon.Weapon.WeaponType.Type == "As" {
		return attacks * shootingWeapon.Weapon.WeaponType.Number
	}	

	if re.Dice == "3D3" {
		return re.getRndDice()
	}

	for _, ability := range shootingWeapon.Weapon.Abilities {

		if ability == "Blast" {
			return shootingWeapon.Weapon.ApplyBlast(attacks)
		}
	}

	if shootingWeapon.Weapon.WeaponType.Number > 0 {
		attacks = attacks * shootingWeapon.Weapon.WeaponType.Number
	}

	return attacks
}

func (re WeaponType) getRndDice() int {
	dmg := 0

	for i := 0; i < 3; i++ {
		dmg = dmg + rand.Intn(3) + 1
	}

	return dmg
}
