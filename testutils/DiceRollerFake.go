package testutils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
)

type DiceRollerFake struct {
	Success int
	Dice    []int
	Model   models.Model
}

func (re DiceRollerFake) PlaySound() {}

func (re DiceRollerFake) Roll(msg string, diceRollType interfaces.DiceRollType, onRoll func(die int) int) (int, []int) {
	return re.Success, re.Dice
}

func (re DiceRollerFake) GetUIPanel(dice []int) *ebiten.Image {
	return nil
}
