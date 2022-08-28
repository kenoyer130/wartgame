package models

import (
	"math/rand"
)

func (re Weapon) ApplyBlast(max int) int {

	shot := 1

	targetCount := len(Game().SelectedTargetUnit.Models)

	dice := max

	if targetCount > 11 {
		shot = max
	} else {
		shot = rand.Intn(dice) + 1
	}

	if targetCount > 6 && shot < 3 {
		shot = 3
	}

	return shot
}
