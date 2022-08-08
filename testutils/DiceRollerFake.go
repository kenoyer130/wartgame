package testutils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/models"
)

type DiceRollerFake struct {
	Success	int
	Dice     []int
	Model	models.Model
}

func (re DiceRollerFake) Roll(msg string, diceRollType models.DiceRollType, onRolled func(int, []int)) {
	onRolled(re.Success, re.Dice)
}

func (re DiceRollerFake) GetUIPanel(dice []int) *ebiten.Image {
	return nil
}

func (re DiceRollerFake) GetDice() []int {
	return re.Dice
}
