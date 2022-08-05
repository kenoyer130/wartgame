package models

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/ui"
)

type Unit struct {
	Name      string
	Army      string
	Models    []Model
	Power     int
	Location  Location
	Rect      ui.Rect
	UnitState []UnitState
}

type UnitState struct {
	Advanced bool
	FellBack bool
	Shot     bool
}

func (re Unit) CanShoot() bool {
	for _, unitState := range re.UnitState {

		if unitState.Shot || unitState.Advanced || unitState.FellBack {
			return false
		}
	}

	return true
}

func (unit Unit) DrawUnitSelected(screen *ebiten.Image) {

	if (unit.Rect == ui.Rect{}) {
		return
	}

	rect := ui.Rect{}

	rect.X = unit.Rect.X * ui.TileSize
	rect.Y = unit.Rect.Y * ui.TileSize

	rect.W = (unit.Rect.W + 1) * ui.TileSize
	rect.H = (unit.Rect.H + 1) * ui.TileSize

	ui.DrawSelectorBox(rect, screen)
}
