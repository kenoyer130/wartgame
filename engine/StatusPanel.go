package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kenoyer130/wartgame/ui"
)

func getStatusPanel(msg string) *ebiten.Image {

	panel := ebiten.NewImage(800, 35)

	text.Draw(panel, msg, ui.GetFontNormalFace(), ui.Margin+10,25, ui.GetTextColor())

	return panel
}
