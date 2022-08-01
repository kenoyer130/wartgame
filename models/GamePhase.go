package models

type GamePhase string

const (
	CommandPhase                  GamePhase = "Command Phase"
	MovementPhase                 GamePhase = "Movement Phase"
	ReinformentPhase              GamePhase = "Reinforment Phase"
	AircraftPhase                 GamePhase = "Aircraft Phase"
	PsychicPhase                  GamePhase = "Psychic Phase"
	ShootingPhase_UnitSelection   GamePhase = "ShootingPhase Unit Selection"
	ShootingPhase_TargetSelection GamePhase = "ShootingPhase Target Selection"
	ChargePhase                   GamePhase = "Charge Phase"
	FightPhase                    GamePhase = "Fight Phase"
	MoralePhase                   GamePhase = "Morale Phase"
	EndPhase                      GamePhase = "End Phase"
)
