package phases

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

type UnitCycler struct {
	player            *models.Player
	validUnit         func(unit *models.Unit) bool
	onUnitSelected    func(unit *models.Unit)
	currentUnit       *models.Unit
	suppressSelection bool
	checkedUnits      map[string]bool
}

func NewUnitCycler(player *models.Player,
	validUnit func(unit *models.Unit) bool,
	onUnitSelected func(unit *models.Unit),
	suppressSelection bool) *UnitCycler {
	return &UnitCycler{
		player:            player,
		validUnit:         validUnit,
		onUnitSelected:    onUnitSelected,
		suppressSelection: suppressSelection,
	}
}

func indexOfUnit(element *models.Unit, data []*models.Unit) int {
	for k, v := range data {
		if element.Name == v.Name {
			return k
		}
	}
	return -1 //not found.
}

func (re *UnitCycler) CycleUnits() {

	re.checkedUnits = make(map[string]bool)

	// cycle units and select first valid unit
	re.selectNextUnit(0, -1)

	// if no valid unit return nil
	if re.currentUnit == nil {
		re.onUnitSelected(nil)
		return
	}

	models.Game().SelectedUnit = re.currentUnit

	// register Q and E to cycle units
	if !re.suppressSelection {
		engine.KeyBoardRegistry[ebiten.KeyQ] = func() {

			index := indexOfUnit(models.Game().SelectedUnit, re.player.Army.Units)
			index--

			// if no valid unit return nil
			re.cycleUnits(index)
		}

		engine.KeyBoardRegistry[ebiten.KeyE] = func() {
			index := indexOfUnit(models.Game().SelectedUnit, re.player.Army.Units)
			index++

			// if no valid unit return nil
			re.cycleUnits(index)
		}
	}

	if re.suppressSelection {
		re.onUnitSelected(re.currentUnit)
	} else {

		engine.KeyBoardRegistry[ebiten.KeySpace] = func() {
			re.onUnitSelected(re.currentUnit)
		}
	}
}

func (re *UnitCycler) cycleUnits(index int) bool {
	re.selectNextUnit(index, index)

	if re.currentUnit == nil {
		re.onUnitSelected(nil)
		return true
	}

	models.Game().SelectedUnit = re.currentUnit

	return false
}

func (re *UnitCycler) selectNextUnit(index int, start int) {

	if(len(re.player.Army.Units) == 0) {
		re.currentUnit = nil
		return
	}

	// fix index for cycling
	index = re.wrapIndex(index)

	key := re.player.Army.Units[index].Name

	if re.checkedUnits[key] {
		return
	}

	re.checkedUnits[key] = true

	if start == -1 {
		start = 0
	}

	if !re.validUnit(re.player.Army.Units[index]) {
		index++
		re.selectNextUnit(index, start)
		return
	}

	re.currentUnit = re.player.Army.Units[index]
}

func (re *UnitCycler) wrapIndex(index int) int {
	if index < 0 {
		index = len(re.player.Army.Units) - 1
	}

	if index > len(re.player.Army.Units)-1 {
		index = 0
	}
	return index
}
