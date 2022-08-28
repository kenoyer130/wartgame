package phases

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
)

type ShootingWeaponPhase struct {
	UnitWeapons  map[string]*models.ShootingWeapon
	deductAttack int
	OnCompleted  func()
}

func (re ShootingWeaponPhase) GetName() (interfaces.GamePhase, interfaces.PhaseStep) {
	return interfaces.ShootingPhase, interfaces.ShootingPhaseTargeting
}

func (re ShootingWeaponPhase) Start() {

	re.UnitWeapons = make(map[string]*models.ShootingWeapon)
	re.deductAttack = 0

	models.Game().SelectedUnit = models.Game().SelectedPhaseUnit

	models.Game().StatusMessage.Messsage = "Weapon Selection Phase!"
	models.Game().StatusMessage.Keys = "Press [Space] to attack with current weapon or [X] to skip!"

	for _, model := range models.Game().SelectedPhaseUnit.Models {
		for _, weapon := range model.Weapons {
			if re.UnitWeapons[weapon.Name] == nil {

				inRangeEntities := models.InRange(models.Game().SelectedPhaseUnit.ID, models.Game().CurrentPlayerIndex, weapon.Range, models.Game().SelectedPhaseUnit.Location.X, models.Game().SelectedPhaseUnit.Location.Y)

				if len(inRangeEntities) > 0 {
					shootingWeapon := models.ShootingWeapon{Model: *model, Weapon: weapon, Targets: inRangeEntities, Count: 1}
					re.UnitWeapons[weapon.Name] = &shootingWeapon
				}
			} else {
				shootingWeapon := re.UnitWeapons[weapon.Name]
				shootingWeapon.Count++
				re.UnitWeapons[weapon.Name] = shootingWeapon
			}
		}
	}

	// check for gernades
	for key, currentWeapon := range re.UnitWeapons {
		if currentWeapon.Weapon.WeaponType.Type == "Gre" {
			re.UnitWeapons[key].Count = 1
			re.deductAttack = 1
		}
	}

	re.loop()
}

func (re ShootingWeaponPhase) loop() {

	if len(re.UnitWeapons) == 0 {
		models.Game().SelectedPhaseUnit.AddState(models.UnitShot)
		re.OnCompleted()
		return
	}

	shootingWeapon := models.ShootingWeapon{}

	for _, currentWeapon := range re.UnitWeapons {
		shootingWeapon = *currentWeapon
		break
	}

	models.Game().SelectedWeapon = &shootingWeapon

	if models.Game().SelectedWeapon.Weapon.WeaponType.Type != "Gre" && re.deductAttack > 0 {
		models.Game().SelectedWeapon.Count = models.Game().SelectedWeapon.Count - 1
		if models.Game().SelectedWeapon.Count < 1 {
			models.Game().SelectedWeapon.Count = 1
		}
	}

	engine.WriteMessage("Selected Weapon is " + models.Game().SelectedWeapon.Weapon.Name)

	engine.KeyBoardRegistry[ebiten.KeySpace] = func() {
		shootingTargetingPhase := ShootingTargetingPhase{}

		shootingTargetingPhase.OnCompleted = func() {
			delete(re.UnitWeapons, shootingWeapon.Weapon.Name)
			re.loop()
		}

		models.Game().Players[0].Army.RemoveDestroyedUnits()
		models.Game().Players[1].Army.RemoveDestroyedUnits()

		shootingTargetingPhase.Start()
	}
}
