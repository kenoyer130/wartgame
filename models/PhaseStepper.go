package models

type PhaseStepper interface {
	Move(phase GamePhase, phaseStep PhaseStep)
	GetPhaseName() string
}