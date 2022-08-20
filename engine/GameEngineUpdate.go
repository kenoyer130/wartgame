package engine

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/ui"
)

func UpdateGameEngine() error {

	updateState()
	return nil
}

func updateState() {

	s := models.Game().CurrentGameState

	checkInputs()

	switch s {
	case models.GameStart:

	}
}

func checkInputs() {
	checkEsc()
	checkGridInputs()
	checkUnitSelection()
	checkKeyboardRegistery()
}

func checkKeyboardRegistery() {
	for key := range KeyBoardRegistry {
		if inpututil.IsKeyJustPressed(key) {
			KeyBoardRegistry[key]()
		}
	}
}

func checkUnitSelection() {
	
}

func checkEsc() {
	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		//pprof.StopCPUProfile()
		os.Exit(0)
	}
}

func checkGridInputs() {

	checkGridToggle()

	checkGridDrag()
}

func checkGridDrag() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		//log.Printf("cursor in viewport: %d cx %d cy %d topViewPortX %d topViewPortY %d bottomViewPortX %d bottomViewPortY",
		//	cursorX, cursorY, topViewPortX, topViewPortY, bottomViewPortX, bottomViewPortY)
		if cursorInViewport() && !models.Game().UIState.GridDragging.InDrag {
			cursorX, cursorY := ebiten.CursorPosition()

			models.Game().UIState.GridDragging = models.DraggingGrid{
				InDrag:       true,
				CursorStartX: cursorX,
				CursorStartY: cursorY}
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		models.Game().UIState.GridDragging = models.DraggingGrid{}
	}
}

func cursorInViewport() bool {
	cursorX, cursorY := ebiten.CursorPosition()

	topViewPortX := 0
	topViewPortY := 0

	bottomViewPortX := (models.Game().BattleGround.ViewPort.X) + (models.Game().BattleGround.ViewPort.Width)*ui.TileSize
	bottomViewPortY := (models.Game().BattleGround.ViewPort.Y) + (models.Game().BattleGround.ViewPort.Height)*ui.TileSize

	cursorInViewport := (cursorX > topViewPortX && cursorY > topViewPortY) && (cursorX < bottomViewPortX && cursorY < bottomViewPortY)

	return cursorInViewport
}

func checkGridToggle() {
	if inpututil.IsKeyJustReleased(ebiten.KeyG) {
		if models.Game().UIState.ShowGrid {
			models.Game().UIState.ShowGrid = false
		} else {
			models.Game().UIState.ShowGrid = true
		}
	}
}
