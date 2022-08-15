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

			x := float64((entityX * ui.TileSize) + 1)
			y := float64((entitY * ui.TileSize) + 1)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(x, y)
			background.DrawImage(token, op)
		}
	}

	DrawMoveRange(background)
}

func DrawMoveRange(background interfaces.Draw) {
	if models.Game().PhaseStepper.GetPhase() != interfaces.MovementPhase || models.Game().SelectedPhaseUnit.Name == "" {
		return
	}

	mUnit := models.Game().SelectedPhaseUnit

	for _, loc := range mUnit.MovementRange {

		tile := ebiten.NewImage(31, 31)
		color := ui.GetMoveRangeColor()
		tile.Fill(color)

		x := loc.X
		y := loc.Y

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((x*ui.TileSize)+1), float64((y*ui.TileSize)+1))
		background.DrawImage(tile, op)
	}
}
