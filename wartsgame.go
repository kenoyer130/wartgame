package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/initilizer"
	"github.com/kenoyer130/wartgame/models"
)

func main() {

	models.InitGameState()
	
	setMainWindow()	

	if err := initilizer.StartGame(); err != nil {
		engine.Error(err.Error())
		return
	}

	gameEngine := engine.GameEngine{}

	if err := ebiten.RunGame(gameEngine); err != nil {
		engine.Error(err.Error())
		return
	}
}

func setMainWindow() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.MaximizeWindow()
	ebiten.SetWindowTitle("Wartgame!")
}
