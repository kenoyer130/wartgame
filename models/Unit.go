package models

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/ui"
)

type Unit struct {
	Name               string
	Army               string
	Models             []Model
	DestroyedModels    []Model
	Power              int
	Location           Location
	Rect               ui.Rect
	UnitState          []UnitState
	Destroyed          bool
	OriginalModelCount int
}

type UnitState string

const (
	UnitAdvanced UnitState = "UnitAdvanced"
	UnitFellBack UnitState = "UnitFellBack"
	UnitShot     UnitState = "UnitShot"
)

func (re *Unit) AddState(state UnitState) {
	re.UnitState = append(re.UnitState, state)
}

func (re *Unit) ClearStates() {
	re.UnitState = []UnitState{}
}

func (re Unit) GetMoraleCheck() int {

	leadership := 0

	for _, model := range re.Models {
		leadership = int(math.Max(float64(leadership), float64(model.Leadership)))
	}

	return leadership
}

func (re Unit) CanShoot() bool {
	for _, unitState := range re.UnitState {

		if unitState == UnitShot || unitState == UnitAdvanced || unitState == UnitFellBack {
			return false
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
		re.removeModel(target)
	}
}

func (re *Unit) MoraleFailure() {

	if len(re.Models) == 0 {
		return
	}

	model := 0

	re.removeModel(model)
}

func (re *Unit) removeModel(index int) {
	destroyedModel := re.Models[index]

	// add model to killed list
	re.DestroyedModels = append(re.DestroyedModels, destroyedModel)

	// remove from active duty
	re.Models = append(re.Models[:index], re.Models[index+1:]...)

	// remove from map
	Game().BattleGround.RemoveEntity(destroyedModel.ID)
}
