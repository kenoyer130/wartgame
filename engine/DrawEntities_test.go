package engine

import (
	"testing"

	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/testutils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	testutils.InitGameState()

	phaseStepper := models.Game().PhaseStepper
	s := phaseStepper.(testutils.PhaseStepperFake)
	s.CurrentPhase = interfaces.MovementPhase
	models.Game().PhaseStepper = s

	models.Game().BattleGround = *models.NewBattleGround(50, 50)

	unit := models.Unit{Name: "Test"}
	unit.Location = models.Location{
		X: 10,
		Y: 10,
	}

	models.Game().Players[0].Army.Units = append(models.Game().Players[0].Army.Units, &unit)
	models.Game().SelectedPhaseUnit = &unit

	for i := 0; i < 9; i++ {
		models.Game().Players[0].Army.Units[0].Models = append(models.Game().Players[0].Army.Units[0].Models, &models.Model{
			Movement: 6,
		})
	}

	unit.Place()

	m.Run()
}

// assuming a 9 model unit then the rect would be 3 x 3
func TestUnitRange(t *testing.T) {
	// assemble
	unit := models.Game().Players[0].Army.Units[0]

	// act
	drawMoveRange(testutils.DrawerFake{})

	// assert
	assert.Equal(t, 3, unit.Width)
	assert.Equal(t, 3, unit.Height)
}
