package phases

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
)

type PhaseShootingUnitSelector struct {
	
}

func (re PhaseShootingUnitSelector) GetName() (interfaces.GamePhase, interfaces.PhaseStep) {
	return interfaces.ShootingPhase, interfaces.Nil
}

func (re PhaseShootingUnitSelector) Start() {

	models.Game().StatusMessage.Phase = "Shooting Phase"

	engine.WriteStatusMessage("Select next unit to shoot!")
	engine.WriteStatusKeys( "Press [Q] and [E] to cycle units! Press [Space] to select!")
	re.loop()
}

func (re PhaseShootingUnitSelector) loop() {
	unitCycler := NewUnitCycler(models.Game().CurrentPlayer, re.UnitIsValidShooter, re.ShooterSelected, false)
	unitCycler.CycleUnits()
}

func (re PhaseShootingUnitSelector) UnitIsValidShooter(unit *models.Unit) bool {
	return unit.CanShoot()
}

func (re PhaseShootingUnitSelector) ShooterSelected(unit *models.Unit) {
	if unit == nil {
		engine.WriteMessage("No valid units for shooting phase.")
		models.Game().PhaseEventBus.Fire("ShooterPhaseEnded")
		return
	}

	models.Game().SelectedPhaseUnit = unit
	engine.WriteMessage("Selected Unit to shoot: " + unit.Name)
	models.Game().PhaseEventBus.Fire("ShooterSelected")	
}
