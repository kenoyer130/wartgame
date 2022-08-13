package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	interfaces "github.com/kenoyer130/wartgame/engine/Interfaces"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/ui"
)

func DrawEntities(background interfaces.Draw) {
	entites := models.Game().BattleGround.Grid

	for _, entity := range entites {
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

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((entity.GetLocation().X*ui.TileSize)+1), float64((entity.GetLocation().Y*ui.TileSize)+1))
		background.DrawImage(token, op)
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

	matrix := [unit.Rect.W][unit.Rect.H]int{}

	for r := 0; r < rect.H; r++ {
		for c := 0; c < rect.W; c++ {

			if unit.Rect.InBounds(rect.X+r, rect.Y+c) {
				continue
			}

			tile := ebiten.NewImage(31, 31)
			color := ui.GetTokenColor()
			tile.Fill(color)

			x := unit.Location.X + c
			y := unit.Location.Y + r
		}
	}
}
