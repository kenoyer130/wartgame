package weaponabilities

import "github.com/kenoyer130/wartgame/models"

func ApplyWeaponAbilityShot(weapon models.Weapon) int {
		
	for i := 0; i < len(weapon.Abilities); i++ {
		// blast requires special rules
		if weapon.Abilities[i] == "Blast" {
			return ApplyBlast(weapon)
		}
	}

	return weapon.WeaponType.Number
}

func GetWeaponAbility() {
	
}
