package phases

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
)

type PhaseShootingTargetting struct {
	
}

func (re PhaseShootingTargetting) GetName() (interfaces.GamePhase, interfaces.PhaseStep) {
	return interfaces.ShootingPhase, interfaces.Nil
}

func (re PhaseShootingTargetting) Start() {

	unit := models.Game().SelectedPhaseUnit

	engine.WriteStatusMessage("Targetting Phase! Select a unit to target!")
	engine.WriteStatusKeys("Press [Q] and [E] to cycle targets! Press [Space] to select!")

	opponent := models.Game().OpponetPlayerIndex

	unitCycler := NewUnitCycler(&models.Game().Players[opponent], func(target *models.Unit) bool {
		return re.canTarget(unit, target)
	}, func(target *models.Unit) {
		re.targetSelected(unit, target)
	}, false)

	unitCycler.CycleUnits()
}

func (re PhaseShootingTargetting) canTarget(unit *models.Unit, target *models.Unit) bool {

	for _, entity := range models.Game().SelectedWeapon.Targets {
		if entity.GetID() == target.GetID() {
			return true
		}
	}

	return false
}

func (re PhaseShootingTargetting) targetSelected(unit *models.Unit, target *models.Unit) {
	if target == nil {
		models.Game().PhaseEventBus.Fire("ShooterAttackCompleted")
		return
	}

	models.Game().SelectedTargetUnit = target
	engine.WriteMessage("Selected Target: " + target.Name)

	models.Game().PhaseEventBus.Fire("ShooterWeaponTargetSelected")
}
