package phases

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

type UnitCycler struct {
	player         *models.Player	
	validUnit      func(unit *models.Unit) bool
	onUnitSelected func(unit *models.Unit)
}

func NewUnitCycler(player *models.Player,	
	validUnit func(unit *models.Unit) bool,
	onUnitSelected func(unit *models.Unit)) *UnitCycler {
	return &UnitCycler{
		player:         player,		
		validUnit:      validUnit,
		onUnitSelected: onUnitSelected,
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

func (re UnitCycler) CycleUnits() {
	// cycle units and select first valid unit
	unit := re.selectNextUnit(0, -1)

	// if no valid unit return nil
	if unit == nil {
		re.onUnitSelected(nil)
		return
	}

	models.Game().SelectedUnit = unit

	// register Q and E to cycle units
	engine.KeyBoardRegistry[ebiten.KeyQ] = func() {

		index := indexOfUnit(models.Game().SelectedUnit, re.player.Army.Units)
		index--

		// if no valid unit return nil
		cycleUnits(re, index)		
	}

	engine.KeyBoardRegistry[ebiten.KeyE] = func() {
		index := indexOfUnit(models.Game().SelectedUnit, re.player.Army.Units)
		index++

		// if no valid unit return nil
		cycleUnits(re, index)		
	}

	engine.KeyBoardRegistry[ebiten.KeySpace] = func() {
		re.onUnitSelected(models.Game().SelectedUnit)
	}
}

func cycleUnits(re UnitCycler, index int) bool {
	unit := re.selectNextUnit(index, index)

	if unit == nil {
		re.onUnitSelected(nil)
		return true
	}

	models.Game().SelectedUnit = unit
	return false
}

func (re UnitCycler) selectNextUnit(index int, start int) *models.Unit {

	// fix index for cycling
	index = re.wrapIndex(index)

	if(index == start) {
		return nil
	}

	if(start == -1) {
		start = 0
	}

	if !re.validUnit(re.player.Army.Units[index]) {
		index++
		return re.selectNextUnit(index, start)
	}

	return re.player.Army.Units[index]
}

func (re UnitCycler) wrapIndex(index int) int {
	if index < 0 {
		index = len(re.player.Army.Units) - 1
	}

	if index > len(re.player.Army.Units)-1 {
		index = 0
	}
	return index
}
