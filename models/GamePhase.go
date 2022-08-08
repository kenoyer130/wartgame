package models

type GamePhase string

const (
	CommandPhase     GamePhase = "Command Phase"
	MovementPhase    GamePhase = "Movement Phase"
	ReinformentPhase GamePhase = "Reinforment Phase"
	AircraftPhase    GamePhase = "Aircraft Phase"
	PsychicPhase     GamePhase = "Psychic Phase"
	ShootingPhase    GamePhase = "ShootingPhase"
	ChargePhase      GamePhase = "Charge Phase"
	FightPhase       GamePhase = "Fight Phase"
	MoralePhase      GamePhase = "Morale Phase"
	EndPhase         GamePhase = "End Phase"
)

type PhaseStep string

const (
	// not all phases have steps
	Nil PhaseStep = "Nil"

	//shooting
	ShootingPhaseAttack    PhaseStep = "ShootingPhaseAttack"
	ShootingPhaseTargeting PhaseStep = "ShootingPhaseTargeting"
	ShootingPhaseWeapons   PhaseStep = "ShootingPhaseWeapons"
)

type PhaseExecute interface {
	Start()
	GetName() (GamePhase, PhaseStep)
}
