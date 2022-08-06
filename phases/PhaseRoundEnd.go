package phases

import (
	"fmt"

	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func StartRoundEnd() {

	lost := false

	for _, player := range models.Game().Players {
		lost := checkPlayerHasUnits(player)

		if lost {
			break
		}
	}

	if !lost {
		startNextTurn()
	}
}

func checkPlayerHasUnits(player models.Player) bool {
	if len(player.Army.Units) < 1 {
		engine.WriteMessage(fmt.Sprintf("%s player has lost the game due to no units left!", player.Name))
		return true
	}

	return false
}

func startNextTurn() {
	models.Game().CurrentPlayer.Gone = true

	allGone := true

	for i, player := range models.Game().Players {
		if !player.Gone {
			allGone = false
			models.Game().CurrentPlayer = &player
			models.Game().CurrentPlayerIndex = i
		}
	}

	if allGone {
		startNewRound()
	}

	MoveToPhase(models.ShootingPhase)
}

func startNewRound() {
	models.Game().Round++
	models.Game().CurrentPlayerIndex = models.Game().StartPlayerIndex
	models.Game().CurrentPlayer = &models.Game().Players[models.Game().StartPlayerIndex]
	PhaseMoraleCleanup()
}
