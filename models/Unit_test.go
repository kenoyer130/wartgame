package models

import (
	"testing"

	"github.com/kenoyer130/wartgame/ui"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	Game().BattleGround = *NewBattleGround(50, 50)

	Game().Players[0].Army.Units = append(Game().Players[0].Army.Units, &Unit{})

	for i := 0; i < 9; i++ {
		Game().Players[0].Army.Units[0].Models = append(Game().Players[0].Army.Units[0].Models, &Model{
			Movement: 6,
		})
	}

	m.Run()
}

// assuming a 9 model unit then the rect would be 3 x 3
func TestUnitSetsSizeBasedOnModelCount(t *testing.T) {
	// assemble
	unit := Game().Players[0].Army.Units[0]

	// act
	unit.Place()

	// assert
	assert.Equal(t, 3, unit.Width)
	assert.Equal(t, 3, unit.Height)
}

func TestUnitMoveRange(t *testing.T) {
	// assemble
	unit := Game().Players[0].Army.Units[0]

	// act
	unit.Place()

	// assert
	assert.Equal(t, ui.Rect{X: -6, Y: -6, W: 15, H: 15}, unit.MovementRect)
}
