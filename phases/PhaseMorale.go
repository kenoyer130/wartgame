package phases

import (
	"fmt"

	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func StartPhaseMorale() {

	models.Game().StatusMessage.Phase = "Morale Phase"
	models.Game().StatusMessage.Messsage = "Select next unit to shoot!"
	models.Game().StatusMessage.Keys = "Press [Q] and [E] to cycle units! Press [Space] to select!"

	needMoraleCheck := false

	for i := 0; i < len(models.Game().Players); i++ {
		if(!models.Game().Players[i].MoraleChecked) {
			needMoraleCheck = true
			break
		}
	}

	if(!needMoraleCheck) {
		engine.WriteMessage("All Unit Morale Tests have completed!")
		MoveToPhase(models.EndPhase)
	}

	// loop through current target first
	for i := 0; i < len(models.Game().Players); i++ {

		if(models.Game().Players[i].MoraleChecked) {
			continue
		}

		models.Game().Players[i].MoraleChecked = true
		unitCycler := NewUnitCycler(&models.Game().Players[i], UnitTookCasulaties, MoraleCheckSelected)
		unitCycler.CycleUnits()
	}
}

func UnitTookCasulaties(unit *models.Unit) bool {
	return len(unit.DestroyedModels) > 0
}

func MoraleCheckSelected(unit *models.Unit) {

	if unit == nil {
		MoveToPhase(models.MoralePhase)	
		return
	}

	leadership := unit.GetMoraleCheck()

	engine.WriteMessage(fmt.Sprintf("Morale check for Unit %s with Leadership %d against %d", unit.Name, leadership, len(unit.DestroyedModels)))

	engine.RollDice("Rolling for Morale Test", engine.DiceRollType{
		Dice:      1,
		Target:    leadership,
		AddToDice: len(unit.DestroyedModels),
	}, func(success int, dice []int) {
		if success < 0 {
			MoveToPhase(models.MoralePhase)
		} else {
			failMorale(unit, leadership)
			MoveToPhase(models.MoralePhase)
		}
	})
}

func failMorale(unit *models.Unit, leadership int) {
	// always lose one
	unit.MoraleFailure()

	target := 1

	if ((len(unit.Models)) / 2) < (unit.OriginalModelCount) {
		target = 2
	}

	for _, model := range unit.Models {
		engine.RollDice(fmt.Sprintf("%s rolling for Combat Attrition Test", model.Name), engine.DiceRollType{
			Dice:   1,
			Target: target,
		}, func(failure int, dice []int) {
			if failure > 0 {
				unit.MoraleFailure()
			}
		})
	}

	StartPhaseMorale()
}

func PhaseMoraleCleanup () {	
	// clear out morale check for all players
	for i := 0; i < len(models.Game().Players); i++ {
		models.Game().Players[i].MoraleChecked = false
	}
}
