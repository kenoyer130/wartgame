package phases

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

type MovePhase struct {
	ShootingTargetingPhase *ShootingTargetingPhase
	ShootingAttackPhase    *ShootingAttackPhase
	ShootingWeaponPhase    *ShootingWeaponPhase
}

func (re MovePhase) GetName() (models.GamePhase, models.PhaseStep) {
	return models.ShootingPhase, models.Nil
}

func (re MovePhase) Start() {

	models.Game().StatusMessage.Phase = "Movement Phase"
	models.Game().StatusMessage.Messsage = "Select next unit to Move!"
	models.Game().StatusMessage.Keys = "Press [Q] and [E] to cycle units! Press [X] to Remain Stationary. Press [Space] to select!"

	unitCycler := NewUnitCycler(models.Game().CurrentPlayer, re.UnitCanMove, re.MoverSelected)

	unitCycler.CycleUnits()
}

func (re MovePhase) UnitCanMove(unit *models.Unit) bool {
	return unit.CanMove()
}

func (re MovePhase) MoverSelected(unit *models.Unit) {
	if unit == nil {
		engine.WriteMessage("No valid units for movement phase.")
		models.Game().PhaseStepper.Move(models.ShootingPhase)
		return
	}

	models.Game().SelectedPhaseUnit = unit
	engine.WriteMessage("Selected Unit to move: " + unit.Name)

	models.Game().StatusMessage.Keys = fmt.Sprintf("%d moves left! Use [Num] keys to move. Press [Space] to for next unit!", unit.CurrentMoves)

	unit.CurrentMoves = unit.Models[0].Movement
	
	re.registerMovementKeys(unit)
}

func (re MovePhase) registerMovementKeys(unit *models.Unit) {	

	engine.KeyBoardRegistry[ebiten.KeySpace] = func() {
		unit.AddState(models.UnitMoved)
		re.Start()
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad8] = func() {
		unit.AddState(models.UnitMoved)
		re.Start()
	}
}
