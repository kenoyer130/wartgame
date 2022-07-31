package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {

	image := ebiten.NewImage(16, 16)
	image.Fill(color.RGBA{245, 40, 145, 1 })
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(image, op)
}
