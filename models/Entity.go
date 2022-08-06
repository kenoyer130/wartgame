package models

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type EntityType int64

const (
	ModelEntityType EntityType = 0
)

type Entity interface {
	GetLocation() Location
	GetEntityType() EntityType
	GetToken() *ebiten.Image
	GetID() string
}