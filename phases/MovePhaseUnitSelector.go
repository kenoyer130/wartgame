package phases

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

type MovePhaseUnitSelector struct {
	
}

func (re MovePhaseUnitSelector) Start() {

	re.clearMovementKeys()

	models.Game().StatusMessage.Phase = "Movement Phase"	

	unitCycler := NewUnitCycler(models.Game().CurrentPlayer, re.UnitCanMove, re.MoverSelected, true)

	unitCycler.CycleUnits()
}

func (re MovePhaseUnitSelector) UnitCanMove(unit *models.Unit) bool {
	return unit.CanMove()
}

func (re MovePhaseUnitSelector) MoverSelected(unit *models.Unit) {
	if unit == nil {
		engine.WriteMessage("No valid units for movement phase.")
		models.Game().PhaseEventBus.Fire("MovePhaseEnded")
		return
	}

	models.Game().SelectedPhaseUnit = unit
	engine.WriteStatusMessage("Selected Unit to move: " + unit.Name)
	engine.WriteStatusKeys("Moving! Press [Space] to end movement. Press [A] to advance!")

	unit.CurrentMoves = unit.Models[0].Movement

	re.registerMovementKeys(unit)
}

func (re MovePhaseUnitSelector) clearMovementKeys() {

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

func (re MovePhaseUnitSelector) registerMovementKeys(unit *models.Unit) {
	
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
		unit.Location.X = unit.Location.X + 1
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

func (re MovePhaseUnitSelector) doMove(unit *models.Unit) {
	unit.CurrentMoves--

	engine.WriteStatusKeys( fmt.Sprintf("%d moves left! Use [Num] keys to move. Press [Space] to for next unit!", unit.CurrentMoves))

	if unit.CurrentMoves < 1 {
		unit.AddState(models.UnitMoved)
		re.Start()
	}
}
