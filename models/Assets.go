package models

type Assets struct {
	Squads  map[string]Squad
	Weapons map[string]Weapon
}

func NewAssets() *Assets {
	var a Assets
	a.Squads = make(map[string]Squad)
	a.Weapons = make(map[string]Weapon)
	return &a
}
