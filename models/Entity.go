package models

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type EntityType int64

const (
	UnitEntityType EntityType = 0
)

type Entity interface {
	GetLocation() Location
	GetEntityType() EntityType
	GetToken() *ebiten.Image
}