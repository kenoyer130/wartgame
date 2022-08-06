package phases

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/engine"
)

func StartPhaseShootingWeapons(g *engine.Game) {

	g.SelectedUnit = g.SelectedPhaseUnit

	g.StatusMesssage = "Weapon Selection Phase! Press [Space] to attack with current weapon or [X] to skip!"

	for _, model := range g.SelectedPhaseUnit.Models {
		weapon := model.GetUnfiredWeapon()

		if weapon != "" {
			g.SelectedWeapon = weapon
			break
		}		
	}

	//TODO skip weapons out of range	
	engine.KeyBoardRegistry[ebiten.KeySpace] = func() {
		StartPhaseShootingAttack(g)
	}
}
