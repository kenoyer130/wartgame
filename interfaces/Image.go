package interfaces

import "github.com/hajimehoshi/ebiten/v2"

type Draw interface {
	DrawImage(img *ebiten.Image, options *ebiten.DrawImageOptions)
}

type Drawer struct {
	*ebiten.Image
}