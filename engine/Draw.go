package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/ui"
)

func (g *Game) Draw(screen *ebiten.Image) {

	background := ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())

	DrawMainView(g, background)

	drawGameInfoPanel(g, screen)
	drawModelPanel(g, screen)
	drawMessagePanel(g, screen)

	screen.DrawImage(background, nil)
}

func drawMessagePanel(g *Game, screen *ebiten.Image) {
	messagePanel := getMessagePanel(g)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(getLeftXStartingPixel(g), 125)
	screen.DrawImage(messagePanel, op)
}

func drawGameInfoPanel(g *Game, screen *ebiten.Image) {
	gameInfoPanel := getGameInfoPanel(g)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(getLeftXStartingPixel(g), 0)
	screen.DrawImage(gameInfoPanel, op)
}

func drawModelPanel(g *Game, screen *ebiten.Image) {
	ModelPanel := getModelPanel(g)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(getLeftXStartingPixel(g), 275)
	screen.DrawImage(ModelPanel, op)
}

func getLeftXStartingPixel(g *Game) float64 {
	return float64(g.BattleGround.ViewPort.GetPixelRectangle().Width + ui.Margin)
}

func drawSelectedModelInfo(g *Game, background *ebiten.Image) {
	if g.SelectedModel == nil {
		return
	}
}
