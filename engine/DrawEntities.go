package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/ui"
)

func DrawEntities(g *Game, background *ebiten.Image) {
	entites := g.BattleGround.Grid

	for _, entity := range entites {
		token := entity.GetToken()

		entityX := entity.GetLocation().X
		entitY := entity.GetLocation().Y

		// no need to render if outside viewport

		if entityX < g.BattleGround.ViewPort.X && entitY > g.BattleGround.ViewPort.Y {
			continue
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((entity.GetLocation().X*ui.TileSize)+1), float64((entity.GetLocation().Y*ui.TileSize)+1))
		background.DrawImage(token, op)
	}
}
