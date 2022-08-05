package models

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kenoyer130/wartgame/ui"
)

type ModelType string

const (
	LeaderModelType ModelType = "Leader"
)

type Model struct {
	Name           string
	Count          int
	ModelNumber    ModelNumber
	Movement       int
	WeaponSkill    string
	BallisticSkill string
	Strength       int
	Toughness      int
	Wounds         int
	Attacks        int
	Leadership     int
	Save           string
	Weapons        []string
	Location       Location
	ModelType      ModelType
	Token          Token
}

type ModelNumber struct {
	Min int
	Max int
}

func (re Model) GetLocation() Location {
	return re.Location
}

func (re Model) GetEntityType() EntityType {
	return ModelEntityType
}

func (re Model) GetToken() *ebiten.Image {
	token := ebiten.NewImage(31, 31)
	color := color.RGBA{uint8(re.Token.RGBA.R), uint8(re.Token.RGBA.G), uint8(re.Token.RGBA.B), uint8(re.Token.RGBA.A)}

	token.Fill(color)

	text.Draw(token, re.Token.ID, ui.GetFontNormalFace(), 2, 24, ui.GetTextColor())

	return token
}
