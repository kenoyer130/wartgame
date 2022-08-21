package phases

import (
	"fmt"
	"math"
	"strconv"

	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/interfaces"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/weaponabilities"
)

type ShootingAttackPhase struct {
	ShootingWeaponPhase *ShootingWeaponPhase
	model               models.Model
	weapon              models.Weapon
	weaponCount         int
	OnCompleted         func()
	TargetUnits         models.Stack
	WeaponAttacks       int
}

func (re ShootingAttackPhase) GetName() (interfaces.GamePhase, interfaces.PhaseStep) {
	return interfaces.ShootingPhase, interfaces.Nil
}

func (re ShootingAttackPhase) Start() {

	// figure out how many weapons are making the attack
	re.model = models.Game().SelectedWeapon.Model
	re.weapon = models.Game().SelectedWeapon.Weapon
	re.weaponCount = models.Game().SelectedWeapon.Count

	re.TargetUnits = *re.setModelsByToughness()

	models.Game().StatusMessage.Messsage = fmt.Sprintf("%s is shooting %s with %d %s ", models.Game().SelectedPhaseUnit.Name, models.Game().SelectedTargetUnit.Name, re.weaponCount, re.weapon.Name)

	shots := weaponabilities.ApplyWeaponAbilityShot(re.weapon)
	re.WeaponAttacks = shots

	re.shootWeapon(re.weaponCount*shots, re.model)
}

// break apart the unit models into toughness brackets for quicker rolling
func (re ShootingAttackPhase) setModelsByToughness() *models.Stack {

	targetUnits := models.Stack{}

	toughCheck := -1
	group := models.Stack{}

	count := len(models.Game().SelectedTargetUnit.Models)

	for i := 0; i < count; i++ {
		checkModel := models.Game().SelectedTargetUnit.Models[i]
		toughness := checkModel.Toughness

		if toughCheck == -1 {
			toughCheck = toughness
			group.Push(checkModel)

		} else if toughness != toughCheck {
			targetUnits.Push(group)
			group = models.Stack{}
			group.Push(checkModel)

		} else {
			group.Push(checkModel)
		}

		if i == count-1 {
			targetUnits.Push(group)
		}
	}

	randomizeUnit(&targetUnits)

	return &targetUnits
}

func randomizeUnit(targetUnits *models.Stack) {
	targetUnits.Randomize()

	for i, _ := range targetUnits.Array() {
		modelPeek := targetUnits.Array()[i]
		modelStack := modelPeek.(models.Stack)
		modelStack.Randomize()
		targetUnits.Array()[i] = modelStack
	}
}

func (re ShootingAttackPhase) shootWeapon(weaponCount int, model models.Model) {
	models.Game().DiceRoller.Roll("Rolling for Attack",
		models.DiceRollType{
			Dice:   weaponCount,
			Target: model.GetBallisticSkill(),
		}, func(success int, dice []int) {
			re.onAttackRolled(success, &model)
		})
}

func (re ShootingAttackPhase) onAttackRolled(hits int, model *models.Model) {
	target, _ := re.TargetUnits.Pop()
	re.rollWoundsToUnit(model, hits, target.(models.Stack))
}

func (re ShootingAttackPhase) rollWoundsToUnit(model *models.Model, hits int, targetModels models.Stack) {

	peeked, _ := targetModels.Peek()

	targetModel := peeked.(*models.Model)

	target := getWoundTarget(model.Strength, targetModel.Toughness)

	engine.WriteMessage("Wound target is " + strconv.Itoa(target))

	models.Game().DiceRoller.Roll("Rolling for Wounds", models.DiceRollType{
		Dice:   hits,
		Target: target,
	}, func(success int, dice []int) {
		re.onWoundsRolled(success, targetModels, model)
	})
}

func (re ShootingAttackPhase) onWoundsRolled(hits int, targetModels models.Stack, model *models.Model) {

	if hits == 0 {
		re.OnCompleted()
		return
	}

	ap := re.weapon.ArmorPiercing

	peeked, _ := targetModels.Peek()

	targetModel := peeked.(*models.Model)

	save := targetModel.GetIntSkill(targetModel.Save) - ap

	count := int(math.Min(float64(hits), float64(targetModels.Count())))

	models.Game().DiceRoller.Roll("Rolling for Saves", models.DiceRollType{
		Dice:   count,
		Target: save,
	}, func(success int, dice []int) {
		re.allocateAttacks(hits, success, count, targetModels, model)
		hits = hits - count
		re.onWoundsRolled(hits, targetModels, model)
	})
}

func (re ShootingAttackPhase) allocateAttacks(hits int, success int, count int, targetModels models.Stack, model *models.Model) {

	failed := count - success

	attackCount := 0

	if failed > 0 {
		attackCount = (failed / re.WeaponAttacks) + 1
	}

	engine.WriteMessage(fmt.Sprintf("Unit %s took %d wounds from %d attacks!", models.Game().SelectedTargetUnit.Name, failed, attackCount))

	allocatedHits := failed

	for attacks := 0; attacks < attackCount; attacks++ {

		if allocatedHits < 1 {
			break
		}

		popped, _ := targetModels.Pop()
		targetModel := popped.(*models.Model)

		if targetModel == nil {
			break
		}

		allocatedHits = re.allocateAttack(allocatedHits, model, targetModel)
	}

	if len(models.Game().SelectedTargetUnit.Models) <= 0 {
		engine.WriteMessage(fmt.Sprintf("Unit %s wiped out!", models.Game().SelectedTargetUnit.Name))
		models.Game().SelectedTargetUnit.Destroyed = true
		re.OnCompleted()

	} else {
		engine.WriteMessage(fmt.Sprintf("%s took %d casulties!", models.Game().SelectedTargetUnit.Name, len(models.Game().SelectedTargetUnit.DestroyedModels)))
	}
}

func (re ShootingAttackPhase) allocateAttack(hits int, model *models.Model, targetModel *models.Model) int {

	// we need to resolve per attack per model to allow extra wounds to be lost
	// for example an Assult 3 makes 5 wounds, we need to allocate 3 wounds to one model and 2 wounds to a second model
	for shots := 0; shots < re.WeaponAttacks; shots++ {

		if hits < 1 {
			break
		}

		hits = re.inflictWound(targetModel, model, hits)
	}

	return hits
}

func (re ShootingAttackPhase) inflictWound(target *models.Model, model *models.Model, hits int) int {

	dmg := re.weapon.Damage
	dead := false

	for hits > 0 && !dead {
		hits--
		engine.WriteMessage(fmt.Sprintf("inflicting %d damage wound!", dmg))
		dead, deadModel := models.Game().SelectedTargetUnit.InflictWounds(*target, dmg)

		if dead {
			engine.WriteMessage(fmt.Sprintf("%s was destroyed!", deadModel.Name))
			hits = 0
		}
	}

	return hits
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
