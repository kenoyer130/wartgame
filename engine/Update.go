package engine

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kenoyer130/wartgame/models"
)

func (g *Game) Update() error {

	updateState(g, g.CurrentGameState)
	return nil
}

func updateState(g *Game, s models.GameState) {

	if(inpututil.IsKeyJustReleased(ebiten.KeyEscape)) {
		os.Exit(0)
	}

	if(inpututil.IsKeyJustReleased(ebiten.KeyG)) {
		if(g.ShowGrid) {
			g.ShowGrid = false
		} else {
			g.ShowGrid = true
		}
	}

	switch s {
	case models.Start:

	}
}
