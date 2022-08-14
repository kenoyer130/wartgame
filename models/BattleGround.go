package models

import (
	"github.com/kenoyer130/wartgame/ui"
)

type BattleGround struct {
	Size     Size
	ViewPort ViewPort
	Grid     [][]Entity
}

type ViewPort struct {
	X      int
	Y      int
	Height int
	Width  int
}

func (re *BattleGround) RemoveEntity(l Location) {
	re.Grid[l.X][l.Y] = nil
}

func (re ViewPort) GetPixelRectangle() Rectangle {
	return Rectangle{X: re.X * ui.TileSize, Y: re.Y * ui.TileSize, Width: re.Width * ui.TileSize, Height: re.Height * ui.TileSize}
}

func NewBattleGround(x int, y int) *BattleGround {
	var b BattleGround

	b.Size.X = x
	b.Size.Y = y

	b.Grid = make([][]Entity, x)

	for i := range b.Grid {
		b.Grid[i] = make([]Entity, y)
	}

	b.ViewPort = ViewPort{X: 0, Y: 0, Height: 28, Width: 45}

	return &b
}

type Size struct {
	Y int
	X int
}

// put an entity on the battleground thus taking up space
func PlaceBattleGroundEntity(entity Entity, battleGround *BattleGround) {
	l := entity.GetLocation()
	battleGround.Grid[l.X][l.Y] = entity
}

// put an entity on the battleground thus taking up space
func RemoveBattleGroundEntity(entity Entity, battleGround *BattleGround) {
	l := entity.GetLocation()
	battleGround.Grid[l.X][l.Y] = nil
}

// put an entity on the battleground thus taking up space
func UpdateBattleGroundEntity(entity Entity, battleGround *BattleGround) {

	l := entity.GetLocation()
	battleGround.Grid[l.X][l.Y] = entity

}

// check if a location is empty
func IsBattleGroundLocationFree(l Location, battleGround *BattleGround) bool {
	if battleGround.Grid[l.X][l.Y] != nil {
		return false
	} else {
		return true
	}
}
