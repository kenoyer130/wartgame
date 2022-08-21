package weaponabilities

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/interfaces"
)

type Shuriken struct {
}

func (re Shuriken) GetType() string {
	return "Shuriken"
}

func (re Shuriken) GetPhase() interfaces.WeaponAbilityPhase {
	return interfaces.WeaponAbilityPhaseWounds
}

func  (re Shuriken) Apply(die int, weapon interfaces.IWeaponAbility) interfaces.IWeaponAbility {
	if die == 6 {
		weapon.SetArmorPiercing(-2)
		engine.WriteMessage("shuriken: die of 6 raised AP to 2")
	}

	return weapon
}
