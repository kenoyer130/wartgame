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

	if models.Game().DraggingUnit != nil {

		marker := ebiten.NewImage(31, 31)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(models.Game().DraggingUnitLocation.X), float64(models.Game().DraggingUnitLocation.Y))
		color := ui.GetGridOutlineColor()
		marker.Fill(color)

		//text.Draw(token, re.Token.ID, ui.GetFontNormalFace(), 2, 24, ui.GetTextColor())
		background.DrawImage(marker, op)

	}
}