package initilizer

import (
	"math/rand"
	"time"

	"github.com/kenoyer130/wartgame/consts"
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/phases"
)

func StartGame(g *engine.Game) error {

	engine.WriteMessage("Wartgame!")

	loadAssets(g)

	// todo: hardcoded players for now
	g.Players[0].Name = "playerOne"
	g.Players[0].Army.ID = "test_army"

	g.Players[1].AI = true
	g.Players[1].Name = "SimpleSimon"
	g.Players[1].Army.ID = "ai_army"

	// load player armies
	for i := 0; i < len(g.Players); i++ {
		if err := LoadPlayerArmy(&g.Players[i], g.Assets); err != nil {
			return err
		}
	}

	// for now just place units across from each other
	setPlayerUnitStartingLocation(0, 24, 12, g)
	setPlayerUnitStartingLocation(1, 4, 12, g)

	// roll and set first player
	rand.Seed(time.Now().UnixNano())
	die := rand.Intn(consts.MaxPlayers)

	g.CurrentPlayer = &g.Players[die]
	g.CurrentPlayerIndex = die

	g.Round = 1
	phases.MoveToNextPhaseOrder(models.ShootingPhase_UnitSelection, g)
	return nil
}

func setPlayerUnitStartingLocation(player int, x int, y int, g *engine.Game) {
	for i := 0; i < len(g.Players[player].Army.Units); i++ {

		unit := &g.Players[player].Army.Units[i]

		unit.Location = models.Location{X: x, Y: y}

		engine.SetUnitFormation(engine.StandardUnitFormation, unit, &g.BattleGround)
	}
}
