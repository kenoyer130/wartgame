package initilizer

import (
	"math/rand"
	"time"

	"github.com/kenoyer130/wartgame/consts"
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/phases"
)

func StartGame() error {

	engine.WriteMessage("Wartgame!")

	loadAssets()

	// todo: hardcoded players for now
	models.Game().Players[0].Name = "playerOne"
	models.Game().Players[0].Army.ID = "test_army"

	models.Game().Players[1].AI = true
	models.Game().Players[1].Name = "SimpleSimon"
	models.Game().Players[1].Army.ID = "ai_army"

	// load player armies
	for i := 0; i < len(models.Game().Players); i++ {
		if err := LoadPlayerArmy(&models.Game().Players[i], models.Game().Assets); err != nil {
			return err
		}
	}

	// for now just place units across from each other
	setPlayerUnitStartingLocation(0, 24, 12)
	setPlayerUnitStartingLocation(1, 4, 12)

	// roll and set first player
	rand.Seed(time.Now().UnixNano())
	die := rand.Intn(consts.MaxPlayers)

	models.Game().CurrentPlayer = &models.Game().Players[die]
	models.Game().CurrentPlayerIndex = die
	models.Game().StartPlayerIndex = die

	models.Game().Round = 1
	phases.MoveToPhase(models.ShootingPhase)
	return nil
}

func setPlayerUnitStartingLocation(player int, x int, y int) {
	for i := 0; i < len(models.Game().Players[player].Army.Units); i++ {

		unit := models.Game().Players[player].Army.Units[i]

		unit.Location = models.Location{X: x, Y: y}

		engine.SetUnitFormation(engine.StandardUnitFormation, unit, &models.Game().BattleGround)
	}
}
