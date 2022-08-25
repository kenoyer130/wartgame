package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	Game().BattleGround = *NewBattleGround(50, 50)

	unit := Unit{}
	unit.ID = "5"
	unit.Location = Location{X: 10, Y: 10}

	target := Unit{}
	target.ID = "4"
	target.Location = Location{X: 5, Y: 5}
	Game().BattleGround.PlaceBattleGroundEntity(&target)

	outOfRange := Unit{}
	outOfRange.ID = "3"
	outOfRange.Location = Location{X: 25, Y: 25}

	Game().BattleGround.PlaceBattleGroundEntity(&outOfRange)

	m.Run()
}

func TestShootingPhase(t *testing.T) {
	entites := InRange("5", 6, 10, 10)
	assert.Equal(t, len(entites), 1)
}
