package weaponabilities

import "github.com/kenoyer130/wartgame/models"

func ApplyWeaponAbilityShot(weapon models.Weapon) int {

	shot := 1

	for i := 0; i < len(weapon.Abilities); i++ {
		if weapon.Abilities[i] == "Blast" {
			return ApplyBlast(weapon)
		}
	}

	return shot
}
