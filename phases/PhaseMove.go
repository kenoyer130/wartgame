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

	re.clearMovementKeys()

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
	
	unit.CurrentMoves = unit.Models[0].Movement	

	re.registerMovementKeys(unit)
}

func (re MovePhase) clearMovementKeys() {

	engine.KeyBoardRegistry[ebiten.KeyNumpad9] = func() {
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad8] = func() {
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad7] = func() {
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad6] = func() {
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad4] = func() {
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad3] = func() {
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad2] = func() {
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad1] = func() {
	}
}

func (re MovePhase) registerMovementKeys(unit *models.Unit) {

	engine.KeyBoardRegistry[ebiten.KeySpace] = func() {
		unit.AddState(models.UnitMoved)
		re.Start()
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad9] = func() {
		unit.Location.Y = unit.Location.Y - 1
		unit.Location.X = unit.Location.X + 1
		re.doMove(unit)
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad8] = func() {
		unit.Location.Y = unit.Location.Y - 1
		re.doMove(unit)
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad7] = func() {
		unit.Location.Y = unit.Location.Y - 1
		unit.Location.X = unit.Location.X - 1
		re.doMove(unit)
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad6] = func() {
		unit.Location.X = unit.Location.X + 1
		re.doMove(unit)
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad4] = func() {
		unit.Location.X = unit.Location.X - 1
		re.doMove(unit)
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad3] = func() {
		unit.Location.Y = unit.Location.Y + 1
		unit.Location.Y = unit.Location.X + 1
		re.doMove(unit)
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad2] = func() {
		unit.Location.Y = unit.Location.Y + 1
		re.doMove(unit)
	}

	engine.KeyBoardRegistry[ebiten.KeyNumpad1] = func() {
		unit.Location.Y = unit.Location.Y + 1
		unit.Location.X = unit.Location.X - 1
		re.doMove(unit)
	}
}

func (re MovePhase) doMove(unit *models.Unit) {
	unit.CurrentMoves--

	models.Game().StatusMessage.Keys = fmt.Sprintf("%d moves left! Use [Num] keys to move. Press [Space] to for next unit!", unit.CurrentMoves)

	engine.SetUnitFormation(engine.StandardUnitFormation, unit)
	if unit.CurrentMoves < 1 {
		unit.AddState(models.UnitMoved)
		re.Start()
	}
}
