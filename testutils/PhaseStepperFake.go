package testutils

import (
	"github.com/kenoyer130/wartgame/interfaces"
)

type PhaseStepperFake struct {
	CurrentPhase interfaces.GamePhase
}

func (re PhaseStepperFake) GetPhase() interfaces.GamePhase {
	return re.CurrentPhase
}

func (re PhaseStepperFake) GetPhaseName() interfaces.GamePhase {
	return re.CurrentPhase
}

func (re PhaseStepperFake) Move(phase interfaces.GamePhase) {
}
