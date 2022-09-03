package initilizer

import (
	"math/rand"
	"time"

	"github.com/kenoyer130/wartgame/consts"
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/phases"
	"github.com/kenoyer130/wartgame/weaponabilities"
)

func StartGame() error {

	engine.WriteMessage("Wartgame!")

	loadAssets()

	initGameState()

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
	setPlayerUnitStartingLocation(0, 12, 6)
	setPlayerUnitStartingLocation(1, 2, 6)

	// roll and set first player
	rand.Seed(time.Now().UnixNano())
	die := rand.Intn(consts.MaxPlayers)

	models.Game().CurrentPlayer = &models.Game().Players[die]
	models.Game().CurrentPlayerIndex = die
	models.Game().StartPlayerIndex = die

	opponent := 0
	if die == 0 {
		opponent = 1
	}

	models.Game().OpponetPlayerIndex = opponent

	models.Game().Round = 1
	models.Game().PhaseEventBus.Fire("StartBattleRound")
	return nil
}

func initGameState() {
	models.Game().PhaseEventBus = *models.NewPhaseEventBus()
	models.Game().PhaseStepper = *phases.NewPhaseStepper()
	models.Game().DiceRoller = &engine.DiceRoller{}
	models.Game().Drawer = interfaces.Drawer{}

	wa := weaponabilities.WeaponAbilityList{}
	wa.Init()

	models.Game().WeaponAbilityList = &wa
}

func setPlayerUnitStartingLocation(player int, x int, y int) {
	for i := 0; i < len(models.Game().Players[player].Army.Units); i++ {

		unit := models.Game().Players[player].Army.Units[i]
		unit.SetLocation(interfaces.Location{X: x, Y: y + 1 + i})
	}
}
