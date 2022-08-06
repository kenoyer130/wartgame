package phases

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func StartPhaseShootingTargetting(unit *models.Unit, g *engine.Game) {

	g.StatusMesssage = "Targetting Phase! Select a unit to target! Press [Q] and [E] to cycle targets! Press [Space] to select target!"

	opponent := 0
	if g.CurrentPlayerIndex == 0 {
		opponent = 1
	}

	unitCycler := NewUnitCycler(&g.Players[opponent], g, func(target *models.Unit) bool {
		return canTarget(unit, target)
	}, func(target *models.Unit, g *engine.Game) {
		targetSelected(unit, target, g)
	})

	unitCycler.CycleUnits()
}

func canTarget(unit *models.Unit, target *models.Unit) bool {
	return true
}

func targetSelected(unit *models.Unit, target *models.Unit, g *engine.Game) {
	if unit == nil { 
		//TODO: move to next phase
		return
	}

	g.SelectedTargetUnit = target
	engine.WriteMessage("Selected Target: " + target.Name)

	engine.ClearKeyBoardRegistry()

	StartPhaseShootingWeapons(g)
}
