package models

type UnitType string

const (
	Leader UnitType = "Leader"
)

type Unit struct {
	Name           string
	UnitNumber     UnitNumber
	Movement       int
	WeaponSkill    string
	BallisticSkill string
	Strength       int
	Toughness      int
	Wounds         int
	Attackes       int
	Leadership     int
	Save           string
	Weapons        []Weapon
	Location       Location
	UnitType       UnitType
}

type UnitNumber struct {
	Min int
	Max int
}

func (re Unit) GetLocation() Location {
	return re.Location
}
