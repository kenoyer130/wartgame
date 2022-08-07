package models

type Player struct {
	AI            bool
	Name          string
	Army          Army
	Gone          bool
	MoraleChecked bool
}

func (re *Player) PhaseCleanup() {
	re.MoraleChecked = false
}

func (re *Player) RoundCleanup() {
	re.Gone = false
}
