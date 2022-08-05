package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kenoyer130/wartgame/ui"
)

type Panel struct {
	Img *ebiten.Image
}

func NewPanel(x, y int) *Panel {
	var p Panel
	p.Img = ebiten.NewImage(x, y)
	return &p
}

func (re Panel) addTitle(title string) {
	text.Draw(re.Img, title, ui.GetFontBold(), ui.Margin, 25, ui.GetTextColor())
}

func (re Panel) addMessage(msg string, r int) {
	text.Draw(re.Img, msg, ui.GetFontNormalFace(), ui.Margin+10, r*25, ui.GetTextColor())
}

func (re Panel) addRow(label, value string, r int) {
	re.addLabel(label, r)
	re.addValue(value, r, 0)
}

func (re Panel) addLabel(label string, r int) {
	text.Draw(re.Img, label, ui.GetFontBold(), ui.Margin, r*25, ui.GetTextColor())
}

func (re Panel) addValue(value string, r int, c int) {
	text.Draw(re.Img, value, ui.GetFontNormalFace(), (ui.Margin+150)+(c*35), r*25, ui.GetTextColor())
}
