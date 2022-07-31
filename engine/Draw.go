package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {

	image := ebiten.NewImage(g.BattleGround.ViewPort.Width*32, g.BattleGround.ViewPort.Height*32)
	image.Fill(color.RGBA{166, 142, 154, 1})

	screen.DrawImage(image, nil)

	if g.ShowGrid {
		drawGrid(screen, g)
	}

	entites := g.BattleGround.Grid

	for _, entity := range entites {
		token := entity.GetToken()

		entityX := entity.GetLocation().X
		entitY := entity.GetLocation().Y

		// no need to render if outside viewport
		if ((entityX < g.BattleGround.ViewPort.X && entitY > g.BattleGround.ViewPort.Y) ||  (entityX < g.BattleGround.ViewPort.X && entitY > g.BattleGround.ViewPort.Y)) {
			continue
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((entity.GetLocation().X*32)+1), float64((entity.GetLocation().Y*32)+1))
		screen.DrawImage(token, op)
	}
}

func drawGrid(screen *ebiten.Image, g *Game) {
	for x := 0; x < g.BattleGround.ViewPort.Width; x++ {

		image := ebiten.NewImage((g.BattleGround.Size.X * 32), 1)
		image.Fill(color.RGBA{35, 31, 33, 1})

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, float64(x*32))

		screen.DrawImage(image, op)
	}

	for y := 0; y < g.BattleGround.ViewPort.Height; y++ {

		image := ebiten.NewImage(1, (g.BattleGround.Size.Y * 32))
		image.Fill(color.RGBA{35, 31, 33, 1})

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(y*32), 0)

		screen.DrawImage(image, op)
	}
}
