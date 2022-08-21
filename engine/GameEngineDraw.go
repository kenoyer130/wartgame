package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/ui"
)

func GameEngineDraw(screen *ebiten.Image) {

	background := ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())

	drawMainView(background)

	drawGameInfoPanel(background)
	drawModelPanel(background)	
	drawTopStatusPanel(background)
	drawStatusPanel(background)
	drawMessagePanel(background)

	screen.DrawImage(background, nil)
}

func drawMainView(screen *ebiten.Image) {
	mainViewPanel := GetMainView()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 40)
	screen.DrawImage(mainViewPanel, op)
}

func drawStatusPanel(screen *ebiten.Image) {
	statusPanel := getStatusPanel()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(25, 950)
	screen.DrawImage(statusPanel, op)
}

func drawTopStatusPanel(screen *ebiten.Image) {
	statusPanel := getTopStatusPanel()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(150, 5)
	screen.DrawImage(statusPanel, op)
}

func drawMessagePanel(screen *ebiten.Image) {
	messagePanel := getMessagePanel()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(getLeftXStartingPixel()+ui.Margin, 600)
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
	op.GeoM.Translate(getLeftXStartingPixel(), 155)
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
