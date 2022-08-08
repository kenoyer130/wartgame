package phases

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

type ShootingTargetingPhase struct {
}

func (re ShootingTargetingPhase) GetName() (models.GamePhase, models.PhaseStep) {
	return models.ShootingPhase, models.Nil
}

func (re ShootingTargetingPhase) Start() {

	unit := models.Game().SelectedPhaseUnit

	models.Game().StatusMessage.Messsage = "Targetting Phase! Select a unit to target! "
	models.Game().StatusMessage.Keys = "Press [Q] and [E] to cycle targets! Press [Space] to select!"

	opponent := 0
	if models.Game().CurrentPlayerIndex == 0 {
		opponent = 1
	}

	unitCycler := NewUnitCycler(&models.Game().Players[opponent], func(target *models.Unit) bool {
		return re.canTarget(unit, target)
	}, func(target *models.Unit) {
		re.targetSelected(unit, target)
	})

	unitCycler.CycleUnits()
}

func (re ShootingTargetingPhase) canTarget(unit *models.Unit, target *models.Unit) bool {
	return true
}

func (re ShootingTargetingPhase) targetSelected(unit *models.Unit, target *models.Unit) {
	if unit == nil {
		//TODO: move to next phase
		return
	}

	models.Game().SelectedTargetUnit = target
	engine.WriteMessage("Selected Target: " + target.Name)

	engine.ClearKeyBoardRegistry()

	models.Game().PhaseStepper.Move(models.ShootingPhase, models.ShootingPhaseWeapons)
}
