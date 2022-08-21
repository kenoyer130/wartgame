package models

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

func (re Weapon) SetArmorPiercing(value int) {
	re.ArmorPiercing = value
}

type WeaponType struct {
	Type   string
	Dice   int
	Number int
}
