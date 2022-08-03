package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawSelectorBox(rect Rect, screen *ebiten.Image) {
	// draw top
	drawSelectorLine(rect.X, rect.Y, rect.W, 1, screen)

	// draw left
	drawSelectorLine(rect.X, rect.Y, 1, rect.H, screen)

	// draw right
	drawSelectorLine(rect.X+rect.W, rect.Y, 1, rect.H, screen)

	// draw bottom
	drawSelectorLine(rect.X, rect.Y+rect.H, rect.W, 1, screen)
}

func drawSelectorLine(x int, y int, w int, h int, screen *ebiten.Image) {
	unitOutline := ebiten.NewImage(w, h)

	unitOutline.Fill(color.RGBA{255, 255, 255, 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(unitOutline, op)
}