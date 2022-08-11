package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/ui"
)

func DrawEntities(background *ebiten.Image) {
	entites := models.Game().BattleGround.Grid

	for _, entity := range entites {
		token := entity.GetToken()

		entityX := entity.GetLocation().X
		entitY := entity.GetLocation().Y

		if(models.Game().PhaseStepper.GetPhase() == models.MovementPhase && models.Game().DraggingUnit != nil) {
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
}
