package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/ui"
)

func getStatusPanel() *ebiten.Image {

	panel := ebiten.NewImage(800, 35)

	text.Draw(panel, models.Game().StatusMessage.Phase, ui.GetFontNormalFace(), ui.Margin, 25, ui.GetTextColor())
	text.Draw(panel, models.Game().StatusMessage.Messsage, ui.GetFontNormalFace(), ui.Margin+150, 25, ui.GetTextColor())
	text.Draw(panel, models.Game().StatusMessage.Keys, ui.GetFontNormalFace(), ui.Margin+550, 25, ui.GetTextColor())

	return panel
}
