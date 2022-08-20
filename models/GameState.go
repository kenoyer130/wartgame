package models

import (
	"github.com/kenoyer130/wartgame/consts"
	"github.com/kenoyer130/wartgame/interfaces"
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
	Initialized          bool
	CurrentGameState     GameStates
	Round                int
	BattleGround         BattleGround
	PhaseStepper         interfaces.PhaseStepper
	DiceRoller           DiceRoller
	Players              [consts.MaxPlayers]Player
	CurrentPlayer        *Player
	CurrentPlayerIndex   int
	StartPlayerIndex     int
	SelectedUnit         *Unit
	DraggingUnit         *Unit
	DraggingUnitLocation Location
	SelectedPhaseUnit    *Unit
	SelectedTargetUnit   *Unit
	SelectedModel        *Model
	SelectedWeapon       *ShootingWeapon		
	Assets               Assets
	UIState              UIState
	StatusMessage        StatusMessage
	Dice                 []int
	GameStateUpdater     GameStateUpdater
	Drawer               interfaces.Draw
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
