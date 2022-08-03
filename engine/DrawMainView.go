package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/ui"
)

func DrawMainView(g *Game, background *ebiten.Image) {
	mainView := ebiten.NewImage(g.BattleGround.ViewPort.Width*ui.TileSize, g.BattleGround.ViewPort.Height*ui.TileSize)
	mainView.Fill(color.RGBA{166, 142, 154, 1})

	if g.UIState.ShowGrid {
		DrawGrid(background, g)
	}

	drawSelectedModelInfo(g, mainView)

	DrawEntities(g, mainView)

	background.DrawImage(mainView, nil)
}