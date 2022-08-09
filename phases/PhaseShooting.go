package phases

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

type ShootingPhase struct {
	ShootingTargetingPhase *ShootingTargetingPhase
	ShootingAttackPhase    *ShootingAttackPhase
	ShootingWeaponPhase    *ShootingWeaponPhase
}

func (re ShootingPhase) GetName() (models.GamePhase, models.PhaseStep) {
	return models.ShootingPhase, models.Nil
}

func (re ShootingPhase) Start() {

	if re.ShootingTargetingPhase == nil {
		re.init()
	}

	models.Game().StatusMessage.Phase = "Shooting Phase"
	models.Game().StatusMessage.Messsage = "Select next unit to shoot!"
	models.Game().StatusMessage.Keys = "Press [Q] and [E] to cycle units! Press [Space] to select!"

	unitCycler := NewUnitCycler(models.Game().CurrentPlayer, re.UnitIsValidShooter, re.ShooterSelected)

	unitCycler.CycleUnits()
}

func (re *ShootingPhase) init() {
	re.ShootingTargetingPhase = &ShootingTargetingPhase{}
	re.ShootingWeaponPhase = &ShootingWeaponPhase{}
	re.ShootingAttackPhase = &ShootingAttackPhase{}

	// this sets up our combat cycle
	re.ShootingTargetingPhase.ShootingWeaponPhase = re.ShootingWeaponPhase
	re.ShootingWeaponPhase.ShootingAttackPhase = re.ShootingAttackPhase
	re.ShootingAttackPhase.ShootingPhase = re
}

func (re ShootingPhase) UnitIsValidShooter(unit *models.Unit) bool {
	return unit.CanShoot()
}

func (re ShootingPhase) ShooterSelected(unit *models.Unit) {
	if unit == nil {
		engine.WriteMessage("No valid units for shooting phase.")
		models.Game().PhaseStepper.Move(models.MoralePhase)
		return
	}

	models.Game().SelectedPhaseUnit = unit
	engine.WriteMessage("Selected Unit to shoot: " + unit.Name)

	re.ShootingTargetingPhase.Start()
}
