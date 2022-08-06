package phases

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func StartPhaseShootingTargetting(unit *models.Unit) {

	models.Game().StatusMessage.Messsage = "Targetting Phase! Select a unit to target! "
	models.Game().StatusMessage.Keys = "Press [Q] and [E] to cycle targets! Press [Space] to select!"

	opponent := 0
	if models.Game().CurrentPlayerIndex == 0 {
		opponent = 1
	}

	unitCycler := NewUnitCycler(&models.Game().Players[opponent], func(target *models.Unit) bool {
		return canTarget(unit, target)
	}, func(target *models.Unit) {
		targetSelected(unit, target)
	})

	unitCycler.CycleUnits()
}

func canTarget(unit *models.Unit, target *models.Unit) bool {
	return true
}

func targetSelected(unit *models.Unit, target *models.Unit) {
	if unit == nil { 
		//TODO: move to next phase
		return
	}

	models.Game().SelectedTargetUnit = target
	engine.WriteMessage("Selected Target: " + target.Name)

	engine.ClearKeyBoardRegistry()

	StartPhaseShootingWeapons()
}
