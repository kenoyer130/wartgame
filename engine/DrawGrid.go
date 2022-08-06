package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/ui"
)

func DrawGrid(screen *ebiten.Image) {
	drawXLines(screen)
	drawYLines(screen)
}

func drawXLines(screen *ebiten.Image) {
	for x := 0; x < models.Game().BattleGround.ViewPort.Height; x++ {

		image := ebiten.NewImage(screen.Bounds().Dx(), 1)
		image.Fill(ui.GetBattleGroundBackgroundColor())

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, float64(x*ui.TileSize))

		screen.DrawImage(image, op)
	}
}

func drawYLines(screen *ebiten.Image) {
	for y := 0; y < models.Game().BattleGround.ViewPort.Width; y++ {

		image := ebiten.NewImage(1, screen.Bounds().Dy())
		image.Fill(ui.GetBattleGroundBackgroundColor())

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(y*ui.TileSize), 0)

		screen.DrawImage(image, op)
	}
}
