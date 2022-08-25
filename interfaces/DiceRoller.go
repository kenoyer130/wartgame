package interfaces

import "github.com/hajimehoshi/ebiten/v2"

type DiceRoller interface {
	PlaySound()
	Roll(msg string, diceRollType DiceRollType, onRoll func(die int) int)(int, []int)
	GetUIPanel(dice []int) *ebiten.Image		
}

type DiceRollType struct {
	Dice      int
	Target    int
	AddToDice int
}
