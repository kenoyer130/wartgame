package models

type Phase string
type PhaseStep string

const (
	MovementPhase               Phase     = "MovementPhase"
	MovementPhase_UnitSelection PhaseStep = "MovementPhase_UnitSelection"
	MovementPhase_Moving        PhaseStep = "MovementPhase_Moving"
)