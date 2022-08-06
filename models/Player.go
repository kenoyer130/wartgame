package models

type Player struct {
	AI            bool
	Name          string
	Army          Army
	Gone          bool
	MoraleChecked bool
}
