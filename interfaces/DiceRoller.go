package interfaces

import "github.com/hajimehoshi/ebiten/v2"

type DiceRoller interface {
	Roll(msg string, diceRollType DiceRollType, onRoll func(die int) int, onRolled func(int, []int))
	GetUIPanel(dice []int) *ebiten.Image
	GetDice() []int
	Suppress(suppress bool)
}

type DiceRollType struct {
	Dice      int
	Target    int
	AddToDice int
}
