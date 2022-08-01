package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/ui"
)

func (g *Game) Draw(screen *ebiten.Image) {

	background := ebiten.NewImage(g.BattleGround.ViewPort.Width*models.TileSize, g.BattleGround.ViewPort.Height*models.TileSize)
	background.Fill(color.RGBA{166, 142, 154, 1})

	if g.UIState.ShowGrid {
		drawGrid(background, g)
	}

	drawSelectedModelInfo(g, background)

	drawEntities(g, background)

	screen.DrawImage(background, nil)

	drawGameInfoPanel(g, screen)
	drawModelPanel(g, screen)
}

func drawGameInfoPanel(g *Game, screen *ebiten.Image) {
	gameInfoPanel := getGameInfoPanel(g)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.BattleGround.ViewPort.GetPixelRectangle().Width+ui.Margin), 0)
	screen.DrawImage(gameInfoPanel, op)
}

func drawModelPanel(g *Game, screen *ebiten.Image) {
	ModelPanel := getModelPanel(g)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.BattleGround.ViewPort.GetPixelRectangle().Width+ui.Margin), 210)
	screen.DrawImage(ModelPanel, op)
}

func drawSelectedModelInfo(g *Game, background *ebiten.Image) {
	if g.SelectedModel == nil {
		return
	}
}

func drawEntities(g *Game, background *ebiten.Image) {
	entites := g.BattleGround.Grid

	for _, entity := range entites {
		token := entity.GetToken()

		entityX := entity.GetLocation().X
		entitY := entity.GetLocation().Y

		// no need to render if outside viewport

		if (entityX < g.BattleGround.ViewPort.X && entitY > g.BattleGround.ViewPort.Y) || (entityX < g.BattleGround.ViewPort.X && entitY > g.BattleGround.ViewPort.Y) {
			continue
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((entity.GetLocation().X*32)+1), float64((entity.GetLocation().Y*32)+1))
		background.DrawImage(token, op)
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
