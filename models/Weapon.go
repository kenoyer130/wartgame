package models

type Weapon struct {
	Name          string
	Range         int
	WeaponType    WeaponType
	Strength      int
	ArmorPiercing int
	Damage        int
	Abilities     []string
}

type WeaponType struct {
	Type   string
	Dice   int
	Number int
}
