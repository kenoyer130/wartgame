package models

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kenoyer130/wartgame/ui"
)

type UnitType string

const (
	LeaderUnitType UnitType = "Leader"
)

type Unit struct {
	Name           string
	UnitNumber     UnitNumber
	Movement       int
	WeaponSkill    string
	BallisticSkill string
	Strength       int
	Toughness      int
	Wounds         int
	Attackes       int
	Leadership     int
	Save           string
	Weapons        []Weapon
	Location       Location
	UnitType       UnitType
	Token          Token
}

type UnitNumber struct {
	Min int
	Max int
}

func (re Unit) GetLocation() Location {
	return re.Location
}

func (re Unit) GetEntityType() EntityType {
	return UnitEntityType
}

func (re Unit) GetToken() *ebiten.Image {
	token := ebiten.NewImage(31, 31)
	color := color.RGBA{uint8(re.Token.RGBA.R), uint8(re.Token.RGBA.G), uint8(re.Token.RGBA.B), uint8(re.Token.RGBA.A)}

	token.Fill(color)

	text.Draw(token, re.Token.ID, ui.GetFontNormalFace(), 2, 24, ui.GetTextColor())

	return token
}
