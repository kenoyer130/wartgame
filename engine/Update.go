package engine

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kenoyer130/wartgame/models"
)

func (g *Game) Update() error {

	updateState(g, g.CurrentGameState)
	return nil
}

func updateState(g *Game, s models.GameState) {

	checkInputs(g)

	switch s {
	case models.Start:

	}
}

func checkInputs(g *Game) {
	checkEsc()
	checkGridInputs(g)
	checkUnitNav(g)
}

func checkUnitNav(g *Game) {
	
	if inpututil.IsKeyJustReleased(ebiten.KeyB) {
		squadSelector(1, g)
	}
	
	if inpututil.IsKeyJustReleased(ebiten.KeyV) {
		squadSelector(-1, g)
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyN) {
		unitSelector(1, g)
	}
	
	if inpututil.IsKeyJustReleased(ebiten.KeyM) {
		unitSelector(-1, g)
	}
}

func squadSelector(direction int, g *Game) {
	squadSelected := ensureSelectedSquad(g)
	if squadSelected {
		return
	}

	currentSquad := 0

	// loop through squad array and find our selected squad
	for i := 0; i < len(g.CurrentPlayer.Army.Squads); i++ {
		if(&g.CurrentPlayer.Army.Squads[i] == g.SelectedSquad) {
			currentSquad = i
		}
	}

	currentSquad = currentSquad + direction

	if currentSquad< 0 {
		currentSquad = len(g.CurrentPlayer.Army.Squads) - 1
	}

	if currentSquad > len(g.CurrentPlayer.Army.Squads) - 1 {
		currentSquad = 0
	}

	g.SelectedSquad = &g.CurrentPlayer.Army.Squads[currentSquad]
	g.SelectedUnit = &g.CurrentPlayer.Army.Squads[currentSquad].Units[0]
}

func ensureSelectedSquad(g *Game) bool {
	if g.SelectedSquad == nil {
		g.SelectedSquad = &g.CurrentPlayer.Army.Squads[0]
		g.SelectedUnit = &g.CurrentPlayer.Army.Squads[0].Units[0]
		return true
	}
	return false
}

func unitSelector(direction int, g *Game) {
	squadSelected := ensureSelectedSquad(g)
	if squadSelected {
		return
	}

	currentUnit := 0

	currentUnit = currentUnit + direction

	if currentUnit< 0 {
		currentUnit = len(g.SelectedSquad.Units) - 1
	}

	if currentUnit > len(g.SelectedSquad.Units) - 1 {
		currentUnit = 0
	}

	g.SelectedUnit = &g.SelectedSquad.Units[currentUnit]
}

func checkEsc() {
	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		os.Exit(0)
	}
}

func checkGridInputs(g *Game) {

	checkGridToggle(g)

	checkGridDrag(g)
}

func checkGridDrag(g *Game) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		//log.Printf("cursor in viewport: %d cx %d cy %d topViewPortX %d topViewPortY %d bottomViewPortX %d bottomViewPortY",
		//	cursorX, cursorY, topViewPortX, topViewPortY, bottomViewPortX, bottomViewPortY)
		if cursorInViewport(g) && !g.UIState.GridDragging.InDrag {
			cursorX, cursorY := ebiten.CursorPosition()

			g.UIState.GridDragging = DraggingGrid{
				InDrag: true,
				CursorStartX: cursorX,
				CursorStartY: cursorY }
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.UIState.GridDragging = DraggingGrid {}
	}
}

func cursorInViewport(g *Game) bool {
	cursorX, cursorY := ebiten.CursorPosition()

	topViewPortX := 0
	topViewPortY := 0

	bottomViewPortX := (g.BattleGround.ViewPort.X) + (g.BattleGround.ViewPort.Width)*models.TileSize
	bottomViewPortY := (g.BattleGround.ViewPort.Y) + (g.BattleGround.ViewPort.Height)*models.TileSize

	cursorInViewport := (cursorX > topViewPortX && cursorY > topViewPortY) && (cursorX < bottomViewPortX && cursorY < bottomViewPortY)

	return cursorInViewport
}

func checkGridToggle(g *Game) {
	if inpututil.IsKeyJustReleased(ebiten.KeyG) {
		if g.UIState.ShowGrid {
			g.UIState.ShowGrid = false
		} else {
			g.UIState.ShowGrid = true
		}
	}
}
