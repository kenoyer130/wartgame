package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kenoyer130/wartgame/ui"
)

type BattleGround struct {
	Size     Size
	ViewPort ViewPort
	// we store the grid as a sparse array
	Grid map[string]Entity
}

type ViewPort struct {
	X      int
	Y      int
	Height int
	Width  int
}

func (re *BattleGround) RemoveEntity(id string) {
	for key, entity := range re.Grid {
		if entity.GetID() == id {
			delete(re.Grid, key)
		}
	}
}

func (re ViewPort) GetPixelRectangle() Rectangle {
	return Rectangle{X: re.X * ui.TileSize, Y: re.Y * ui.TileSize, Width: re.Width * ui.TileSize, Height: re.Height * ui.TileSize}
}

func NewBattleGround(x int, y int) *BattleGround {
	var b BattleGround

	b.Size.X = x
	b.Size.Y = y

	b.Grid = make(map[string]Entity)

	b.ViewPort = ViewPort{X: 0, Y: 0, Height: 28, Width: 45}

	return &b
}

type Size struct {
	Y int
	X int
}

// put an entity on the battleground thus taking up space
func PlaceBattleGroundEntity(entity Entity, battleGround *BattleGround) {
	locationKey := getLocationKey(entity.GetLocation())
	battleGround.Grid[locationKey] = entity
}

// put an entity on the battleground thus taking up space
func RemoveBattleGroundEntity(entity Entity, battleGround *BattleGround) {
	locationKey := getLocationKey(entity.GetLocation())
	delete(battleGround.Grid, locationKey)
}

// put an entity on the battleground thus taking up space
func UpdateBattleGroundEntity(entity Entity, battleGround *BattleGround) {

	for key, value := range battleGround.Grid {
		if value.GetID() == entity.GetID() {
			entity.SetLocation(getLocationFromKey(key))
			battleGround.Grid[key] = entity
			break
		}
	}
}

// check if a location is empty
func IsBattleGroundLocationFree(location Location, battleGround *BattleGround) bool {
	locationKey := getLocationKey(location)
	_, ok := battleGround.Grid[locationKey]

	if ok {
		return false
	} else {
		return true
	}
}

// create a string key from the location x,y
func getLocationKey(location Location) string {
	return fmt.Sprintf("%d_%d", location.X, location.Y)
}

func getLocationFromKey(key string) Location {
	values := strings.Split(key, "_")

	x, _ := strconv.Atoi(values[0])
	y, _ := strconv.Atoi(values[1])

	return Location{
		X: x,
		Y: y,
	}
}
