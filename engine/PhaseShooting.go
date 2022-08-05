package engine

import (
	"github.com/kenoyer130/wartgame/models"
)


func StartPhaseShooting(g *Game) {

	g.StatusMesssage = "Shooting Phase! Select a unit to shoot! Press [Q] and [E] to cycle units! Press [Space] to select unit to shoot!"

	unitCycler := NewUnitCycler(g.CurrentPlayer, g, UnitIsValidShooter, ShooterSelected) 

	unitCycler.CycleUnits()
}

func UnitIsValidShooter(unit *models.Unit) bool {
	return unit.CanShoot()
}

func ShooterSelected(unit *models.Unit, g *Game) {
	if(unit == nil) {
		//TODO: move to next phase
		return
	}

	g.SelectedPhaseUnit = unit
	WriteMessage("Selected Unit to shoot: " + unit.Name)
	StartPhaseShootingTargetting(unit, g)
}