package phases

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func MoveToPhase(phase models.GamePhase) {

	printMsg := phase != models.Game().CurrentPhase

	switch phase {
	case models.ShootingPhase:

		printPhase(printMsg, "Starting Shooting Phase")
		models.Game().CurrentPhase = models.ShootingPhase
		StartPhaseShooting()

	case models.ChargePhase:
		printPhase(printMsg, "Starting Charge Phase")
	}
}

func printPhase(print bool, msg string) {
	if print {
		engine.WriteMessage(msg)
	}
}
