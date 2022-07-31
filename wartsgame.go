package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/initilizer"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.MaximizeWindow()
	ebiten.SetWindowTitle("Wartgame!")

	game := &engine.Game{}

	if err := initilizer.StartGame(game); err != nil {
		engine.Error(err.Error())
		return
	}

	if err := ebiten.RunGame(game); err != nil {
		engine.Error(err.Error())
		return
	}
}
