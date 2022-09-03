package models

import "github.com/kenoyer130/wartgame/interfaces"

type ShootingWeapon struct {
	Unit        Unit
	Model       Model
	Weapon      Weapon
	Count       int
	Targets     []interfaces.Entity
	TargetRange int
}
