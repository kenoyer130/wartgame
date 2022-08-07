package phases

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func StartPhaseShootingWeapons() {

	models.Game().SelectedUnit = models.Game().SelectedPhaseUnit

	models.Game().StatusMessage.Messsage = "Weapon Selection Phase!"
	models.Game().StatusMessage.Keys = "Press [Space] to attack with current weapon or [X] to skip!"

	models.Game().SelectedWeaponName = ""

	for _, model := range models.Game().SelectedPhaseUnit.Models {
		weapon := model.GetUnfiredWeapon()

		if weapon != "" {
			models.Game().SelectedWeaponName = weapon
				for i := 0; i < len(models.Game().SelectedPhaseUnit.Models); i++ {
					models.Game().SelectedPhaseUnit.Models[i].SetFiredWeapon(weapon)
				}			
			
			break
		}		
	}

	if(models.Game().SelectedWeaponName == "") {		
		models.Game().SelectedPhaseUnit.AddState(models.UnitShot)
		StartPhaseShooting()
		return
	}

	//TODO skip weapons out of range	
	engine.KeyBoardRegistry[ebiten.KeySpace] = func() {		
		StartPhaseShootingAttack()
	}
}
