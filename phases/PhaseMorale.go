package phases

import (
	"fmt"

	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
)

type MoralePhase struct {
	moraleLoss int
}

func (re MoralePhase) GetName() (interfaces.GamePhase, interfaces.PhaseStep) {
	return interfaces.MoralePhase, interfaces.Nil
}

func (re MoralePhase) Start() {

	re.moraleLoss = 0

	models.Game().StatusMessage.Phase = "Morale Phase"
	
	models.Game().Players[0].Army.RemoveDestroyedUnits()
	models.Game().Players[1].Army.RemoveDestroyedUnits()
	
	re.checkMoraleForPlayer(0, func() {
		re.checkMoraleForPlayer(1, func() {
			models.Game().PhaseEventBus.Fire("MoralePhaseEnded")
		})
	})

	engine.WriteMessage(fmt.Sprintf("%d lost due to failed moral check!", re.moraleLoss))
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

	if 6+len(unit.DestroyedModels) <= leadership {
		engine.WriteMessage("Automatic pass!")
		onCompleted()
		return
	}

	success, _ := models.Game().DiceRoller.Roll("Rolling for Morale Test", interfaces.DiceRollType{
		Dice:      1,
		Target:    leadership,
		AddToDice: len(unit.DestroyedModels),
	},
		nil)

	if success > 0 {
		re.failMorale(unit, leadership, onCompleted)
	}

	onCompleted()
}

// always lose one
func (re *MoralePhase) failMorale(unit *models.Unit, leadership int, onCompleted func()) {
	engine.WriteMessage("1 lost model due to failed moral check")
	re.moraleLoss++
	unit.MoraleFailure()

	target := 1

	count := len(unit.Models)

	if count > (unit.OriginalModelCount / 2) {
		target = 2
	}

	_, dice := models.Game().DiceRoller.Roll(fmt.Sprintf("%s rolling for Combat Attrition Test", unit.Name), interfaces.DiceRollType{
		Dice:   count,
		Target: target,
	},
		nil)

	fail := 0

	for _, die := range dice {
		if die == 1 {
			fail++
			re.moraleLoss++
		}
	}

	engine.WriteMessage(fmt.Sprintf("%d lost model due to failed moral check", fail))

	if fail == 0 {
		return
	}

	for fail > 0 {
		unit.MoraleFailure()
		fail--
	}
}
