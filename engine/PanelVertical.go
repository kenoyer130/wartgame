package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kenoyer130/wartgame/ui"
)

type PanelVertical struct {
	Img       *ebiten.Image
	colWidths map[int]int
}

func NewPanelVertical(x int, y int, colwidths map[int]int) *PanelVertical {
	var p PanelVertical
	p.Img = ebiten.NewImage(x, y)
	p.colWidths = colwidths
	return &p
}

func (re PanelVertical) addTitle(title string) {
	text.Draw(re.Img, title, ui.GetFontBold(), ui.Margin, 25, ui.GetTextColor())
}

func (re *PanelVertical) addLabel(label string, c int) {
	text.Draw(re.Img, label, ui.GetFontBold(), ui.Margin+(re.getColWidth(c)), 50, ui.GetTextColor())
}

func (re *PanelVertical) addValue(value string, r int, c int) {
	text.Draw(re.Img, value, ui.GetFontNormalFace(), ui.Margin+(re.getColWidth(c)), r*35, ui.GetTextColor())
}

func (re *PanelVertical) getColWidth(c int) int {

	return c * 50

	// currentC := 0
	// currentW := 0

	// for _, colWidth := range re.colWidths {
	// 	if currentC == c {
	// 		return currentW
	// 	}

	// 	currentW += colWidth

	// 	currentC++
	// }

	// return 0
}
