package testutils

import (
	"fmt"

	"github.com/kenoyer130/wartgame/models"
)

type PhaseStepperFake struct {
	CurrentPhase models.GamePhase	
}

func (re PhaseStepperFake) GetPhase() models.GamePhase {
	return models.AircraftPhase
}

func (re PhaseStepperFake) GetPhaseName() string {
	return fmt.Sprint(re.CurrentPhase)
}

func (re PhaseStepperFake) Move(phase models.GamePhase) {
}