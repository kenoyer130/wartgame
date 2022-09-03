package phases

import (
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
)

type PhaseStepper struct {
	Phase interfaces.GamePhase
}

func (re PhaseStepper) GetPhase() interfaces.GamePhase { 
	return re.Phase
}

func NewPhaseStepper() *PhaseStepper {

	phaseStepper := PhaseStepper{}

	models.Game().PhaseEventBus.RegisterEventHandler("StartBattleRound", func() {
		phaseStepper.Phase = interfaces.MovementPhase
		movePhaseUnitSelector := MovePhaseUnitSelector{}
		movePhaseUnitSelector.Start()
	})

	models.Game().PhaseEventBus.RegisterEventHandler("MovePhaseEnded", func() {
		phaseStepper.Phase = interfaces.ShootingPhase
		phaseStepper.SelectNextShooter()
	})

	models.Game().PhaseEventBus.RegisterEventHandler("ShooterSelected", func() {
		phaseShootingWeaponSelector := PhaseShootingWeaponSelector{}
		phaseShootingWeaponSelector.Start()
	})

	models.Game().PhaseEventBus.RegisterEventHandler("ShooterWeaponSelected", func() {
		phaseShootingTargetting := PhaseShootingTargetting{}
		phaseShootingTargetting.Start()
	})

	models.Game().PhaseEventBus.RegisterEventHandler("ShooterWeaponTargetSelected", func() {
		phaseShootingAttackResolver := PhaseShootingAttackResolver{}
		phaseShootingAttackResolver.Start()
	})

	models.Game().PhaseEventBus.RegisterEventHandler("ShooterAttackCompleted", func() {
		phaseStepper.SelectNextShooter()
	})

	models.Game().PhaseEventBus.RegisterEventHandler("ShooterPhaseEnded", func() {
		phaseStepper.Phase = interfaces.MoralePhase
		moralePhase := MoralePhase{}
		moralePhase.Start()
	})

	models.Game().PhaseEventBus.RegisterEventHandler("MoralePhaseEnded", func() {
		phaseStepper.Phase = interfaces.EndPhase
		endPhase := EndPhase{}
		endPhase.Start()
	})

	return &phaseStepper
}

func (re PhaseStepper) SelectNextShooter() {
	phaseShootingUnitSelector := PhaseShootingUnitSelector{}
	phaseShootingUnitSelector.Start()
}
