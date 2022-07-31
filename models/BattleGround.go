package models

import "fmt"

type BattleGround struct {
	Size     Size
	ViewPort ViewPort
	// we store the grid as a sparse array
	Grid map[string]Entity
}

type ViewPort struct {
	X int
	Y int
	Height int
	Width int
}

func NewBattleGround(x int, y int) *BattleGround {
	var b BattleGround

	b.Size.X = x
	b.Size.Y = y

	b.Grid = make(map[string]Entity)

	b.ViewPort = ViewPort{ X: x/2, Y: y/2, Height: 20, Width: 40}

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
