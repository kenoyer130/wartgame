package models

type PhaseStepper interface {
	Move(phase GamePhase)
	GetPhaseName() string
}