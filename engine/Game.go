package engine

import (
	"github.com/kenoyer130/wartgame/models"
)

type Game struct {
	CurrentGameState models.GameState
	Round            int
	BattleGround     models.BattleGround
	CurrentPhase     models.GamePhase	
	Players          [models.MaxPlayers]models.Player
	CurrentPlayer    *models.Player
	SelectedUnit    *models.Unit
	SelectedModel     *models.Model
	Assets           models.Assets
	UIState          UIState
}

type UIState struct {
	ShowGrid     bool
	GridDragging DraggingGrid
}

type DraggingGrid struct {
	InDrag       bool
	CursorStartX int
	CursorStartY int
}
