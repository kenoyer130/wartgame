package engine

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/ui"
)

type DiceRoller struct {
	DieImage map[int]*ebiten.Image
	Dice     []int
}

func (re DiceRoller) GetDice() []int {
	return re.Dice
}

// rolls the indicated dice count and returns how many are equal to or greater then the target. Also returns each die result.
func (re DiceRoller) Roll(msg string, diceRollType models.DiceRollType, onRolled func(int, []int)) {

	WriteMessage(msg)
	WriteMessage(fmt.Sprintf("Rolling %d to hit target %d", diceRollType.Dice, diceRollType.Target))

	success := 0
	results := []int{}

	PlaySound("Roll")

	for i := 0; i < diceRollType.Dice; i++ {

		die := rand.Intn(6) + 1

		if diceRollType.AddToDice > 0 {
			die = die + diceRollType.AddToDice
		}

		results = append(results, die)

		if die >= diceRollType.Target {
			success++
		}
	}

	diceTimer := time.NewTimer(1 * time.Second)

	go func() {
		<-diceTimer.C

		rolled := "Dice Rolled:"

		for i := 0; i < len(results); i++ {
			rolled += fmt.Sprintf(" %d", results[i])
		}

		WriteMessage(rolled)
		WriteMessage(fmt.Sprintf("%d successes out of %d", success, diceRollType.Dice))
		models.Game().Dice = results

		models.Game().StatusMessage.Keys = "Press [Space] to continue"

		dicePauseTimer := time.NewTimer(500 * time.Millisecond)

		go func() {
			<-dicePauseTimer.C
			onRolled(success, results)
		}()
	}()
}

type DiceRollerUI struct {
	dice []int
}

func (re DiceRoller) GetUIPanel(dice []int) *ebiten.Image {

	panel := ebiten.NewImage(400, 300)

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
