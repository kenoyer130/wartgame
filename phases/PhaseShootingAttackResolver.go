package phases

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
)

type PhaseShootingAttackResolver struct {
	shooter     models.Model
	targetUnit  *models.Unit
	targetName  string
	weapon      *models.Weapon
	weaponCount int
	attacks     int
	hits        int
	wounds      int
	saves       int
	kills       int
}

func (re PhaseShootingAttackResolver) GetName() (interfaces.GamePhase, interfaces.PhaseStep) {
	return interfaces.ShootingPhase, interfaces.Nil
}

func (re *PhaseShootingAttackResolver) Start() {

	re.targetUnit = models.Game().SelectedTargetUnit
	re.targetName = re.targetUnit.Name
	re.shooter = models.Game().SelectedWeapon.Model

	// figure out how many weapons are making the attack
	re.weaponCount = models.Game().SelectedWeapon.Weapon.WeaponType.GetAttacks(models.Game().SelectedWeapon.Count, models.Game().SelectedTargetUnit, *models.Game().SelectedWeapon)

	re.attacks = 0
	re.hits = 0
	re.wounds = 0
	re.saves = 0
	re.kills = 0

	engine.WriteMessage(fmt.Sprintf("Firing %d %s at %s ", re.weaponCount, models.Game().SelectedWeapon.Weapon.Name, re.targetUnit.Name))

	max := re.weaponCount
	re.attacks = max

	models.Game().DiceRoller.PlaySound()

	engine.PlaySound(models.Game().SelectedWeapon.Weapon.Name)

	for i := 0; i < max; i++ {
		re.setShootingWeapon(i, max)

		if re.targetUnit == nil {
			break
		}
	}

	re.endShootingWeaponPhase()
}

func (re *PhaseShootingAttackResolver) setShootingWeapon(i int, max int) {
	selectedWeapon := *models.Game().SelectedWeapon
	re.weapon = &selectedWeapon.Weapon
	re.shootWeapon(i, max)
}

func (re *PhaseShootingAttackResolver) endShootingWeaponPhase() {

	engine.WriteMessage(fmt.Sprintf("%s %s attack on %s completed.", re.shooter.Name, re.weapon.Name, re.targetName))
	engine.WriteMessage(fmt.Sprintf("%d attacks %d hits, %d wounds, %d saves, %d killed!", re.attacks, re.hits, re.wounds, re.saves, re.kills))
	engine.WriteMessage("Press [Space] to continue")
	engine.KeyBoardRegistry[ebiten.KeySpace] = func() {
		models.Game().PhaseEventBus.Fire("ShooterAttackCompleted")
	}
}

func (re *PhaseShootingAttackResolver) shootWeapon(i int, max int) {
	engine.WriteMessage(fmt.Sprintf("%s firing %s %d of %d attacks", re.shooter.Name, re.weapon.Name, i+1, max))

	hits, _ := models.Game().DiceRoller.Roll("Rolling for Attack",
		interfaces.DiceRollType{
			Dice:   1,
			Target: re.shooter.GetBallisticSkill(),
		},
		nil)

	if hits == 0 {
		engine.WriteMessage(fmt.Sprintf("%s attack failed.", re.shooter.Name))
		return
	}

	re.onAttackRolled()
}

func (re *PhaseShootingAttackResolver) onAttackRolled() {

	re.hits++
	target := &models.Model{}

	for _, model := range re.targetUnit.Models {
		if model.CurrentWounds < model.Wounds {
			target = model
			break
		}
	}

	if target.ID == "" {
		target = re.targetUnit.Models[rand.Intn(len(re.targetUnit.Models))]
	}

	re.rollWoundsToModel(*target)
}

func (re *PhaseShootingAttackResolver) rollWoundsToModel(target models.Model) {

	toughnessTarget := getWoundTarget(re.weapon.Strength, target.Toughness)

	engine.WriteMessage("Wound target is " + strconv.Itoa(toughnessTarget))

	wounds, _ := models.Game().DiceRoller.Roll("Rolling for Wounds", interfaces.DiceRollType{
		Dice:   1,
		Target: toughnessTarget,
	},
		re.onWoundDie)

	if wounds == 0 {
		engine.WriteMessage(fmt.Sprintf("No Wound. %s attack failed.", re.shooter.Name))
		return
	}

	re.wounds++

	re.onWoundsRolled(target)
}

func (re *PhaseShootingAttackResolver) onWoundsRolled(target models.Model) {

	ap := re.weapon.ArmorPiercing

	save := target.GetIntSkill(target.Save)

	success, _ := models.Game().DiceRoller.Roll("Rolling for Save", interfaces.DiceRollType{
		Dice:      1,
		Target:    save,
		AddToDice: ap,
	},
		nil)

	if success < 1 {
		re.allocateAttacks(target)
	} else {
		re.saves++
		engine.WriteMessage(fmt.Sprintf("%s made save!", target.Name))
		engine.WriteMessage(fmt.Sprintf("%s attack failed.", re.shooter.Name))
	}
}

func (re *PhaseShootingAttackResolver) onWoundDie(die int) int {

	models.Game().WeaponAbilityList.ApplyWeaponAbilities(interfaces.WeaponAbilityPhaseWounds, die, re.weapon)
	return die
}

func (re *PhaseShootingAttackResolver) allocateAttacks(target models.Model) {

	engine.WriteMessage(fmt.Sprintf("%s failed save!", target.Name))

	re.InflictWounds(target)

	if len(re.targetUnit.Models) <= 0 {
		re.targetUnit.Destroyed = true
		models.Game().BattleGround.RemoveBattleGroundEntity(re.targetUnit)
		opponent := models.Game().OpponetPlayerIndex
		models.Game().Players[opponent].Army.RemoveDestroyedUnits()
		re.targetUnit = nil
	}
}

func (re *PhaseShootingAttackResolver) InflictWounds(target models.Model) {

	dmg := re.weapon.Damage
	dead := false

	engine.WriteMessage(fmt.Sprintf("inflicting %d damage wound on %s!", dmg, target.Name))
	dead, deadModel := re.targetUnit.InflictWounds(target, dmg)

	if dead {
		re.kills++
		engine.WriteMessage(fmt.Sprintf("%s was destroyed!", deadModel.Name))
		if re.targetUnit.Destroyed {
			engine.WriteMessage(fmt.Sprintf("Unit %s wiped out!", re.targetUnit.Name))
		}
	}
}

func getWoundTarget(str int, toughness int) int {
	if str > toughness*2 {
		return 2
	}

	if str > toughness {
		return 3
	}

	if str == toughness {
		return 4
	}

	if str < toughness {
		return 5
	}

	if (str) < toughness/2 {
		return 6
	}

	return 4
}
