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
	DefaultWeapons []string
	Weapons        []Weapon
	SelectedWeapon Weapon
	Location       Location
	ModelType      ModelType
	Token          Token
	PlayerIndex    int
}

func (re Model) GetID() string {
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

func (re Model) GetUnfiredWeapon() *Weapon {

	isUnthrownGernade, grenadeWeapon := re.checkGrenade()
	if isUnthrownGernade {
		return grenadeWeapon
	}

	shouldReturn, returnValue := re.getUnfired()
	if shouldReturn {
		return returnValue
	}

	return nil
}

func (re Model) checkGrenade() (bool, *Weapon) {
	for _, weapon := range re.Weapons {

		if weapon.WeaponType.Type == "Gre" && !weapon.Fired {
			return true, &weapon
		}
	}

	return false, nil
}

func (re Model) getUnfired() (bool, *Weapon) {
	for _, weapon := range re.Weapons {

		if !weapon.Fired {
			return true, &weapon
		}
	}
	return false, nil
}

func (re *Model) SetFiredWeapon(weapon *Weapon) {

	// for gernades the unit cannot fire other weapons
	shouldReturn := setGernadeFired(weapon, re)
	if shouldReturn {
		return
	}
	weapon.Fired = true
}

func setGernadeFired(weapon *Weapon, re *Model) bool {
	if weapon.WeaponType.Type == "Gre" {
		weapon.Fired = true
		return true
	}
	return false
}

func (re *Model) ClearFiredWeapon() {
	for i := 0; i < len(re.Weapons); i++ {
		re.Weapons[i].Fired = false
	}
}

func (re Model) GetLocation() Location {
	return re.Location
}

func (re *Model) SetLocation(location Location) {
	re.Location = location
}

func (re Model) GetEntityType() EntityType {
	return ModelEntityType
}

func (re Model) GetToken() *ebiten.Image {
	token := ebiten.NewImage(31, 31)
	color := color.RGBA{uint8(re.Token.RGBA.R), uint8(re.Token.RGBA.G), uint8(re.Token.RGBA.B), uint8(re.Token.RGBA.A)}

	token.Fill(color)

	text.Draw(token, re.Token.ID, ui.GetFontNormalFace(), 2, 24, ui.GetTextColor())

	if re.Wounds != re.CurrentWounds {
		wounds := re.Wounds - re.CurrentWounds
		text.Draw(token, strconv.Itoa(wounds), ui.GetFontTiny(), 2, 10, ui.GetWoundColor())
	}

	return token
}
