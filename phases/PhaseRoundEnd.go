package phases

import (
	"fmt"

	"github.com/kenoyer130/wartgame/engine"
	interfaces "github.com/kenoyer130/wartgame/engine/Interfaces"
	"github.com/kenoyer130/wartgame/models"
)

type EndPhase struct {
}

func (re EndPhase) GetName() (interfaces.GamePhase, interfaces.PhaseStep) {
	return interfaces.EndPhase, interfaces.Nil
}

func (re EndPhase) Start() {

	lost := false

	for _, player := range models.Game().Players {
		lost := re.checkPlayerHasUnits(player)

		if lost {
			break
		}
	}

	if !lost {
		re.startNextTurn()
	}
}

func (re EndPhase) checkPlayerHasUnits(player models.Player) bool {
	if len(player.Army.Units) < 1 {
		engine.WriteMessage(fmt.Sprintf("%s player has lost the game due to no units left!", player.Name))
		return true
	}

	return false
}

func (re EndPhase) startNextTurn() {
	models.Game().CurrentPlayer.Gone = true

	allGone := true

	for i := 0; i < len(models.Game().Players); i++ {

		player := &models.Game().Players[i]

		if player.Name == models.Game().CurrentPlayer.Name {
			player.Gone = true
		} else {

			if !player.Gone {
				allGone = false
				models.Game().CurrentPlayer = player
				models.Game().CurrentPlayerIndex = i
			}
		}
	}

	re.EndPhaseCleanup()

	if allGone {
		re.startNewRound()
	}

	models.Game().PhaseStepper.Move(interfaces.ShootingPhase)
}

func (re EndPhase) startNewRound() {

	for i := 0; i < len(models.Game().Players); i++ {
		models.Game().Players[i].RoundCleanup()
	}

	models.Game().CurrentPlayerIndex = models.Game().StartPlayerIndex
	models.Game().CurrentPlayer = &models.Game().Players[models.Game().StartPlayerIndex]

	models.Game().Round++
}

func (re EndPhase) EndPhaseCleanup() {
	for i := 0; i < len(models.Game().Players); i++ {
		models.Game().Players[i].PhaseCleanup()
		for j := 0; j < len(models.Game().Players[j].Army.Units); j++ {
			models.Game().Players[i].Army.Units[j].Cleanup()
		}
	}
}
