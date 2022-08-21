package phases

import (
	"fmt"

	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
)

type MoralePhase struct {
}

func (re MoralePhase) GetName() (interfaces.GamePhase, interfaces.PhaseStep) {
	return interfaces.MoralePhase, interfaces.Nil
}

func (re MoralePhase) Start() {

	models.Game().StatusMessage.Phase = "Morale Phase"
	models.Game().StatusMessage.Messsage = "Select unit to perform moral check!"
	models.Game().StatusMessage.Keys = "Press [Q] and [E] to cycle units! Press [Space] to select!"

	re.checkMoraleForPlayer(0, func() {
		re.checkMoraleForPlayer(1, func() {
			models.Game().PhaseStepper.Move(interfaces.EndPhase)
		})
	})
}

func (re MoralePhase) checkMoraleForPlayer(player int, onCompleted func()) {
	engine.WriteMessage(fmt.Sprintf("Checking Morale for %s", models.Game().Players[player].Name))

	models.Game().Players[player].MoraleChecked = true

	unitCycler := NewUnitCycler(&models.Game().Players[player], re.UnitTookCasulaties, func(unit *models.Unit) {
		if unit == nil {
			engine.WriteMessage("No units need morale check!")
			onCompleted()
			return
		}

		re.MoraleCheckSelected(unit, onCompleted)
	}, false)

	unitCycler.CycleUnits()
}

func (re MoralePhase) UnitTookCasulaties(unit *models.Unit) bool {
	return len(unit.DestroyedModels) > 0
}

func (re MoralePhase) MoraleCheckSelected(unit *models.Unit, onCompleted func()) {

	leadership := unit.GetMoraleCheck()

	engine.WriteMessage(fmt.Sprintf("Morale check for Unit %s with Leadership %d adding %d", unit.Token.ID, leadership, len(unit.DestroyedModels)))

	models.Game().DiceRoller.Roll("Rolling for Morale Test", interfaces.DiceRollType{
		Dice:      1,
		Target:    leadership,
		AddToDice: len(unit.DestroyedModels),
	},
	nil,
	 func(success int, dice []int) {
		if success < 1 {
			onCompleted()
		} else {
			re.failMorale(unit, leadership, onCompleted)
		}
	})
}

func (re MoralePhase) failMorale(unit *models.Unit, leadership int, onCompleted func()) {
	// always lose one
	engine.WriteMessage("1 lost model due to failed moral check")
	unit.MoraleFailure()

	target := 1

	count := len(unit.Models)

	if count > (unit.OriginalModelCount / 2) {
		target = 2
	}

	models.Game().DiceRoller.Roll(fmt.Sprintf("%s rolling for Combat Attrition Test", unit.Name), interfaces.DiceRollType{
		Dice:   count,
		Target: target,
	},
	nil,
	func(success int, dice []int) {

		fail := (count - success)
		engine.WriteMessage(fmt.Sprintf("%d lost model due to failed moral check", fail))

		for i := 0; i < fail; i++ {
			unit.MoraleFailure()
			onCompleted()
		}
	})
}
