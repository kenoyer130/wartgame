package engine

import (
	"fmt"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/models"
)

type UnitWeaponsPanel struct {
	ColWidths   map[int]int
	WeaponPanel *PanelVertical
	Unit        *models.Unit
}

func NewUnitWeaponsPanel(unit *models.Unit) *UnitWeaponsPanel {
	var p UnitWeaponsPanel
	p.ColWidths = map[int]int{
		0: 40,
		1: 100,
		2: 40,
		3: 40,
		4: 40,
		5: 40,
	}

	p.Unit = unit

	return &p
}

func (re UnitWeaponsPanel) GetUnitWeaponsPanel() *ebiten.Image {

	re.WeaponPanel = NewPanelVertical(500, 400, re.ColWidths)

	re.WeaponPanel.addTitle("Weapons")

	drawWeaponLabels(re.WeaponPanel)

	weapons := getModelWeapons(re.Unit)
	drawWeapons(weapons, re.WeaponPanel)

	return re.WeaponPanel.Img
}

func drawWeaponLabels(panel *PanelVertical) {
	var labels []string
	c := 0

	labels = append(labels, "R")
	labels = append(labels, "Ty")
	labels = append(labels, "S")
	labels = append(labels, "AP")
	labels = append(labels, "D")
	labels = append(labels, "Abilities")

	for _, label := range labels {
		panel.addLabel(label, c)
		c++
	}
}

func getModelWeapons(unit *models.Unit) map[string]*models.Weapon {
	weapons := make(map[string]*models.Weapon)

	for _, model := range unit.Models {
		for _, weapon := range model.Weapons {
			if weapons[weapon.Name] == nil {
				weapons[weapon.Name] = &weapon
			}
		}
	}
	return weapons
}

func drawWeapons(weapons map[string]*models.Weapon, weaponPanel *PanelVertical) {

	r := 2

	for _, weapon := range weapons {
		var values []string
		c := 0

		weaponType := fmt.Sprintf("%s %d%d", weapon.Type.Type, weapon.Type.Dice, weapon.Type.Number)

		values = append(values, strconv.Itoa(weapon.Range))
		values = append(values, weaponType)
		values = append(values, strconv.Itoa(weapon.Strength))
		values = append(values, strconv.Itoa(weapon.ArmorPiercing))
		values = append(values, strconv.Itoa(weapon.Damage))

		for _, value := range values {
			weaponPanel.addValue(value, r, c)
			c++
		}
	}
}
