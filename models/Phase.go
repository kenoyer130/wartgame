package models

type Phase int64

const (
	Command  Phase = 0
	Movement Phase = 1
	Psychic  Phase = 2
	Shooting Phase = 3
	Charge   Phase = 4
	Fight    Phase = 5
	Morale   Phase = 6
)
