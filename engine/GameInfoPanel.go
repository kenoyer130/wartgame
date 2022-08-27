package engine

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/models"
)

func getGameInfoPanel() *ebiten.Image {

	panel := NewPanel(400, 800)

	panel.addTitle("Game Info")
	panel.addRow("Round: ", strconv.Itoa(models.Game().Round), 2)
	panel.addRow("Current Player: ", models.Game().CurrentPlayer.Name, 3)
	panel.addRow("Current Phase: ", string(models.Game().PhaseStepper.GetPhaseName()), 4)

	return panel.Img
}
