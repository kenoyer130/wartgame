package engine

import (
	"fmt"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/models"
)

type UnitWeaponsPanel struct {
	WeaponPanel *PanelVertical
	Unit        *models.Unit
}

var ColWidths = map[int]int{
	0: 20,
	1: 150,
	2: 30,
	3: 30,
	4: 60,
	5: 30,
	6: 30,
	7: 30,
	8: 30,
}

func NewUnitWeaponsPanel(unit *models.Unit) *UnitWeaponsPanel {
	var p UnitWeaponsPanel
	p.Unit = unit
	return &p
}

func (re UnitWeaponsPanel) GetUnitWeaponsPanel(g *Game) *ebiten.Image {

	re.WeaponPanel = NewPanelVertical(500, 400)

	re.WeaponPanel.addTitle("Weapons")

	re.drawWeaponLabels(re.WeaponPanel)

	re.drawWeapons(g, re.WeaponPanel)

	return re.WeaponPanel.Img
}

func (re UnitWeaponsPanel) drawWeaponLabels(panel *PanelVertical) {
	var labels []string

	labels = append(labels, " ")
	labels = append(labels, "Name")
	labels = append(labels, "#")
	labels = append(labels, "R")
	labels = append(labels, "Type")
	labels = append(labels, "S")
	labels = append(labels, "AP")
	labels = append(labels, "D")
	labels = append(labels, "Abilities")

	totalWidth := 0

	for i, label := range labels {
		panel.addLabel(label, 2, i, totalWidth)
		totalWidth = totalWidth + ColWidths[i]
	}
}

func (re UnitWeaponsPanel) drawWeapons(g *Game, weaponPanel *PanelVertical) {

	r := 3

	assetWeapons := make(map[string]bool)
	weaponCount := make(map[string]int)

	for _, model := range re.Unit.Models {

		for _, weapon := range model.Weapons {
			total := weaponCount[weapon]
			total++
			weaponCount[weapon] = total
		}
	}

	for _, model := range re.Unit.Models {

		for _, weapon := range model.Weapons {
			if !assetWeapons[weapon] {

				var values []string

				thisWeapon := g.Assets.Weapons[weapon]

				weaponType := fmt.Sprintf("%s %d%d", thisWeapon.WeaponType.Type, thisWeapon.WeaponType.Dice, thisWeapon.WeaponType.Number)

				selected := ""

				if(thisWeapon.Name == g.SelectedWeapon) {
					selected = "X"
				}

				values = append(values,selected)
				values = append(values, thisWeapon.Name)
				values = append(values, strconv.Itoa(weaponCount[thisWeapon.Name]))
				values = append(values, strconv.Itoa(thisWeapon.Range))
				values = append(values, weaponType)
				values = append(values, strconv.Itoa(thisWeapon.Strength))
				values = append(values, strconv.Itoa(thisWeapon.ArmorPiercing))
				values = append(values, strconv.Itoa(thisWeapon.Damage))

				totalWidth := 0

				for i, value := range values {
					weaponPanel.addValue(value, r, i, totalWidth)
					totalWidth = totalWidth + ColWidths[i]
				}

				r++

				assetWeapons[weapon] = true
			}
		}
	}
}
