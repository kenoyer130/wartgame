package engine

import (
	"github.com/kenoyer130/wartgame/consts"
	"github.com/kenoyer130/wartgame/models"
)

type Game struct {
	CurrentGameState   models.GameState
	Round              int
	BattleGround       models.BattleGround
	CurrentPhase       models.GamePhase
	Players            [consts.MaxPlayers]models.Player
	CurrentPlayer      *models.Player
	CurrentPlayerIndex int
	SelectedUnit       *models.Unit
	SelectedPhaseUnit  *models.Unit
	SelectedTargetUnit *models.Unit
	SelectedModel      *models.Model
	Assets             models.Assets
	UIState            UIState
	StatusMesssage     string
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
