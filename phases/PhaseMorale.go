package phases

import (
	"fmt"

	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func StartPhaseMorale() {

	models.Game().StatusMessage.Phase = "Morale Phase"
	models.Game().StatusMessage.Messsage = "Select unit to perform moral check!"
	models.Game().StatusMessage.Keys = "Press [Q] and [E] to cycle units! Press [Space] to select!"

	checkMoraleForPlayer(0, func() {
		checkMoraleForPlayer(1, func() {
			MoveToPhase(models.EndPhase)
		})
	})
}

func checkMoraleForPlayer(player int, onCompleted func()) {
	engine.WriteMessage(fmt.Sprintf("Checking Morale for %s", models.Game().Players[player].Name))

	models.Game().Players[player].MoraleChecked = true

	unitCycler := NewUnitCycler(&models.Game().Players[player], UnitTookCasulaties, func(unit *models.Unit) {
		if unit == nil {
			engine.WriteMessage("No units need morale check!")
			onCompleted()
			return
		}

		MoraleCheckSelected(unit, onCompleted)
	})

	unitCycler.CycleUnits()
}

func UnitTookCasulaties(unit *models.Unit) bool {
	return len(unit.DestroyedModels) > 0
}

func MoraleCheckSelected(unit *models.Unit, onCompleted func()) {

	leadership := unit.GetMoraleCheck()

	engine.WriteMessage(fmt.Sprintf("Morale check for Unit %s with Leadership %d against %d", unit.Name, leadership, len(unit.DestroyedModels)))

	engine.RollDice("Rolling for Morale Test", engine.DiceRollType{
		Dice:      1,
		Target:    leadership,
		AddToDice: len(unit.DestroyedModels),
	}, func(success int, dice []int) {
		if success < 1 {
			onCompleted()
		} else {
			failMorale(unit, leadership, onCompleted)
		}
	})
}

func failMorale(unit *models.Unit, leadership int, onCompleted func()) {
	// always lose one
	engine.WriteMessage("1 lost model due to failed moral check")
	unit.MoraleFailure()

	target := 1

	count := len(unit.Models)

	if count > (unit.OriginalModelCount / 2) {
		target = 2
	}

	engine.RollDice(fmt.Sprintf("%s rolling for Combat Attrition Test", unit.Name), engine.DiceRollType{
		Dice:   count,
		Target: target,
	}, func(success int, dice []int) {

		fail := (count - success)
		engine.WriteMessage(fmt.Sprintf("%d lost model due to failed moral check", fail))

		for i := 0; i < fail; i++ {
			unit.MoraleFailure()
			onCompleted()
		}
	})
}
