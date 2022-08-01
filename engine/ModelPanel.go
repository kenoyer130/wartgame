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
)

func getModelPanel(g *Game) *ebiten.Image {

	panel := NewPanel(400, 800)

	panel.addTitle("Model Information")

	if g.SelectedModel == nil {
		panel.addRow("No Model Selected", "")
	} else {
		panel.addRow("Unit:", g.SelectedUnit.Name)
		panel.addRow("Model:", g.SelectedModel.Name)
		panel.addRow("Movement:", strconv.Itoa(g.SelectedModel.Movement))
		panel.addRow("Weapon Skill:", g.SelectedModel.WeaponSkill)
		panel.addRow("Ballistic Skill:", g.SelectedModel.BallisticSkill)
		panel.addRow("Strength:", strconv.Itoa(g.SelectedModel.Strength))
		panel.addRow("Toughness:", strconv.Itoa(g.SelectedModel.Toughness))
		panel.addRow("Wounds:", strconv.Itoa(g.SelectedModel.Wounds))
		panel.addRow("Attacks:", strconv.Itoa(g.SelectedModel.Attacks))
		panel.addRow("Leadership:", strconv.Itoa(g.SelectedModel.Leadership))
		panel.addRow("Save:", g.SelectedModel.Save)

		drawModelImage(g, panel)
	}

	return panel.Img
}

func drawModelImage(g *Game, panel *Panel) {
	ModelImg := getProfilePic(g)

	if(ModelImg == nil) {		
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(275, 84)

	panel.Img.DrawImage(ModelImg, op)
}


var ModelPics = make(map[string]*ebiten.Image)

func getProfilePic(g *Game) *ebiten.Image {

	imgPath := getImgPath(g.SelectedModel.Name)

	if(ModelPics[imgPath] != nil) {
		return ModelPics[imgPath]
	}

	path := fmt.Sprintf("./assets/armies/%s/images/%s.png",getImgPath(g.SelectedUnit.Army), imgPath)

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
