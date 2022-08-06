package models

import (
	"image/color"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kenoyer130/wartgame/ui"
)

type ModelType string

const (
	LeaderModelType ModelType = "Leader"
)

type Model struct {
	ID             string
	Name           string
	Count          int
	ModelNumber    ModelNumber
	Movement       int
	WeaponSkill    string
	BallisticSkill string
	Strength       int
	Toughness      int
	CurrentWounds  int
	Wounds         int
	Attacks        int
	Leadership     int
	Save           string
	Weapons        []string
	FiredWeapons   []string
	SelectedWeapon string
	Location       Location
	ModelType      ModelType
	Token          Token
}

func (re Model)GetID() string {
	return re.ID
}

type ModelNumber struct {
	Min int
	Max int
}

func (re Model) GetBallisticSkill() int {
	n, _ := strconv.Atoi(strings.Replace(re.BallisticSkill, "+", "", -1))
	return n
}

func (re Model) GetIntSkill(value string) int {
	n, _ := strconv.Atoi(strings.Replace(value, "+", "", -1))
	return n
}

func (re Model) GetUnfiredWeapon() string {
	for _, weapon := range re.Weapons {

		hasFired := false

		for _, fired := range re.FiredWeapons {
			if weapon == fired {
				hasFired = true
				break
			}
		}

		if !hasFired {
			return weapon
		}
	}

	return ""
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

	if(re.Wounds != re.CurrentWounds) {
		wounds := re.Wounds - re.CurrentWounds
		text.Draw(token, strconv.Itoa(wounds), ui.GetFontTiny(), 2, 24, ui.GetWoundColor())
	}

	return token
}
