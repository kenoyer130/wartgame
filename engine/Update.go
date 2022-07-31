package engine

import "github.com/kenoyer130/wartgame/models"

func (g *Game) Update() error {

	updateState(g, g.CurrentGameState)
	return nil
}

func updateState(g *Game, s models.GameState) {

	switch s {
	case models.Start:

	}
}
