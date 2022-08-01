package engine

import (
	"fmt"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

func getGameInfoPanel(g *Game) *ebiten.Image {

	panel := NewPanel(400, 800)

	panel.addTitle("Game Info")
	panel.addRow("Round: ", strconv.Itoa(g.Round))
	panel.addRow("Current Player: ", g.CurrentPlayer.Name)
	panel.addRow("Current Phase: ", fmt.Sprintf("%s", g.CurrentPhase))

	return panel.Img
}
