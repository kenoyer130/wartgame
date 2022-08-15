package models

import (
	"fmt"
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
	UnitState          []UnitState
	Destroyed          bool
	OriginalModelCount int
	CurrentMoves       int
	PlayerIndex        int
	Width              int
	Height             int
	Rect               ui.Rect
	MovementRect       ui.Rect
	MovementRange      map[string]Location
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
	re.Height = re.Height + 2

	re.Rect = ui.Rect{X: re.Location.X, Y: re.Location.Y, W: re.Width, H: re.Height}

	re.setMovementRect()
}

func (re *Unit) Remove() {
	for i := 0; i < len(re.Models); i++ {
		model := re.Models[i]
		RemoveBattleGroundEntity(model, &Game().BattleGround)
	}
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

	if (re.Location == Location{}) {
		return
	}

	rect := ui.Rect{}

	rect.X = re.Location.X * ui.TileSize
	rect.Y = re.Location.Y * ui.TileSize

	rect.W = (re.Width + 1) * ui.TileSize
	rect.H = (re.Height + 1) * ui.TileSize

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
	Game().BattleGround.RemoveEntity(destroyedModel.Location)

	// add model to killed list
	re.DestroyedModels = append(re.DestroyedModels, destroyedModel)

	// remove from active duty
	re.Models = append(re.Models[:index], re.Models[index+1:]...)
}

func (re *Unit) SetMoveRange() {

	// remove unit temporarily so it doesn't conflict with itself
	re.Remove()

	movement := re.Models[0].Movement + 1

	movementRange := make(map[string]Location)

	floodfill(movementRange, re.Location.X, re.Location.Y, re.Width, re.Height, movement)

	re.Place()

	for k, l := range movementRange {
		if !IsBattleGroundLocationFree(l, &Game().BattleGround) {
			delete(movementRange, k)
		}
	}

	re.MovementRange = movementRange
}

func floodfill(tiles map[string]Location, x int, y int, w int, h int, moves int) {

	if !IsBattleGroundLocationRectFree(x, y, w, h, &Game().BattleGround) {
		return
	}

	for c := 0; c < w; c++ {
		for r := 0; r < h; r++ {
			l := Location{X: x + c, Y: y + r}

			key := fmt.Sprintf("%d_%d", l.X, l.Y)

			tiles[key] = l
		}
	}

	moves = moves - 1

	if moves == 0 {
		return
	}

	// check up
	floodfill(tiles, x, y-1, w, h, moves)

	floodfill(tiles, x-1, y-1, w, h, moves)
	floodfill(tiles, x+1, y+1, w, h, moves)
	floodfill(tiles, x-1, y+1, w, h, moves)
	floodfill(tiles, x+1, y-1, w, h, moves)

	// check right
	floodfill(tiles, x+1, y, w, h, moves)

	// check down
	floodfill(tiles, x, y+1, w, h, moves)

	// check left
	floodfill(tiles, x-1, y, w, h, moves)
}
