package phases

import (
	"fmt"

	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

type PhaseStepper struct {
	CurrentPhase         models.GamePhase
	CurrentStep          models.PhaseStep
	CurrentPhaseExecuter models.PhaseExecute
}

func (re PhaseStepper) GetPhaseName() string {
	return fmt.Sprintf("%s %s", re.CurrentPhase, re.CurrentStep)
}

func (re PhaseStepper) Move(phase models.GamePhase, step models.PhaseStep) {

	newPhase := false
	newStep := false

	if re.CurrentPhase == "" {
		newPhase = true
	} else {
		newPhase = (re.CurrentPhase != phase)
	}

	if re.CurrentStep == "" {
		newStep = true
	} else {
		newStep = (re.CurrentStep != step)
	}

	if newPhase && newStep {
		re.cleanupPreviousPhase()	
		re.printPhase(fmt.Sprintf("Starting phase %s %s", phase, step))
	}

	re.CurrentPhaseExecuter.Start()
}

func (re PhaseStepper) printPhase(msg string) {
	engine.WriteMessage(msg)
}

func (re PhaseStepper) cleanupPreviousPhase() {
	models.Game().SelectedPhaseUnit = nil
	models.Game().SelectedTargetUnit = nil
	models.Game().SelectedUnit = nil
	models.Game().StatusMessage.Clear()
}

func (re PhaseStepper) phaseLookup(phase models.GamePhase, step models.PhaseStep) models.PhaseExecute {
	if phase == models.ShootingPhase && step == models.Nil {
		return ShootingPhase{}
	} else if phase == models.ShootingPhase && step == models.ShootingPhaseAttack {
		return ShootingAttackPhase{}
	} else if phase == models.ShootingPhase && step == models.ShootingPhaseTargeting {
		return ShootingTargetingPhase{}
	} else if phase == models.ShootingPhase && step == models.ShootingPhaseWeapons {
		return ShootingWeaponPhase{}
	} else if phase == models.MoralePhase {
		return MoralePhase{}
	} else if phase == models.EndPhase {
		return EndPhase{}
	}

	return nil
}
