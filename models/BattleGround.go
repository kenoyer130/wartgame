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

	b.Grid = New2DArray[Entity](x, y)

	b.ViewPort = ViewPort{X: 0, Y: 0, Height: 14, Width: 23}

	return &b
}

type Size struct {
	Y int
	X int
}

// put an entity on the battleground thus taking up space
func (re *BattleGround) PlaceBattleGroundEntity(entity Entity) {
	l := entity.GetLocation()
	re.Grid[l.X][l.Y] = entity
}

// put an entity on the battleground thus taking up space
func (re *BattleGround) RemoveBattleGroundEntity(entity Entity) {
	l := entity.GetLocation()
	re.Grid[l.X][l.Y] = nil
}

// put an entity on the battleground thus taking up space
func (re *BattleGround) UpdateBattleGroundEntity(entity Entity) {

	l := entity.GetLocation()
	re.Grid[l.X][l.Y] = entity

}

// return entity at location if any
func (re *BattleGround) GetEntityAtLocation(l Location) Entity {

	if (l.X < 0) || (l.Y < 0) {
		return nil
	}

	if (l.X > len(Game().BattleGround.Grid)-1) || (l.Y > len(Game().BattleGround.Grid[l.X])-1) {
		return nil
	}

	if re.Grid[l.X][l.Y] != nil {
		return re.Grid[l.X][l.Y]
	} else {
		return nil
	}
}

// check if a location is empty
func (re *BattleGround) IsBattleGroundLocationFree(l Location) bool {
	if re.Grid[l.X][l.Y] != nil {
		return false
	} else {
		return true
	}
}

func (re *BattleGround) IsBattleGroundLocationRectFree(x int, y int, w int, h int) bool {

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {

			if (x+i < 0 || x+i > len(re.Grid)) || (x+i < 0 || y+j > len(re.Grid[i])) {
				return false
			}

			if re.Grid[x+i][y+j] != nil {
				return false
			}
		}
	}

	return true
}
