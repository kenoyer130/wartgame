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
