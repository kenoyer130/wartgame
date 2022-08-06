package phases

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func MoveToNextPhaseOrder(phase models.GamePhase, g *engine.Game) {
	switch phase {
	case models.ShootingPhase_UnitSelection:

		engine.WriteMessage("Starting Shooting Phase")
		g.CurrentPhase = models.ShootingPhase_UnitSelection
		StartPhaseShooting(g)
	}
}
