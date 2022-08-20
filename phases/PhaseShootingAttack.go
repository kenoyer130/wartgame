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
	
	re.shootWeapon(re.weaponCount * shots, re.model)
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

	return &targetUnits
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

	if(hits == 0) {
		re.OnCompleted();
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
		re.allocateAttack(hits, success, count, targetModels, model)
		hits = hits - count
		re.onWoundsRolled(hits, targetModels, model)
	})
}

func (re ShootingAttackPhase) allocateAttack(hits int, success int, count int, targetModels models.Stack, model *models.Model) {

	failed := count - success

	engine.WriteMessage(fmt.Sprintf("%d models failed Saves!", failed))

	for i := 0; i < failed; i++ {

		if hits < 1 {
			break
		}

		popped, _ := targetModels.Pop()
		targetModel := popped.(*models.Model)

		hits = re.inflictWound(targetModel, model, hits)
	}

	if len(models.Game().SelectedTargetUnit.Models) <= 0 {
		engine.WriteMessage(fmt.Sprintf("Unit %s wiped out!", models.Game().SelectedTargetUnit.Name))
		models.Game().SelectedTargetUnit.Destroyed = true
		re.OnCompleted();

	} else {
		engine.WriteMessage(fmt.Sprintf("%s took %d casulties!", models.Game().SelectedTargetUnit.Name, len(models.Game().SelectedTargetUnit.DestroyedModels)))			
	}
}

func (re ShootingAttackPhase) inflictWound(target *models.Model, model *models.Model, hits int) int {

	dmg := re.weapon.Damage

	engine.WriteMessage(fmt.Sprintf("Model Saved Failed! %d wounds infliced!", dmg))

	dead := false

	for hits > 0 && !dead {
		hits--
		dead = models.Game().SelectedTargetUnit.InflictWounds(*target, dmg)
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