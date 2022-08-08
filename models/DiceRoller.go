package models

import "github.com/hajimehoshi/ebiten/v2"

type DiceRoller interface {
	Roll(msg string, diceRollType DiceRollType, onRolled func(int, []int))
	GetUIPanel(dice []int) *ebiten.Image
	GetDice() []int
}

type DiceRollType struct {
	Dice      int
	Target    int
	AddToDice int
}
