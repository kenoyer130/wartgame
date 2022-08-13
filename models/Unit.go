package models

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/ui"
)

type Unit struct {
	ID                 string
	Name               string
	Army               string
	Models             []*Model
	DestroyedModels    []*Model
	Power              int
	Location           Location
	Rect               ui.Rect
	UnitState          []UnitState
	Destroyed          bool
	OriginalModelCount int
	CurrentMoves       int
	PlayerIndex        int
	Width              int
	Height             int
	MovementRect       ui.Rect
	RangeRect          ui.Rect
}

func (re *Unit) Place() {
	size := re.findLargest(len(re.Models))

	index := 0

	row := 0
	col := 0

	for index < len(re.Models) {

		model := re.Models[index]
		RemoveBattleGroundEntity(model, &Game().BattleGround)

		model.Location = Location{re.Location.X + row, re.Location.Y + col}

		PlaceBattleGroundEntity(model, &Game().BattleGround)

		Game().GameStateUpdater.UpdateModel(model.PlayerIndex, model)

		re.Width = int(math.Max(float64(re.Width), float64(col)))
		re.Height = int(math.Max(float64(re.Height), float64(row)))

		col++

		if col > size {
			col = 0
			row++
		}

		index++

	}

	// adjust height
	re.Height++

	re.setMovementRect()
}

func (re *Unit) setMovementRect() {
	m := re.Models[0].Movement

	x := re.Location.X - m
	w := (m * 2) + re.Width

	y := re.Location.Y - m
	h := (m * 2) + re.Height

	re.MovementRect = ui.Rect{X: x, Y: y, W: w, H: h}
}

func (re Unit) findLargest(x int) int {
	return int(math.Sqrt(float64(x)))
}

func (re *Unit) Cleanup() {
	re.UnitState = []UnitState{}
	re.DestroyedModels = []*Model{}

	for i := 0; i < len(re.Models); i++ {
		re.Models[i].ClearFiredWeapon()
	}
}

type UnitState string

const (
	UnitAdvanced UnitState = "UnitAdvanced"
	UnitFellBack UnitState = "UnitFellBack"
	UnitShot     UnitState = "UnitShot"
	UnitMoved    UnitState = "UnitMoved"
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

func (re Unit) CanMove() bool {
	for _, unitState := range re.UnitState {

		if unitState == UnitMoved {
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

func (re *Unit) InflictWounds(targetModel Model, str int) bool {

	var model = re.GetModelByID(targetModel.ID)

	hp := model.CurrentWounds - str
	model.CurrentWounds = hp

	if model.CurrentWounds <= 0 {
		re.removeModel(model)
		return true
	} else {
		UpdateBattleGroundEntity(model, &Game().BattleGround)
		return false
	}
}

func (re *Unit) GetModelByID(id string) *Model {
	for _, model := range re.Models {
		if model.ID == id {
			return model
		}
	}

	return nil
}

func (re *Unit) GetModelIndexByID(id string) int {
	for i, model := range re.Models {
		if model.ID == id {
			return i
		}
	}

	return -1
}

func (re *Unit) GetDestroyedModelByID(id string) *Model {
	for _, model := range re.DestroyedModels {
		if model.ID == id {
			return model
		}
	}

	return nil
}

func (re *Unit) MoraleFailure() {

	if len(re.Models) == 0 {
		return
	}

	re.removeModel(re.Models[0])
}

func (re *Unit) removeModel(destroyedModel *Model) {

	index := re.GetModelIndexByID(destroyedModel.ID)

	// remove from map
	Game().BattleGround.RemoveEntity(destroyedModel.ID)

	// add model to killed list
	re.DestroyedModels = append(re.DestroyedModels, destroyedModel)

	// remove from active duty
	re.Models = append(re.Models[:index], re.Models[index+1:]...)
}
