package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/ui"
)

func GameEngineDraw(screen *ebiten.Image) {

	background := ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())

	DrawMainView(background)

	drawGameInfoPanel(screen)
	drawModelPanel(screen)
	drawMessagePanel(screen)
	drawStatusPanel(screen)
	drawDiceRollerPanel(screen)

	screen.DrawImage(background, nil)
}

func drawDiceRollerPanel(screen *ebiten.Image) {
	diceRollerPanel := getDiceRollerPanel(models.Game().Dice)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(getLeftXStartingPixel() + ui.Margin, 700)
	screen.DrawImage(diceRollerPanel, op)
}

func drawStatusPanel(screen *ebiten.Image) {
	statusPanel := getStatusPanel(models.Game().StatusMesssage)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(25, 950)
	screen.DrawImage(statusPanel, op)
}

func drawMessagePanel(screen *ebiten.Image) {
	messagePanel := getMessagePanel()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(getLeftXStartingPixel(), 125)
	screen.DrawImage(messagePanel, op)
}

func drawGameInfoPanel(screen *ebiten.Image) {
	gameInfoPanel := getGameInfoPanel()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(getLeftXStartingPixel(), 0)
	screen.DrawImage(gameInfoPanel, op)
}

func drawModelPanel(screen *ebiten.Image) {
	ModelPanel := getUnitPanel(models.Game().SelectedUnit)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(getLeftXStartingPixel(), 275)
	screen.DrawImage(ModelPanel, op)
}

func getLeftXStartingPixel() float64 {
	return float64(models.Game().BattleGround.ViewPort.GetPixelRectangle().Width + ui.Margin)
}

func drawSelectedModelInfo(background *ebiten.Image) {
	if models.Game().SelectedModel == nil {
		return
	}
}
