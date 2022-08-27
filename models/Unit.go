package models

import (
	"errors"
	"fmt"
	"image/color"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kenoyer130/wartgame/ui"
)

type Unit struct {
	ID                 string
	Name               string
	Army               string
	Models             []*Model
	ModelCount         map[string]int
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
	Token              Token
}

func (re *Unit) GetAssetPath(assetName string, ext string) string {
	path := fmt.Sprintf("./assets/armies/%s/images/%s.%s", re.getImgPath(re.Army), re.getImgPath(assetName), ext)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return ""
	}

	return path
}

func (re *Unit) getImgPath(path string) string {
	path = strings.Replace(path, " ", "_", -1)
	return path
}

func (re *Unit) GetLocation() Location {
	return re.Location
}

func (re *Unit) SetLocation(location Location) {
	re.Location = location
	Game().BattleGround.PlaceBattleGroundEntity(re)
}

func (re *Unit) GetPlayerIndex() int {
	return re.PlayerIndex
}

func (re *Unit) GetEntityType() EntityType {
	return UnitEntityType
}

func (re *Unit) GetToken() *ebiten.Image {

	x := re.Location.X
	y := re.Location.Y

	token := ebiten.NewImage(60, 60)
	color := color.RGBA{uint8(re.Token.RGBA.R), uint8(re.Token.RGBA.G), uint8(re.Token.RGBA.B), uint8(re.Token.RGBA.A)}

	tline := ui.DrawLine(64)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x-32), float64(y-32))

	token.DrawImage(tline, op)

	token.Fill(color)

	index := 0

	for col := 0; col < 4; col++ {
		for row := 0; row < 4; row++ {
			if len(re.Models)-1 < index {
				break
			}

			model := ebiten.NewImage(12, 12)
			color := ui.GetMoveRangeColor()

			model.Fill(color)
			
			if(re.Models[index].Wounds != re.Models[index].CurrentWounds) {
				current:= strconv.Itoa(re.Models[index].CurrentWounds)
				text.Draw(model, current, ui.GetFontTiny(), 1, 3, ui.GetTextColor())
			}

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(col*14)+3, float64(row*14)+3)

			token.DrawImage(model, op)

			index++
		}
	}

	return token
}

func (re *Unit) GetID() string {
	return re.ID
}

func (re *Unit) Remove() {
	Game().BattleGround.RemoveBattleGroundEntity(re)
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

func (re *Unit) InflictWounds(targetModel Model, str int) (bool, Model) {

	var model = re.GetModelByID(targetModel.ID)

	hp := model.CurrentWounds - str
	model.CurrentWounds = hp

	if model.CurrentWounds <= 0 {
		model.Destroyed = true
		re.ModelCount[model.Name]--
		re.removeModel(model)
		return true, *model
	}

	return false, Model{}
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

	// add model to killed list
	re.DestroyedModels = append(re.DestroyedModels, destroyedModel)

	// remove from active duty
	re.Models = append(re.Models[:index], re.Models[index+1:]...)
}
