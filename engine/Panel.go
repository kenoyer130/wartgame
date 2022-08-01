package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kenoyer130/wartgame/ui"
)

type Panel struct {
	row int
	Img *ebiten.Image
}

func NewPanel(x, y int) *Panel {
	var p Panel

	p.Img = ebiten.NewImage(x,y)
	p.row = 2

	return &p
}

func (re Panel) addTitle(title string) {
	text.Draw(re.Img, title, ui.GetFontBold(), ui.Margin, 25, ui.GetTextColor())
}

func (re *Panel) addRow(label string, value string) {
	text.Draw(re.Img, label, ui.GetFontBold(), ui.Margin, re.row*25, ui.GetTextColor())
	text.Draw(re.Img, value, ui.GetFontNormalFace(), ui.Margin+150, re.row*25, ui.GetTextColor())

	re.row++
}
