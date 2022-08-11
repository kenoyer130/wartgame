package phases

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/weaponabilities"
)

type ShootingAttackPhase struct {
	ShootingWeaponPhase *ShootingWeaponPhase
	onCompleted         func()
	TargetUnits         models.Stack
}

func (re ShootingAttackPhase) GetName() (models.GamePhase, models.PhaseStep) {
	return models.ShootingPhase, models.Nil
}

func (re ShootingAttackPhase) Start() {

	// figure out how many weapons are making the attack
	weaponCount := re.getFiringWeaponCount()

	// if throwing gernade we only throw one
	if models.Game().Assets.Weapons[models.Game().SelectedWeaponName].WeaponType.Type == "Gre" {
		weaponCount = 1
	}

	weapon := models.Game().Assets.Weapons[models.Game().SelectedWeaponName]

	models.Game().SelectedWeapons = []models.Weapon{}

	for i := 0; i < weaponCount; i++ {
		thisWeapon := weapon
		models.Game().SelectedWeapons = append(models.Game().SelectedWeapons, thisWeapon)
	}

	model := re.getModel(weapon)
	re.TargetUnits = *re.setModelsByToughness()

	if model.Name == "" {
		log.Fatal("no model found for " + weapon.Name)
	}

	models.Game().StatusMessage.Messsage = fmt.Sprintf("%s is shooting %s with %d %s ", models.Game().SelectedPhaseUnit.Name, models.Game().SelectedTargetUnit.Name, weaponCount, weapon.Name)

	shots := weaponabilities.ApplyWeaponAbilityShot(weapon)
	index := 0

	re.onCompleted = func() {
		index++

		if index >= shots {
			re.ShootingWeaponPhase.Start()
		} else {
			re.shootWeapon(weaponCount, model)
		}
	}

	re.shootWeapon(weaponCount, model)
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
			group.Push(&checkModel)

		} else if toughness != toughCheck {
			targetUnits.Push(group)
			group = models.Stack{}
			group.Push(&checkModel)

		} else {
			group.Push(&checkModel)
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

	ap := models.Game().SelectedWeapons[models.Game().SelectedWeaponIndex].ArmorPiercing

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
		re.onCompleted()

	} else {
		engine.WriteMessage(fmt.Sprintf("%s took %d casulties!", models.Game().SelectedTargetUnit.Name, len(models.Game().SelectedTargetUnit.DestroyedModels)))
		re.ShootingWeaponPhase.Start()
	}
}

func (re ShootingAttackPhase) inflictWound(target *models.Model, model *models.Model, hits int) int {

	dmg := models.Game().SelectedWeapons[models.Game().SelectedWeaponIndex].Damage

	engine.WriteMessage(fmt.Sprintf("Model Saved Failed! %d wounds infliced!", dmg))

	dead := false
	
	for( hits > 0 && !dead) {
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

func (re ShootingAttackPhase) getModel(weapon models.Weapon) models.Model {
	model := models.Model{}

	for _, m := range models.Game().SelectedPhaseUnit.Models {
		for _, modelWeapon := range m.Weapons {
			if modelWeapon.Name == weapon.Name {
				model = *m
			}
		}
	}

	return model
}

func (re ShootingAttackPhase) getFiringWeaponCount() int {
	weaponCount := 0

	for _, model := range models.Game().SelectedPhaseUnit.Models {
		for _, weapon := range model.Weapons {
			if weapon.Name == models.Game().SelectedWeaponName {
				weaponCount++
			}
		}
	}

	return weaponCount
}
