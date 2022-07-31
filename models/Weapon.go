package models

type Weapon struct {
	Name          string
	Range         int
	Type          WeaponType
	Strength      int
	ArmorPiercing int
	Damage        int
}

type WeaponType struct {
	Type   string
	Dice   int
	Number int
}

type WeaponAbility struct {
	Name string
}
