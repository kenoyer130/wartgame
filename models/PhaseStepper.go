package models

type PhaseStepper interface {
	Move(phase GamePhase)
	GetPhase() GamePhase
	GetPhaseName() string
}