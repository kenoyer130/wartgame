package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawGrid(screen *ebiten.Image, g *Game) {
	for x := 0; x < g.BattleGround.ViewPort.Width; x++ {

		image := ebiten.NewImage((g.BattleGround.Size.X * 32), 1)
		image.Fill(color.RGBA{35, 31, 33, 1})

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, float64(x*32))

		screen.DrawImage(image, op)
	}

	for y := 0; y < g.BattleGround.ViewPort.Height; y++ {

		image := ebiten.NewImage(1, (g.BattleGround.Size.Y * 32))
		image.Fill(color.RGBA{35, 31, 33, 1})

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(y*32), 0)

		screen.DrawImage(image, op)
	}
}
