package phases

import "github.com/kenoyer130/wartgame/interfaces"

type PhaseGameOver struct {
}

func (re PhaseGameOver) Start() {
}

func (re PhaseGameOver) GetName() (interfaces.GamePhase, interfaces.PhaseStep) {
	return interfaces.GameOverPhase, interfaces.Nil
}
