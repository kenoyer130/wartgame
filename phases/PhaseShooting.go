package phases

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
)

type ShootingPhase struct {
	ShootingTargetingPhase *ShootingTargetingPhase
	ShootingAttackPhase    *ShootingAttackPhase
	ShootingWeaponPhase    *ShootingWeaponPhase
}

func (re ShootingPhase) GetName() (interfaces.GamePhase, interfaces.PhaseStep) {
	return interfaces.ShootingPhase, interfaces.Nil
}

func (re ShootingPhase) Start() {

	models.Game().StatusMessage.Phase = "Shooting Phase"
	models.Game().StatusMessage.Messsage = "Select next unit to shoot!"
	models.Game().StatusMessage.Keys = "Press [Q] and [E] to cycle units! Press [Space] to select!"
	re.loop()
}

func (re ShootingPhase) loop() {
	unitCycler := NewUnitCycler(models.Game().CurrentPlayer, re.UnitIsValidShooter, re.ShooterSelected, false)
	unitCycler.CycleUnits()
}

func (re ShootingPhase) UnitIsValidShooter(unit *models.Unit) bool {
	return unit.CanShoot()
}

func (re ShootingPhase) ShooterSelected(unit *models.Unit) {
	if unit == nil {
		engine.WriteMessage("No valid units for shooting phase.")
		models.Game().PhaseStepper.Move(interfaces.MoralePhase)
		return
	}

	models.Game().SelectedPhaseUnit = unit
	engine.WriteMessage("Selected Unit to shoot: " + unit.Name)

	shootingTargetingPhase := ShootingTargetingPhase{}
	shootingTargetingPhase.OnCompleted = func() {
		re.loop()
	}
	shootingTargetingPhase.Start()
}
