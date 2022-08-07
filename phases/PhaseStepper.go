package phases

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func MoveToPhase(phase models.GamePhase) {

	newPhase := phase != models.Game().CurrentPhase

	if newPhase {
		cleanupPreviousPhase()
		models.Game().CurrentPhase = phase
	}

	switch phase {
	case models.ShootingPhase:

		printPhase(newPhase, "Starting Shooting Phase")		
		StartPhaseShooting()

	case models.ChargePhase:
		printPhase(newPhase, "Starting Charge Phase")
	
	case models.MoralePhase:

		printPhase(newPhase, "Starting Morale Phase")		
		StartPhaseMorale()

	case models.EndPhase:
		printPhase(newPhase, "Starting End Phase")
		StartRoundEnd()
	}
}

func printPhase(print bool, msg string) {
	if print {
		engine.WriteMessage(msg)
	}
}

func cleanupPreviousPhase() {
	models.Game().SelectedPhaseUnit = nil
	models.Game().SelectedTargetUnit = nil
	models.Game().SelectedUnit = nil
	models.Game().StatusMessage.Clear()
}
