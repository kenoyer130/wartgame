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

func (re WeaponType) GetDice() int {
	if re.Dice == "3D3" {
		return getRndDice()
	}

	dmg, _ := strconv.Atoi(re.Dice)

	return dmg
}

func getRndDice() int {
	dmg := 0

	for i := 0; i < 3; i++ {
		dmg = dmg + rand.Intn(3)
	}

	return dmg
}
