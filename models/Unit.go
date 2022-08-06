package models

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/ui"
)

type Unit struct {
	Name         string
	Army         string
	Models       []Model
	KilledModels []Model
	Power        int
	Location     Location
	Rect         ui.Rect
	UnitState    []UnitState
}

type UnitState bool

const (
	UnitAdvanced UnitState = false
	UnitFellBack UnitState = false
	UnitShot     UnitState = false
)

func (re *Unit) AddState(state UnitState) {
	re.UnitState = append(re.UnitState, state)
}

func (re *Unit) ClearStates() {
	re.UnitState = []UnitState{}
}

func (re Unit) CanShoot() bool {
	for _, unitState := range re.UnitState {

		if unitState == UnitShot || unitState == UnitAdvanced || unitState == UnitFellBack {
			if unitState {
				return false
			}
		}
	}

	return true
}

func (re Unit) DrawUnitSelected(screen *ebiten.Image) {

	if (re.Rect == ui.Rect{}) {
		return
	}

	rect := ui.Rect{}

	rect.X = re.Rect.X * ui.TileSize
	rect.Y = re.Rect.Y * ui.TileSize

	rect.W = (re.Rect.W + 1) * ui.TileSize
	rect.H = (re.Rect.H + 1) * ui.TileSize

	ui.DrawSelectorBox(rect, screen)
}

func (re *Unit) InflictWounds(target int, str int) {

	model := &re.Models[target]

	hp := re.Models[target].Wounds - str
	model.CurrentWounds = hp

	if model.CurrentWounds <= 0 {
		killedModel := re.Models[target]

		// add model to killed list
		re.KilledModels = append(re.KilledModels, killedModel)

		// remove from active duty
		re.Models = append(re.Models[:target], re.Models[target+1:]...)

	}
}
