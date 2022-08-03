package engine

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/ui"
)

func (g *Game) Update() error {

	updateState(g, g.CurrentGameState)
	return nil
}

func updateState(g *Game, s models.GameState) {

	checkInputs(g)

	switch s {
	case models.GameStart:

	}
}

func checkInputs(g *Game) {
	checkEsc()
	checkGridInputs(g)
	checkModelNav(g)
	checkUnitSelection(g)
}

func checkUnitSelection(g *Game) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		
		for _, player := range g.Players {
			for _, unit := range player.Army.Units {
				cx, cy := ebiten.CursorPosition()
				if unit.Rect.InPixelBounds(cx, cy) {
					g.SelectedUnit = &unit
					g.SelectedModel = &unit.Models[0]
				}
			}
		}
	}
}

func checkModelNav(g *Game) {

	if inpututil.IsKeyJustReleased(ebiten.KeyB) {
		UnitSelector(1, g)
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyV) {
		UnitSelector(-1, g)
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyN) {
		ModelSelector(1, g)
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyM) {
		ModelSelector(-1, g)
	}
}

func UnitSelector(direction int, g *Game) {
	UnitSelected := ensureSelectedUnit(g)
	if UnitSelected {
		return
	}

	currentUnit := 0

	// loop through Unit array and find our selected Unit
	for i := 0; i < len(g.CurrentPlayer.Army.Units); i++ {
		if &g.CurrentPlayer.Army.Units[i] == g.SelectedUnit {
			currentUnit = i
		}
	}

	currentUnit = currentUnit + direction

	if currentUnit < 0 {
		currentUnit = len(g.CurrentPlayer.Army.Units) - 1
	}

	if currentUnit > len(g.CurrentPlayer.Army.Units)-1 {
		currentUnit = 0
	}

	g.SelectedUnit = &g.CurrentPlayer.Army.Units[currentUnit]
	g.SelectedModel = &g.CurrentPlayer.Army.Units[currentUnit].Models[0]
}

func ensureSelectedUnit(g *Game) bool {
	if g.SelectedUnit == nil {
		g.SelectedUnit = &g.CurrentPlayer.Army.Units[0]
		g.SelectedModel = &g.CurrentPlayer.Army.Units[0].Models[0]
		return true
	}
	return false
}

func ModelSelector(direction int, g *Game) {
	UnitSelected := ensureSelectedUnit(g)
	if UnitSelected {
		return
	}

	currentModel := 0

	currentModel = currentModel + direction

	if currentModel < 0 {
		currentModel = len(g.SelectedUnit.Models) - 1
	}

	if currentModel > len(g.SelectedUnit.Models)-1 {
		currentModel = 0
	}

	g.SelectedModel = &g.SelectedUnit.Models[currentModel]
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
				InDrag:       true,
				CursorStartX: cursorX,
				CursorStartY: cursorY}
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.UIState.GridDragging = DraggingGrid{}
	}
}

func cursorInViewport(g *Game) bool {
	cursorX, cursorY := ebiten.CursorPosition()

	topViewPortX := 0
	topViewPortY := 0

	bottomViewPortX := (g.BattleGround.ViewPort.X) + (g.BattleGround.ViewPort.Width)*ui.TileSize
	bottomViewPortY := (g.BattleGround.ViewPort.Y) + (g.BattleGround.ViewPort.Height)*ui.TileSize

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
