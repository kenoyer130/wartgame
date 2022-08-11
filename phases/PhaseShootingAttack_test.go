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

	weapons := []models.Weapon{models.Weapon{Name: "testW1"}}

	for i := 0; i < 1; i++ {
		shooter := models.Model{
			Name:           "shooter",
			Wounds:         1,
			Weapons:        weapons,
			BallisticSkill: "3+",
		}

		shootingUnit.Models = append(shootingUnit.Models, &shooter)
	}

	// set up a unit to target with two models so we don't destroy the model
	for i := 0; i < 2; i++ {
		target := models.Model{
			Name:      "target",
			Wounds:    1,
			Weapons:   weapons,
			Toughness: 3,
		}

		targetUnit.Models = append(targetUnit.Models, &target)
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
		Model:   *models.Game().SelectedPhaseUnit.Models[0],
	}

	models.Game().DiceRoller = roller

	// act
	phase.Start()

	// assert
	assert.Equal(t, len(models.Game().SelectedTargetUnit.DestroyedModels), 0)
}

func TestSetModelsByToughness(t *testing.T) {
	// assemble
	phase := ShootingAttackPhase{}

	targetUnit := models.Unit{}
	targets := [3]*models.Model{}

	targets[0] = &models.Model{Toughness: 3}
	targets[1] = &models.Model{Toughness: 3}
	targets[2] = &models.Model{Toughness: 5}

	targetUnit.Models = targets[:]

	models.Game().SelectedTargetUnit = &targetUnit

	// act
	targetUnits := phase.setModelsByToughness()

	// assert
	assert.Equal(t, 2, targetUnits.Count())

	pop1, _ := targetUnits.Pop()
	firstSet := pop1.(models.Stack)
	assert.Equal(t, 1, firstSet.Count())

	pop2, _ := targetUnits.Pop()
	secondSet := pop2.(models.Stack)
	assert.Equal(t, 2, secondSet.Count())
}
