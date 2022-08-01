package engine

import "github.com/kenoyer130/wartgame/models"

func MoveToNextPhaseOrder(phase models.GamePhase, g *Game) {
	switch phase {
	case models.ShootingPhase_UnitSelection:
		validSelection := selectValidShootingUnit(g)
		if(!validSelection) {
			MoveToNextPhaseOrder(models.ChargePhase, g)
			break
		}
	}
}

func selectValidShootingUnit(g *Game) bool {
	panic("unimplemented")
}
