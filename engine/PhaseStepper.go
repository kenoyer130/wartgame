package engine

import "github.com/kenoyer130/wartgame/models"

func MoveToNextPhaseOrder(phase models.GamePhase, g *Game) {
	switch phase {
	case models.ShootingPhase_UnitSelection:

		WriteMessage("Starting Shooting Phase")
		g.CurrentPhase = models.ShootingPhase_UnitSelection
		StartPhaseShooting(g)
	}
}
