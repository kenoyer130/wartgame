package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/ui"
)

func GetMainView() *ebiten.Image {
	mainView := ebiten.NewImage(models.Game().BattleGround.ViewPort.Width*ui.TileSize, models.Game().BattleGround.ViewPort.Height*ui.TileSize)
	mainView.Fill(ui.GetBattleGroundBackgroundColor())

	if models.Game().UIState.ShowGrid {
		DrawGrid(mainView)
	}

	drawSelectedModelInfo(mainView)

	DrawEntities(mainView)

	return mainView
}
