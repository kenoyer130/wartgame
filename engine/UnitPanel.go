package engine

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)


func getUnitPanel(g *Game) *ebiten.Image {

	panel := NewPanel(400, 800)

	panel.addTitle("Unit Information")

	if g.SelectedUnit == nil {
		panel.addRow("No Unit Selected", "")
	} else {
		panel.addRow("Squad:", g.SelectedSquad.Name)
		panel.addRow("Name:", g.SelectedUnit.Name)
		panel.addRow("Movement:",  strconv.Itoa(g.SelectedUnit.Movement))
		panel.addRow("Weapon Skill:", g.SelectedUnit.WeaponSkill)
		panel.addRow("Ballistic Skill:", g.SelectedUnit.BallisticSkill)
		panel.addRow("Strength:", strconv.Itoa(g.SelectedUnit.Strength))
		panel.addRow("Toughness:", strconv.Itoa(g.SelectedUnit.Toughness))
		panel.addRow("Wounds:", strconv.Itoa(g.SelectedUnit.Wounds))
		panel.addRow("Attacks:", strconv.Itoa(g.SelectedUnit.Attacks))
		panel.addRow("Leadership:", strconv.Itoa(g.SelectedUnit.Leadership))
		panel.addRow("Save:", g.SelectedUnit.Save)
	}

	return panel.img
}

