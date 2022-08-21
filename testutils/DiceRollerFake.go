package testutils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
)

type DiceRollerFake struct {
	Success	int
	Dice     []int
	Model	models.Model
}

func (re DiceRollerFake) Suppress(suppress bool) {

}

func (re DiceRollerFake) Roll(msg string, diceRollType interfaces.DiceRollType,  onRoll func(die int) int, onRolled func(int, []int)) {
	onRolled(re.Success, re.Dice)
}

func (re DiceRollerFake) GetUIPanel(dice []int) *ebiten.Image {
	return nil
}

func (re DiceRollerFake) GetDice() []int {
	return re.Dice
}
