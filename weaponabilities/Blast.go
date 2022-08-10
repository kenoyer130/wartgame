package weaponabilities

import (
	"fmt"
	"math/rand"

	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func ApplyBlast(weapon models.Weapon) int {

	shot := 1

	targetCount := len(models.Game().SelectedTargetUnit.Models)

	if targetCount > 11 {
		shot = weapon.WeaponType.Number
	}

	dice := weapon.WeaponType.Dice

	shot = rand.Intn(dice) + 1

	if targetCount > 6 && shot < 3 {
		shot = 3
	}

	engine.WriteMessage(fmt.Sprintf("Blast weapon hits %d targets!", shot))

	return shot
}
