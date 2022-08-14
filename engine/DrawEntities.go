package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/ui"
)

func DrawEntities(background interfaces.Draw) {

	for c := 0; c < len(models.Game().BattleGround.Grid); c++ {

		for r := 0; r < len(models.Game().BattleGround.Grid[r]); r++ {

			if models.Game().BattleGround.Grid[c][r] == nil {
				continue
			}

			entity := models.Game().BattleGround.Grid[c][r]

			token := entity.GetToken()

			entityX := entity.GetLocation().X
			entitY := entity.GetLocation().Y

			if models.Game().PhaseStepper.GetPhase() == interfaces.MovementPhase && models.Game().DraggingUnit != nil {
				entityX = entityX - 10
				entitY = entitY - 10
			}

			// no need to render if outside viewport

			if entityX < models.Game().BattleGround.ViewPort.X && entitY > models.Game().BattleGround.ViewPort.Y {
				continue
			}

			x := float64((entity.GetLocation().X * ui.TileSize) + 1)
			y := float64((entity.GetLocation().Y * ui.TileSize) + 1)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(x, y)
			background.DrawImage(token, op)
		}
	}
	// draw move range
	// dont draw movement on acutal unit
	drawMoveRange(background)
}

func drawMoveRange(background interfaces.Draw) {
	if models.Game().PhaseStepper.GetPhase() != interfaces.MovementPhase || models.Game().SelectedPhaseUnit.Name == "" {
		return
	}

	unit := models.Game().SelectedPhaseUnit
	rect := models.Game().SelectedPhaseUnit.MovementRect

	drawMatrix := getMoveRange(unit, rect)

	for r := 0; r < len(drawMatrix); r++ {
		for c := 0; c < len(drawMatrix[c]); c++ {

			if drawMatrix[r][c] == 0 {
				continue
			}

			tile := ebiten.NewImage(31, 31)
			color := ui.GetTokenColor()
			tile.Fill(color)

			x := unit.Location.X + c
			y := unit.Location.Y + r

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((x*ui.TileSize)+1), float64((y*ui.TileSize)+1))
			background.DrawImage(tile, op)
		}
	}
}

func getMoveRange(unit *models.Unit, rect ui.Rect) [][]int {

	matrix := make([][]int, unit.Rect.W, unit.Rect.H)

	movement := unit.Models[0].Movement

	for c := unit.Location.X - movement; c < unit.Location.X+unit.Width+movement; c++ {
		for r := unit.Location.Y - movement; r < unit.Location.Y+unit.Height+movement; r++ {
			filled := models.Game().BattleGround.Grid[c][r] != nil

			if !filled {
				matrix[c][r] = 1
			}
		}
	}

	return matrix
}
