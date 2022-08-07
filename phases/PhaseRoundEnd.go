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

	if allGone {
		startNewRound()
	}

	EndPhaseCleanup()

	MoveToPhase(models.ShootingPhase)
}

func startNewRound() {

	for i := 0; i < len(models.Game().Players); i++ {
		models.Game().Players[i].RoundCleanup()
	}

	models.Game().CurrentPlayerIndex = models.Game().StartPlayerIndex
	models.Game().CurrentPlayer = &models.Game().Players[models.Game().StartPlayerIndex]

	models.Game().Round++

}

func EndPhaseCleanup() {
	for i := 0; i < len(models.Game().Players); i++ {
		models.Game().Players[i].PhaseCleanup()
		for j := 0; j < len(models.Game().Players[j].Army.Units); j++ {
			models.Game().Players[i].Army.Units[j].Cleanup()
		}
	}
}
