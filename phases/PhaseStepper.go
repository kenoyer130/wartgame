package phases

import (
	"fmt"

	"github.com/kenoyer130/wartgame/engine"
	interfaces "github.com/kenoyer130/wartgame/engine/Interfaces"
	"github.com/kenoyer130/wartgame/models"
)

type PhaseStepper struct {
	CurrentPhase         interfaces.GamePhase
	CurrentPhaseExecuter interfaces.PhaseExecute
}

func (re PhaseStepper) printPhase(msg string) {
	engine.WriteMessage(msg)
}

func (re PhaseStepper) GetPhase() interfaces.GamePhase {
	return re.CurrentPhase
}

func (re PhaseStepper) GetPhaseName() string {
	return fmt.Sprintf("%s", re.CurrentPhase)
}

func (re *PhaseStepper) Move(phase interfaces.GamePhase) {

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

func (re PhaseStepper) phaseLookup(phase interfaces.GamePhase) interfaces.PhaseExecute {
	if phase == interfaces.MovementPhase {
		return MovePhase{}
	} else if phase == interfaces.ShootingPhase {
		return ShootingPhase{}
	} else if phase == interfaces.MoralePhase {
		return MoralePhase{}
	} else if phase == interfaces.EndPhase {
		return EndPhase{}
	}

	return nil
}
