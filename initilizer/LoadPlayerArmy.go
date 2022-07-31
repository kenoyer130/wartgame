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

	// copy squads and weapons from asset library (this allows customizations and choices per player)

	for i, squad := range army.Squads {

		assetSquad, ok := assets.Squads[squad.Name]

		if !ok {
			engine.Warn(fmt.Sprintf("%s squad id not found!", squad.Name))
			continue
		}

		armySquad := assetSquad

		army.Squads[i] = armySquad

		// break out units into individual
		for _, unit := range assetSquad.Units {
			count := unit.UnitNumber.Min

			for i := 0; i < count-1; i++ {

				squadUnit := unit
				armySquad.Units = append(armySquad.Units, squadUnit)
			}
		}

		assignWeapons(&armySquad, &assetSquad, assets.Weapons)
	}

	p.Army = army

	return nil
}

func assignWeapons(squad *models.Squad, assetSquad *models.Squad, assetWeapons map[string]models.Weapon) {

	for _, unit := range squad.Units {
		assignUnitWeapons(assetSquad, assetWeapons, unit)
	}
}

func assignUnitWeapons(assetSquad *models.Squad, assetWeapons map[string]models.Weapon, unit models.Unit) {
	for _, assetWeaponId := range assetSquad.DefaultWeapons {

		assetWeapon, ok := assetWeapons[assetWeaponId]

		if !ok {
			engine.Warn(fmt.Sprintf("%s weapon id not found!", assetWeaponId))
			continue
		}

		unit.Weapons = append(unit.Weapons, assetWeapon)
	}
}
