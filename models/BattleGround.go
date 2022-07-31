package models

import "fmt"

type BattleGround struct {
	Size     Size
	ViewPort Location
	// we store the grid as a sparse array
	grid map[string]Entity
}

func NewBattleGround(x int, y int) *BattleGround {
	var b BattleGround

	b.Size.x = x
	b.Size.y = y

	b.grid = make(map[string]Entity)

	b.ViewPort = Location{ }

	return &b
}

type Size struct {
	y int
	x int
}

// put an entity on the battleground thus taking up space
func PlaceBattleGroundEntity(entity Entity, battleGround *BattleGround) {
	locationKey := getLocationKey(entity.GetLocation())
	battleGround.grid[locationKey] = entity
}

// check if a location is empty
func IsBattleGroundLocationFree(location Location, battleGround *BattleGround) bool {
	locationKey := getLocationKey(location)
	_, ok := battleGround.grid[locationKey]

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
