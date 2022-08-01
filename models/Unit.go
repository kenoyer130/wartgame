package models

type Unit struct {
	Name           string
	Army           string
	Models         []Model
	DefaultWeapons []string
	Power          int
	Location       Location
	UnitState      UnitState
}

type UnitState struct {
	Advanced bool
	FellBack bool
	Shot     bool
}
