package engine

import "github.com/kenoyer130/wartgame/models"

func MoveToShootingPhaseUnitSelection(g *Game) {

	validSelection := selectValidShootingUnit(g)
	if !validSelection {
		MoveToNextPhaseOrder(models.ChargePhase, g)
		return
	}
}

func selectValidShootingUnit(g *Game) bool {
	return true
}

func isValidShootingUnit(unit *models.Unit) bool {
	return unit.
}
