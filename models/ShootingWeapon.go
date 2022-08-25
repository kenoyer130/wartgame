package models

type ShootingWeapon struct {
	Model   Model
	Weapon  Weapon
	Count   int
	Targets []Entity
}
