package phases

import (
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

type ShootingPhase struct {
}

func (re ShootingPhase) GetName() (models.GamePhase, models.PhaseStep) {
	return models.ShootingPhase, models.Nil
}

func (re ShootingPhase) Start() {

	models.Game().StatusMessage.Phase = "Shooting Phase"
	models.Game().StatusMessage.Messsage = "Select next unit to shoot!"
	models.Game().StatusMessage.Keys = "Press [Q] and [E] to cycle units! Press [Space] to select!"

	unitCycler := NewUnitCycler(models.Game().CurrentPlayer, re.UnitIsValidShooter, re.ShooterSelected)

	unitCycler.CycleUnits()
}

func (re ShootingPhase) UnitIsValidShooter(unit *models.Unit) bool {
	return unit.CanShoot()
}

func (re ShootingPhase) ShooterSelected(unit *models.Unit) {
	if unit == nil {
		engine.WriteMessage("No valid units for shooting phase.")
		models.Game().PhaseStepper.Move(models.MoralePhase, models.Nil)
		return
	}

	models.Game().SelectedPhaseUnit = unit
	engine.WriteMessage("Selected Unit to shoot: " + unit.Name)

	models.Game().PhaseStepper.Move(models.ShootingPhase, models.ShootingPhaseAttack)
}
