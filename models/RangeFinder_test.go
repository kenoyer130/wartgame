package models

import (
	"testing"

	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	Game().BattleGround = *NewBattleGround(50, 50)

	unit := Unit{}
	unit.ID = "5"
	unit.PlayerIndex = 0
	unit.Location = interfaces.Location{X: 10, Y: 10}
	Game().BattleGround.PlaceBattleGroundEntity(&unit)

	friendly := Unit{}
	friendly.ID = "2"
	friendly.PlayerIndex = 0
	friendly.Location = interfaces.Location{X: 11, Y: 11}
	Game().BattleGround.PlaceBattleGroundEntity(&friendly)

	target := Unit{}
	target.ID = "4"
	target.PlayerIndex = 1
	target.Location = interfaces.Location{X: 5, Y: 5}
	Game().BattleGround.PlaceBattleGroundEntity(&target)

	outOfRange := Unit{}
	outOfRange.ID = "3"
	friendly.PlayerIndex = 1
	outOfRange.Location = interfaces.Location{X: 25, Y: 25}

	Game().BattleGround.PlaceBattleGroundEntity(&outOfRange)

	m.Run()
}

func TestShootingPhase(t *testing.T) {
	entites := InRange("5", 0, 6, 10, 10)
	assert.Equal(t, 1, len(entites))
}

func TestShootingPhaseLong(t *testing.T) {
	entites := InRange("160", 0, 49, 49, 10)
	assert.Equal(t, 2, len(entites))
}
