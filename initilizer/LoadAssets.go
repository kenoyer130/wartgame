package initilizer

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func loadAssets(g *engine.Game) error {

	// init global asset state
	g.Assets = *models.NewAssets()

	// todo: hardcoded battle ground size
	g.BattleGround = *models.NewBattleGround(72, 48)

	err := loadAssetFiles(g)

	return err

}

func loadAssetFiles(g *engine.Game) error {
	err := filepath.Walk("./assets",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			loadByPath(g, path)

			return nil
		})

	return err
}

func loadByPath(g *engine.Game, path string) {

	if isAsset(path, "armies") {
		Unit := loadUnit(path)
		g.Assets.Units[Unit.Name] = Unit
	}

	if isAsset(path, "weapons") {
		loadWeapons(g, path)
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

func loadWeapons(g *engine.Game, path string) {
	var weaponAssets []models.Weapon

	engine.UnmarshalJson(&weaponAssets, path)

	// convert to map for faster lookup
	for i := 0; i < len(weaponAssets); i++ {
		weaponName := weaponAssets[i].Name
		g.Assets.Weapons[weaponName] = weaponAssets[i]
	}
}
