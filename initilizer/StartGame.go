package initilizer

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func StartGame(g *engine.Game) error {

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
	for _, squad := range g.Players[0].Army.Squads {

		squad.Location = models.Location{ X: 24, Y: 24}
	}

	return nil
}
