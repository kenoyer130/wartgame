package interfaces

type PhaseStepper interface {
	Move(phase GamePhase)
	GetPhase() GamePhase
	GetPhaseName() GamePhase
}