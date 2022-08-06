package models

import (
	"github.com/kenoyer130/wartgame/consts"
)

var gameState GameState 

func InitGameState() {
	if !gameState.Initialized {
		gameState := GameState{}
		gameState.Initialized = true
	}
}

func Game() *GameState {
	return &gameState
}

type GameState struct {
	Initialized         bool
	CurrentGameState    GameStates
	Round               int
	BattleGround        BattleGround
	CurrentPhase        GamePhase
	Players             [consts.MaxPlayers]Player
	CurrentPlayer       *Player
	CurrentPlayerIndex  int
	SelectedUnit        *Unit
	SelectedPhaseUnit   *Unit
	SelectedTargetUnit  *Unit
	SelectedModel       *Model
	SelectedWeaponName  string
	SelectedWeaponIndex int
	SelectedWeapons     []Weapon
	Assets              Assets
	UIState             UIState
	StatusMesssage      string
	Dice                []int
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
