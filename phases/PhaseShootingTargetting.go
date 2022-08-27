package phases

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
)

type ShootingTargetingPhase struct {
	OnCompleted func()
}

func (re ShootingTargetingPhase) GetName() (interfaces.GamePhase, interfaces.PhaseStep) {
	return interfaces.ShootingPhase, interfaces.Nil
}

func (re ShootingTargetingPhase) Start() {

	unit := models.Game().SelectedPhaseUnit

	models.Game().StatusMessage.Messsage = "Targetting Phase! Select a unit to target! "
	models.Game().StatusMessage.Keys = "Press [Q] and [E] to cycle targets! Press [Space] to select!"

	opponent := models.Game().OpponetPlayerIndex

	unitCycler := NewUnitCycler(&models.Game().Players[opponent], func(target *models.Unit) bool {
		return re.canTarget(unit, target)
	}, func(target *models.Unit) {
		re.targetSelected(unit, target)
	}, false)

	unitCycler.CycleUnits()
}

func (re ShootingTargetingPhase) canTarget(unit *models.Unit, target *models.Unit) bool {

	for _, entity := range models.Game().SelectedWeapon.Targets {
		if entity.GetID() == target.GetID() {
			return true
		}
	}

	return false
}

func (re ShootingTargetingPhase) targetSelected(unit *models.Unit, target *models.Unit) {
	if target == nil {
		engine.WriteMessage("No Targets in range!")
		re.OnCompleted()
		return
	}

	models.Game().SelectedTargetUnit = target
	engine.WriteMessage("Selected Target: " + target.Name)

	engine.ClearKeyBoardRegistry()

	shootingAttackPhase := ShootingAttackPhase{}

	shootingAttackPhase.OnCompleted = func() {
		re.OnCompleted()
	}

	shootingAttackPhase.Start()
}
