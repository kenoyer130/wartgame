package phases

import (
	"testing"

	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/testutils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	testutils.InitGameState()

	targetUnit := models.Unit{}

	shootingUnit := models.Unit{}

	models.Game().SelectedPhaseUnit = &shootingUnit

	weapons := []string{"testW1"}

	for i := 0; i < 1; i++ {
		shooter := models.Model{
			Name:           "shooter",
			Wounds:         1,
			Weapons:        weapons,
			BallisticSkill: "3+",
		}

		shootingUnit.Models = append(shootingUnit.Models, shooter)
	}

	// set up a unit to target with two models so we don't destroy the model
	for i := 0; i < 2; i++ {
		target := models.Model{
			Name:      "target",
			Wounds:    1,
			Weapons:   weapons,
			Toughness: 3,
		}

		targetUnit.Models = append(targetUnit.Models, target)
	}

	models.Game().SelectedWeaponName = "testW1"
	models.Game().SelectedWeaponIndex = 0
	models.Game().SelectedPhaseUnit = &shootingUnit
	models.Game().SelectedTargetUnit = &targetUnit

	m.Run()
}

func TestShootingPhase(t *testing.T) {
	// assemble
	phase := ShootingAttackPhase{}

	roller := testutils.DiceRollerFake{
		Success: 1,
		Dice:    []int{6},
		Model:   models.Game().SelectedPhaseUnit.Models[0],
	}

	models.Game().DiceRoller = roller

	// act
	phase.Start()

	// assert
	assert.Equal(t, len(models.Game().SelectedTargetUnit.DestroyedModels), 0)
}