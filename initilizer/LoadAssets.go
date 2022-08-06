package initilizer

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func loadAssets() error {

	// init global asset state
	models.Game().Assets = *models.NewAssets()

	// todo: hardcoded battle ground size
	models.Game().BattleGround = *models.NewBattleGround(72, 48)

	err := loadAssetFiles()

	return err

}

func loadAssetFiles() error {
	err := filepath.Walk("./assets",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			loadByPath(path)

			return nil
		})

	return err
}

func loadByPath(path string) {

	if isAsset(path, "armies") {
		Unit := loadUnit(path)
		models.Game().Assets.Units[Unit.Name] = Unit
	}

	if isAsset(path, "weapons") {
		loadWeapons(path)
	}
}

func isAsset(path string, asset string) bool {
	return strings.HasPrefix(path, "assets\\"+asset) && strings.HasSuffix(path, ".json")
}

func loadUnit(path string) models.Unit {
	var Unit models.Unit
	engine.UnmarshalJson(&Unit, path)
	return Unit
}

func loadWeapons(path string) {
	var weaponAssets []models.Weapon

	engine.UnmarshalJson(&weaponAssets, path)

	// convert to map for faster lookup
	for i := 0; i < len(weaponAssets); i++ {
		weaponName := weaponAssets[i].Name
		models.Game().Assets.Weapons[weaponName] = weaponAssets[i]
	}
}
