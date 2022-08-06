package phases

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)


func StartPhaseShooting() {

	models.Game().StatusMesssage = "Select next unit to shoot! Press [Q] and [E] to cycle units! Press [Space] to select unit to shoot!"

	unitCycler := NewUnitCycler(models.Game().CurrentPlayer, UnitIsValidShooter, ShooterSelected) 

	unitCycler.CycleUnits()
}

func UnitIsValidShooter(unit *models.Unit) bool {
	return unit.CanShoot()
}

func ShooterSelected(unit *models.Unit) {
	if(unit == nil) {
		engine.WriteMessage("No valid units for shootinmodels.Game().")
		MoveToPhase(models.ChargePhase)
		return
	}

	models.Game().SelectedPhaseUnit = unit
	engine.WriteMessage("Selected Unit to shoot: " + unit.Name)

	unit.AddState(models.UnitShot)

	StartPhaseShootingTargetting(unit)
}