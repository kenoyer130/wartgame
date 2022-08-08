package testutils

import "github.com/kenoyer130/wartgame/models"

func InitGameState() {
	models.Game().DiceRoller = DiceRollerFake{}
	models.Game().PhaseStepper = PhaseStepperFake{}
	initDate()
}

func initDate() {

	models.Game().Assets = *models.NewAssets()
	models.Game().Assets.Weapons["testW1"] = models.Weapon{
		Name: "testW1",
	}

	models.Game().Assets.Weapons["testW2"] = models.Weapon{
		Name: "testW2",
	}

	models.Game().SelectedWeapons = []models.Weapon{
		{
			Damage: 1,
		}}
}
