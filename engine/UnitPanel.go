package engine

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/ui"
)

func getUnitPanel(unit *models.Unit) *ebiten.Image {

	panel := NewPanel(500, 800)

	panel.addTitle("Unit Information")

	if unit == nil {
		panel.addMessage("No Model Selected", 3)
	} else {
		drawModels(panel, unit)
		drawWeaponsPanel(panel, unit)
	}

	return panel.Img
}

func drawWeaponsPanel(panel *Panel, unit *models.Unit) {
	weaponPanel := NewUnitWeaponsPanel(unit).GetUnitWeaponsPanel()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 290)

	panel.Img.DrawImage(weaponPanel, op)
}

func drawModels(panel *Panel, unit *models.Unit) {
	drawLabels(panel)

	panel.addValue(unit.Name, 1, 0)

	if(len(unit.Models) == 0) {
		return
	}

	modelNames := make(map[string]bool)

	c := 0
	for _, model := range unit.Models {
		if !modelNames[model.Name] {
			drawModelInfo(unit.Name, model, panel, c)
			modelNames[model.Name] = true
			c++
		}
	}

	drawModelImage(unit.Models[0], unit, panel)
}

func drawLabels(panel *Panel) {

	var labels []string
	r := 2

	labels = append(labels, "Name")
	labels = append(labels, "Movement")
	labels = append(labels, "Weapon Skill")
	labels = append(labels, "Ballistic Skill")
	labels = append(labels, "Strength")
	labels = append(labels, "Toughness")
	labels = append(labels, "Wounds")
	labels = append(labels, "Attacks")
	labels = append(labels, "Leadership")
	labels = append(labels, "Save")

	for _, label := range labels {
		panel.addLabel(label, r)
		r++
	}
}

func drawModelInfo(unit string, model models.Model, panel *Panel, c int) {

	var values []string
	r := 2

	values = append(values, model.Token.ID)
	values = append(values, strconv.Itoa(model.Movement))
	values = append(values, model.WeaponSkill)
	values = append(values, model.BallisticSkill)
	values = append(values, strconv.Itoa(model.Strength))
	values = append(values, strconv.Itoa(model.Toughness))
	values = append(values, strconv.Itoa(model.Wounds))
	values = append(values, strconv.Itoa(model.Attacks))
	values = append(values, strconv.Itoa(model.Leadership))
	values = append(values, model.Save)

	for _, value := range values {
		panel.addValue(value, r, c)
		r++
	}
}

func drawModelImage(model models.Model, unit *models.Unit, panel *Panel) {
	ModelImg := getProfilePic(model, unit)

	if ModelImg == nil {
		return
	}

	text.Draw(panel.Img, model.Name, ui.GetFontItalic(), 275, 74, ui.GetTextColor())

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(275, 84)

	panel.Img.DrawImage(ModelImg, op)
}

var ModelPics = make(map[string]*ebiten.Image)

func getProfilePic(model models.Model, unit *models.Unit) *ebiten.Image {

	imgPath := getImgPath(model.Name)

	if ModelPics[imgPath] != nil {
		return ModelPics[imgPath]
	}

	path := fmt.Sprintf("./assets/armies/%s/images/%s.png", getImgPath(unit.Army), imgPath)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil
	}

	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	ModelPics[imgPath] = img

	return img
}

func getImgPath(path string) string {
	path = strings.Replace(path, " ", "_", -1)
	return path
}
