package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kenoyer130/wartgame/ui"
)

type PanelVertical struct {
	Img *ebiten.Image
}

func NewPanelVertical(x int, y int) *PanelVertical {
	var p PanelVertical
	p.Img = ebiten.NewImage(x, y)
	return &p
}

func (re PanelVertical) addTitle(title string) {
	text.Draw(re.Img, title, ui.GetFontBold(), ui.Margin, 25, ui.GetTextColor())
}

func (re *PanelVertical) addLabel(label string, r int, c int, w int) {
	text.Draw(re.Img, label, ui.GetFontBold(), ui.Margin+w,  r*25, ui.GetTextColor())
}

func (re *PanelVertical) addValue(value string, r int, c int, w int) {
	text.Draw(re.Img, value, ui.GetFontNormalFace(), ui.Margin+w, r*25, ui.GetTextColor())
}
