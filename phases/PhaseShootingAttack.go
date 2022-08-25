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
	shooter     models.Model
	targetUnit  *models.Unit
	weapon      *models.Weapon
	weaponCount int
	OnCompleted func()
	hits        int
	kills       int
}

func (re ShootingAttackPhase) GetName() (interfaces.GamePhase, interfaces.PhaseStep) {
	return interfaces.ShootingPhase, interfaces.Nil
}

func (re *ShootingAttackPhase) Start() {

	// figure out how many weapons are making the attack
	re.weaponCount = models.Game().SelectedWeapon.Count

	if models.Game().SelectedWeapon.Weapon.WeaponType.Type == "Gre" {
		for _, weapon := range models.Game().SelectedPhaseUnit.Models {
			if weapon.ModelType == "Gre" {
				re.weaponCount--
			}
		}
	}

	if models.Game().SelectedWeapon.Weapon.WeaponType.Type == "Gre" {
		re.weaponCount = re.throwGernade()
	}

	if models.Game().SelectedWeapon.Weapon.WeaponType.Number > 0 {
		re.weaponCount = re.weaponCount * models.Game().SelectedWeapon.Weapon.WeaponType.Number
	}

	re.targetUnit = models.Game().SelectedTargetUnit
	re.shooter = models.Game().SelectedWeapon.Model

	re.hits = 0
	re.kills = 0

	engine.WriteMessage(fmt.Sprintf("Firing %d %s at %s ", re.weaponCount, models.Game().SelectedWeapon.Weapon.Name, re.targetUnit.Name))

	max := re.weaponCount

	models.Game().DiceRoller.PlaySound()

	for i := 0; i < max; i++ {
		re.setShootingWeapon(i, max)

		if len(re.targetUnit.Models) <= 0 {
			break
		}
	}

	if re.targetUnit.Destroyed {
		engine.WriteMessage(fmt.Sprintf("Unit %s wiped out!", re.targetUnit.Name))
	}

	re.endShootingWeaponPhase()
}

func (re *ShootingAttackPhase) throwGernade() int {
	die := rand.Intn(models.Game().SelectedWeapon.Weapon.WeaponType.Dice) + 1

	modelCount := len(models.Game().SelectedTargetUnit.Models)

	if modelCount > 10 {
		die = models.Game().SelectedWeapon.Weapon.WeaponType.Dice
	} else if modelCount > 5 && die < 4 {
		die = 3
	}

	return die
}

func (re *ShootingAttackPhase) setShootingWeapon(i int, max int) {
	selectedWeapon := *models.Game().SelectedWeapon
	re.weapon = &selectedWeapon.Weapon
	re.shootWeapon(i, max)
}

func (re *ShootingAttackPhase) endShootingWeaponPhase() {

	engine.WriteMessage(fmt.Sprintf("%s %d hits and %d destroyed!", re.weapon.Name, re.hits, re.kills))
	engine.WriteMessage("Press [Space] to continue")
	engine.KeyBoardRegistry[ebiten.KeySpace] = func() {
		re.OnCompleted()
	}
}

func (re *ShootingAttackPhase) shootWeapon(i int, max int) {
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

func (re *ShootingAttackPhase) onAttackRolled() {

	re.hits++
	target := re.targetUnit.Models[rand.Intn(len(re.targetUnit.Models))]
	re.rollWoundsToModel(*target)
}

func (re *ShootingAttackPhase) rollWoundsToModel(target models.Model) {

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

	re.onWoundsRolled(target)
}

func (re *ShootingAttackPhase) onWoundsRolled(target models.Model) {

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
		engine.WriteMessage(fmt.Sprintf("%s made save!", target.Name))
		engine.WriteMessage(fmt.Sprintf("%s attack failed.", re.shooter.Name))
	}
}

func (re *ShootingAttackPhase) onWoundDie(die int) int {

	models.Game().WeaponAbilityList.ApplyWeaponAbilities(interfaces.WeaponAbilityPhaseWounds, die, re.weapon)
	return die
}

func (re *ShootingAttackPhase) allocateAttacks(target models.Model) {

	engine.WriteMessage(fmt.Sprintf("%s failed save!", target.Name))

	re.InflictWounds(target)

	if len(re.targetUnit.Models) <= 0 {
		re.targetUnit.Destroyed = true
		opponent := 0
		if models.Game().CurrentPlayerIndex == 0 {
			opponent = 1
		}

		models.Game().Players[opponent].Army.RemoveDestroyedUnits()
	}
}

func (re *ShootingAttackPhase) InflictWounds(target models.Model) bool {

	dmg := re.weapon.Damage
	dead := false

	engine.WriteMessage(fmt.Sprintf("inflicting %d damage wound on %s!", dmg, target.Name))
	dead, deadModel := re.targetUnit.InflictWounds(target, dmg)

	if dead {
		re.kills++
		engine.WriteMessage(fmt.Sprintf("%s was destroyed!", deadModel.Name))
		return dead
	}

	return dead
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
