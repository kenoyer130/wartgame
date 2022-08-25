package engine

import (
	"fmt"
	"image"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/ui"
)

type DiceRoller struct {
	DieImage map[int]*ebiten.Image
	Dice     []int
	msgs     []string
}

func (re DiceRoller) PlaySound() {
	PlaySound("Roll")
}

// rolls the indicated dice count and returns how many are equal to or greater then the target. Also returns each die result.
func (re DiceRoller) Roll(msg string, diceRollType interfaces.DiceRollType, onRoll func(die int) int) (int, []int) {

	re.msgs = []string{}
	re.msgs = append(re.msgs, msg)
	re.msgs = append(re.msgs, fmt.Sprintf("Rolling %d to hit target %d", diceRollType.Dice, diceRollType.Target))

	success := 0
	results := []int{}

	for i := 0; i < diceRollType.Dice; i++ {

		die := rand.Intn(6) + 1

		// allows application of abilities that need the raw roll
		if onRoll != nil {
			onRoll(die)
		}

		if diceRollType.AddToDice != 0 {
			re.msgs = append(re.msgs, fmt.Sprintf("applying modifier of %d to die %d", diceRollType.AddToDice, die))
			die = die + diceRollType.AddToDice
		}

		results = append(results, die)

		if die >= diceRollType.Target {
			success++
		}
	}

	success, dice := re.diceRolled(results, success, diceRollType)

	for _, msg := range re.msgs {
		WriteMessage(msg)
	}

	return success, dice
}

func (re *DiceRoller) diceRolled(results []int, success int, diceRollType interfaces.DiceRollType) (int, []int) {
	rolled := "Dice Rolled:"

	for i := 0; i < len(results); i++ {
		rolled += fmt.Sprintf(" %d", results[i])
	}

	re.msgs = append(re.msgs, rolled)
	msg := fmt.Sprintf("%d successes out of %d", success, diceRollType.Dice)
	re.msgs = append(re.msgs, msg)
	models.Game().Dice = results

	return success, results
}

type DiceRollerUI struct {
}

func (re DiceRoller) GetUIPanel(dice []int) *ebiten.Image {

	panel := ebiten.NewImage(400, 125)

	text.Draw(panel, "Dice", ui.GetFontBold(), ui.Margin, 25, ui.GetTextColor())

	ui.DrawSelectorBox(ui.Rect{
		X: 0,
		Y: 0,
		W: 399,
		H: 299,
	}, panel)

	// r := 0
	// c := 0

	for _, die := range re.Dice {

		re.loadDieImage(die)

		// op := &ebiten.DrawImageOptions{}
		// op.GeoM.Translate(float64(c*50), 10+float64(30+(r*25)))

		// c++

		// if c > 10 {
		// 	r++
		// }

		// panel.DrawImage(DieImage[die], op)
	}

	return panel
}

func (re DiceRoller) loadDieImage(die int) {
	if re.DieImage == nil {
		re.DieImage = make(map[int]*ebiten.Image)
	}

	if re.DieImage[die] == nil {
		img, _, err := ebitenutil.NewImageFromFile("./assets/graphics/dice.png")
		if err != nil {
			log.Fatal(err)
		}

		re.DieImage[die] = ebiten.NewImageFromImage(img.SubImage(image.Rect(4+(0*32), 5, 36, 37)))
	}
}
