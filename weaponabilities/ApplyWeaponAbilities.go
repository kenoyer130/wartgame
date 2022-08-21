package weaponabilities

import (
	"github.com/kenoyer130/wartgame/interfaces"
)

type WeaponAbilityList struct {
	weaponAbilities map[string]interfaces.WeaponAbility
}

func (re *WeaponAbilityList) Init() {
	re.weaponAbilities = make(map[string]interfaces.WeaponAbility)

	re.weaponAbilities["Shuriken"] = Shuriken {}
}

func (re WeaponAbilityList) ApplyWeaponAbilities(phase interfaces.WeaponAbilityPhase, die int, weapon interfaces.IWeaponAbility) interfaces.IWeaponAbility {
	for _, ability := range re.weaponAbilities {
		if ability.GetPhase() == phase {
			return ability.Apply(die, weapon)
		}
	}

	return weapon
}
