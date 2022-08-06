package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameEngine struct {
}

func (re GameEngine) Update() error {
	return UpdateGameEngine()
}

func (re GameEngine) Draw(screen *ebiten.Image) {
	GameEngineDraw(screen)
}

func (re GameEngine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}