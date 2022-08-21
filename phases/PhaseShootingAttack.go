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

type ShootingAttackPhase struct {
	shooter            models.Model
	targetUnit         models.Unit
	weapon             models.Weapon
	weaponCount        int
	OnCompleted        func()
	onShootingComplete func()
	hits               int
	kills              int
}

func (re ShootingAttackPhase) GetName() (interfaces.GamePhase, interfaces.PhaseStep) {
	return interfaces.ShootingPhase, interfaces.Nil
}

func (re *ShootingAttackPhase) Start() {

	// figure out how many weapons are making the attack
	re.weaponCount = models.Game().SelectedWeapon.Count
	re.targetUnit = *models.Game().SelectedTargetUnit
	re.shooter = models.Game().SelectedWeapon.Model

	engine.WriteMessage(fmt.Sprintf("Firing %d %s at %s ", re.weaponCount, models.Game().SelectedWeapon.Weapon.Name, re.targetUnit.Name))

	i := 0
	max := re.weaponCount

	models.Game().DiceRoller.Suppress(true)

	re.onShootingComplete = func() {
		i++
		if i < max {
			re.setShootingWeapon()
			return
		}

		re.endShootingWeaponPhase()
	}

	re.setShootingWeapon()
}

func (re *ShootingAttackPhase) setShootingWeapon() {
	selectedWeapon := *models.Game().SelectedWeapon
	re.weapon = selectedWeapon.Weapon
	re.shootWeapon()
}

func (re *ShootingAttackPhase) endShootingWeaponPhase() {
	models.Game().DiceRoller.Suppress(false)

	engine.WriteMessage(fmt.Sprintf("%d hits and %d destroyed!", re.hits, re.kills))
	engine.WriteMessage("Press [Space] to continue")
	engine.KeyBoardRegistry[ebiten.KeySpace] = func() {
		re.OnCompleted()
	}
}

func (re ShootingAttackPhase) shootWeapon() {
	engine.WriteMessage(fmt.Sprintf("%s firing %s ", re.shooter.Name, re.weapon.Name))
	models.Game().DiceRoller.Roll("Rolling for Attack",
		interfaces.DiceRollType{
			Dice:   1,
			Target: re.shooter.GetBallisticSkill(),
		},
		nil,
		func(success int, dice []int) {
			re.onAttackRolled(success)
		})
}

func (re *ShootingAttackPhase) onAttackRolled(hit int) {

	if hit == 0 {
		re.onShootingComplete()
		return
	}

	re.hits++
	target := re.targetUnit.Models[rand.Intn(len(re.targetUnit.Models))]
	re.rollWoundsToModel(*target)
}

func (re ShootingAttackPhase) rollWoundsToModel(target models.Model) {

	toughnessTarget := getWoundTarget(re.shooter.Strength, target.Toughness)

	engine.WriteMessage("Wound target is " + strconv.Itoa(toughnessTarget))

	models.Game().DiceRoller.Roll("Rolling for Wounds", interfaces.DiceRollType{
		Dice:   1,
		Target: toughnessTarget,
	},
		re.onWoundDie,
		func(success int, dice []int) {

			if success == 0 {
				re.onShootingComplete()
				return
			}

			re.onWoundsRolled(target)
		})
}

func (re ShootingAttackPhase) onWoundsRolled(target models.Model) {

	ap := re.weapon.ArmorPiercing

	save := target.GetIntSkill(target.Save) - ap

	models.Game().DiceRoller.Roll("Rolling for Save", interfaces.DiceRollType{
		Dice:   1,
		Target: save,
	},
		nil,
		func(success int, dice []int) {
			if success > 0 {
				engine.WriteMessage(fmt.Sprintf("%s made save!", target.Name))
				re.onShootingComplete()
				return
			}

			re.allocateAttacks(success, target)
		})
}

func (re *ShootingAttackPhase) onWoundDie(die int) int {

	weapon := models.Game().WeaponAbilityList.ApplyWeaponAbilities(interfaces.WeaponAbilityPhaseWounds, die, re.weapon)
	re.weapon = weapon.(models.Weapon)
	return die
}

func (re ShootingAttackPhase) allocateAttacks(success int, target models.Model) {

	engine.WriteMessage(fmt.Sprintf("%s failed save!", target.Name))
	re.InflictWounds(target)

	if len(re.targetUnit.Models) <= 0 {
		engine.WriteMessage(fmt.Sprintf("Unit %s wiped out!", re.targetUnit.Name))
		re.targetUnit.Destroyed = true
		re.onShootingComplete()

	} else {
		engine.WriteMessage(fmt.Sprintf("%s took %d casulties!", re.targetUnit.Name, len(re.targetUnit.DestroyedModels)))
		re.onShootingComplete()
	}
}

func (re *ShootingAttackPhase) InflictWounds(target models.Model) {

	dmg := re.weapon.Damage
	dead := false

	engine.WriteMessage(fmt.Sprintf("inflicting %d damage wound on %s!", dmg, target.Name))
	dead, deadModel := models.Game().SelectedTargetUnit.InflictWounds(target, dmg)

	if dead {
		re.kills++
		engine.WriteMessage(fmt.Sprintf("%s was destroyed!", deadModel.Name))
	}

	re.onShootingComplete()

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
