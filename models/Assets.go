package models

type Assets struct {
	Units   map[string]Unit
	Weapons map[string]Weapon
}

func NewAssets() *Assets {
	var a Assets
	a.Units = make(map[string]Unit)
	a.Weapons = make(map[string]Weapon)
	return &a
}
