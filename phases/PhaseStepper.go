package phases

import (
	"fmt"

	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

type PhaseStepper struct {
	CurrentPhase         models.GamePhase
	CurrentPhaseExecuter models.PhaseExecute
}

func (re PhaseStepper) printPhase(msg string) {
	engine.WriteMessage(msg)
}

func (re PhaseStepper) GetPhase() models.GamePhase {
	return re.CurrentPhase
}

func (re PhaseStepper) GetPhaseName() string {
	return fmt.Sprintf("%s", re.CurrentPhase)
}

func (re *PhaseStepper) Move(phase models.GamePhase) {

	newPhase := false

	if re.CurrentPhase == "" || re.CurrentPhase != phase {
		newPhase = true
		re.CurrentPhase = phase
	} else {
		newPhase = (re.CurrentPhase != phase)
	}

	if newPhase {
		re.CurrentPhaseExecuter = re.phaseLookup(phase)
		re.cleanupPreviousPhase()
		re.printPhase(fmt.Sprintf("Starting phase %s", phase))
	}

	re.CurrentPhaseExecuter.Start()
}


func (re *PhaseStepper) cleanupPreviousPhase() {
	models.Game().SelectedPhaseUnit = nil
	models.Game().SelectedTargetUnit = nil
	models.Game().SelectedUnit = nil
	models.Game().StatusMessage.Clear()
}

func (re PhaseStepper) phaseLookup(phase models.GamePhase) models.PhaseExecute {
	if phase == models.MovementPhase {
		return MovePhase{}
	} else if phase == models.ShootingPhase {
		return ShootingPhase{}
	} else if phase == models.MoralePhase {
		return MoralePhase{}
	} else if phase == models.EndPhase {
		return EndPhase{}
	}

	return nil
}
