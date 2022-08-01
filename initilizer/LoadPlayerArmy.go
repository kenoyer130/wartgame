package initilizer

import (
	"errors"
	"fmt"
	"os"

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

	for i, Unit := range army.Units {

		assetUnit, ok := assets.Units[Unit.Name]

		if !ok {
			engine.Warn(fmt.Sprintf("%s Unit id not found!", Unit.Name))
			continue
		}

		armyUnit := assetUnit

		// break out Models into individual
		for _, Model := range assetUnit.Models {
			count := Model.ModelNumber.Min

			for i := 0; i < count-1; i++ {

				UnitModel := Model
				armyUnit.Models = append(armyUnit.Models, UnitModel)
			}
		}

		assignWeapons(&armyUnit, &assetUnit, assets.Weapons)

		army.Units[i] = armyUnit

	}

	p.Army = army

	return nil
}

func assignWeapons(Unit *models.Unit, assetUnit *models.Unit, assetWeapons map[string]models.Weapon) {

	for _, Model := range Unit.Models {
		assignModelWeapons(assetUnit, assetWeapons, Model)
	}
}

func assignModelWeapons(assetUnit *models.Unit, assetWeapons map[string]models.Weapon, Model models.Model) {
	for _, assetWeaponId := range assetUnit.DefaultWeapons {

		assetWeapon, ok := assetWeapons[assetWeaponId]

		if !ok {
			engine.Warn(fmt.Sprintf("%s weapon id not found!", assetWeaponId))
			continue
		}

		Model.Weapons = append(Model.Weapons, assetWeapon)
	}
}
