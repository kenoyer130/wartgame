package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Wartgame!")

	game := &Game{}
	
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
