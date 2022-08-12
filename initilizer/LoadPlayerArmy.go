package initilizer

import (
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func LoadPlayerArmy(p *models.Player, assets models.Assets) error {
	// verify profile army exists
	path := fmt.Sprintf("./player_profile/%s/%s.json", p.Name, p.Army.ID)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return err
	}

	// load it from json
	var army models.Army
	engine.UnmarshalJson(&army, path)

	// copy Units and weapons from asset library (this allows customizations and choices per player)

	for u, Unit := range army.Units {

		_, ok := assets.Units[Unit.Name]

		if !ok {
			engine.Warn(fmt.Sprintf("%s Unit id not found!", Unit.Name))
			continue
		}

		loadedModels := []*models.Model{}

		// break out Models into individual
		for _, Model := range Unit.Models {
			count := Model.Count

			for i := 0; i < count; i++ {

				for _, assetModel := range assets.Units[Unit.Name].Models {

					model := *assetModel

					model.CurrentWounds = model.Wounds
					model.ID = uuid.New().String()

					for _, weaponKey := range Model.DefaultWeapons {
						model.Weapons = append(Model.Weapons , assets.Weapons[weaponKey])
					}

					// find matching asset model
					if assetModel.Name == Model.Name {
						loadedModels = append(loadedModels, &model)
					}
				}
			}
		}

		army.Units[u].Models = loadedModels
		army.Units[u].OriginalModelCount = len(army.Units[u].Models)
	}

	p.Army = army

	models.Game().GameStateUpdater = models.GameStateUpdater{}
	models.Game().GameStateUpdater.Init()

	return nil
}
