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
	PhaseStepper        PhaseStepper
	DiceRoller          DiceRoller
	Players             [consts.MaxPlayers]Player
	CurrentPlayer       *Player
	CurrentPlayerIndex  int
	StartPlayerIndex    int
	SelectedUnit        *Unit
	DraggingUnit        *Unit
	SelectedPhaseUnit   *Unit
	SelectedTargetUnit  *Unit
	SelectedModel       *Model
	SelectedWeaponName  string
	SelectedWeaponIndex int
	SelectedWeapons     []Weapon
	Assets              Assets
	UIState             UIState
	StatusMessage       StatusMessage
	Dice                []int
	GameStateUpdater    GameStateUpdater
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
