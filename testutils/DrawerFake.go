package testutils

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type DrawerFake struct {
	tiles [][]int
}

func (re DrawerFake) DrawImage(img *ebiten.Image, options *ebiten.DrawImageOptions) {

}
