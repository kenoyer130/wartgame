package models

import (
	"strconv"
	"strings"
)

type ModelType string

const (
	LeaderModelType ModelType = "Leader"
)

type Model struct {
	ID             string
	Name           string
	ShortName      string
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
	ModelType      ModelType
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
