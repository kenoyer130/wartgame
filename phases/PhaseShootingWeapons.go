package phases

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
)

type ShootingWeaponPhase struct {
	ShootingAttackPhase *ShootingAttackPhase
	ShootingPhase       *ShootingPhase
}

func (re ShootingWeaponPhase) GetName() (interfaces.GamePhase, interfaces.PhaseStep) {
	return interfaces.ShootingPhase, interfaces.ShootingPhaseTargeting
}

func (re ShootingWeaponPhase) Start() {

	models.Game().SelectedUnit = models.Game().SelectedPhaseUnit

	models.Game().StatusMessage.Messsage = "Weapon Selection Phase!"
	models.Game().StatusMessage.Keys = "Press [Space] to attack with current weapon or [X] to skip!"

	models.Game().SelectedWeaponName = ""

	for _, model := range models.Game().SelectedPhaseUnit.Models {
		weapon := model.GetUnfiredWeapon()

		if weapon != nil {
			models.Game().SelectedWeaponName = weapon.Name
			for i := 0; i < len(models.Game().SelectedPhaseUnit.Models); i++ {
				models.Game().SelectedPhaseUnit.Models[i].SetFiredWeapon(weapon)
			}

			break
		}
	}

	if models.Game().SelectedWeaponName == "" {
		models.Game().SelectedPhaseUnit.AddState(models.UnitShot)
		re.ShootingPhase.Start()
		return
	}

	//TODO skip weapons out of range
	engine.KeyBoardRegistry[ebiten.KeySpace] = func() {
		re.ShootingAttackPhase.Start()
	}
}
