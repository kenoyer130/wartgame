package phases

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func StartPhaseShooting() {

	models.Game().StatusMessage.Phase = "Shooting Phase"
	models.Game().StatusMessage.Messsage = "Select next unit to shoot!"
	models.Game().StatusMessage.Keys = "Press [Q] and [E] to cycle units! Press [Space] to select!"

	unitCycler := NewUnitCycler(models.Game().CurrentPlayer, UnitIsValidShooter, ShooterSelected) 

	unitCycler.CycleUnits()
}

func UnitIsValidShooter(unit *models.Unit) bool {
	return unit.CanShoot()
}

func ShooterSelected(unit *models.Unit) {
	if(unit == nil) {
		engine.WriteMessage("No valid units for shooting phase.")
		MoveToPhase(models.MoralePhase)
		return
	}

	models.Game().SelectedPhaseUnit = unit
	engine.WriteMessage("Selected Unit to shoot: " + unit.Name)

	StartPhaseShootingTargetting(unit)
}