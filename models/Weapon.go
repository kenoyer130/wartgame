package models

import (
	"math/rand"
	"strconv"
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

func (re WeaponType) GetDice(weapon Weapon) int {
	if re.Dice == "3D3" {
		return re.getRndDice()
	}

	die, _ := strconv.Atoi(re.Dice)

	for _, ability := range weapon.Abilities {

		if ability == "Blast" {
			return weapon.ApplyBlast(die)
		}
	}

	return die
}

func (re WeaponType) getRndDice() int {
	dmg := 0

	for i := 0; i < 3; i++ {
		dmg = dmg + rand.Intn(3) + 1
	}

	return dmg
}
