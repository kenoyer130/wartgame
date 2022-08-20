package phases

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
)

type ShootingWeaponPhase struct {
	UnitWeapons map[string]*models.ShootingWeapon
	OnCompleted func()
}

func (re ShootingWeaponPhase) GetName() (interfaces.GamePhase, interfaces.PhaseStep) {
	return interfaces.ShootingPhase, interfaces.ShootingPhaseTargeting
}

func (re ShootingWeaponPhase) Start() {

	re.UnitWeapons = make(map[string]*models.ShootingWeapon)

	models.Game().SelectedUnit = models.Game().SelectedPhaseUnit

	models.Game().StatusMessage.Messsage = "Weapon Selection Phase!"
	models.Game().StatusMessage.Keys = "Press [Space] to attack with current weapon or [X] to skip!"

	for _, model := range models.Game().SelectedPhaseUnit.Models {
		for _, weapon := range model.Weapons {
			if re.UnitWeapons[weapon.Name] == nil {
				shootingWeapon := models.ShootingWeapon{Model: *model, Weapon: weapon, Count: 1}
				re.UnitWeapons[weapon.Name] = &shootingWeapon
			} else {
				shootingWeapon := re.UnitWeapons[weapon.Name]
				shootingWeapon.Count++
				re.UnitWeapons[weapon.Name] = shootingWeapon
			}
		}
	}

	re.loop()
}

func (re ShootingWeaponPhase) loop() {

	if len(re.UnitWeapons) == 0 {
		models.Game().SelectedPhaseUnit.AddState(models.UnitShot)
		re.OnCompleted();
		return
	}

	shootingWeapon := models.ShootingWeapon{}

	for _, currentWeapon := range re.UnitWeapons {
		shootingWeapon = *currentWeapon
		break
	}

	models.Game().SelectedWeapon = &shootingWeapon

	//TODO skip weapons out of range
	engine.KeyBoardRegistry[ebiten.KeySpace] = func() {
		shootingAttackPhase := ShootingAttackPhase{}

		shootingAttackPhase.OnCompleted = func() {
			delete(re.UnitWeapons, shootingWeapon.Weapon.Name)
			re.loop()
		}

		shootingAttackPhase.Start()
	}
}
