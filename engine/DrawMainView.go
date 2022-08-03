package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/ui"
)

func DrawMainView(g *Game, background *ebiten.Image) {
	mainView := ebiten.NewImage(g.BattleGround.ViewPort.Width*ui.TileSize, g.BattleGround.ViewPort.Height*ui.TileSize)
	mainView.Fill(ui.GetBattleGroundBackgroundColor())

	if g.UIState.ShowGrid {
		DrawGrid(mainView, g)
	}

	drawSelectedModelInfo(g, mainView)

	DrawEntities(g, mainView)

	background.DrawImage(mainView, nil)
}
