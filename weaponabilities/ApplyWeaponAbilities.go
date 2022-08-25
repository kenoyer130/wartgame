package weaponabilities

import (
	"github.com/kenoyer130/wartgame/interfaces"
)

type WeaponAbilityList struct {
	weaponAbilities map[string]interfaces.WeaponAbility
}

func (re *WeaponAbilityList) Init() {
	re.weaponAbilities = make(map[string]interfaces.WeaponAbility)

	re.weaponAbilities["Shuriken"] = Shuriken{}
}

func (re WeaponAbilityList) ApplyWeaponAbilities(phase interfaces.WeaponAbilityPhase, die int, weapon interfaces.IWeaponAbility) {

	for _, ability := range weapon.GetWeaponAbilities() {

		if re.weaponAbilities[ability] == nil {
			continue
		}

		weaponAbility := re.weaponAbilities[ability]
		if weaponAbility.GetPhase() == phase {
			weaponAbility.Apply(die, weapon)
		}
	}
}
