package phases

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

type ShootingAttackPhase struct {
	ShootingWeaponPhase *ShootingWeaponPhase
}

func (re ShootingAttackPhase) GetName() (models.GamePhase, models.PhaseStep) {
	return models.ShootingPhase, models.Nil
}

func (re ShootingAttackPhase) Start() {

	// figure out how many weapons are making the attack
	weaponCount := re.getFiringWeaponCount()

	weapon := models.Game().Assets.Weapons[models.Game().SelectedWeaponName]

	models.Game().SelectedWeapons = []models.Weapon{}

	for i := 0; i < weaponCount; i++ {
		thisWeapon := weapon
		models.Game().SelectedWeapons = append(models.Game().SelectedWeapons, thisWeapon)
	}

	model := re.getModel(weapon)

	if model.Name == "" {
		log.Fatal("no model found for " + weapon.Name)
	}

	models.Game().StatusMessage.Messsage = fmt.Sprintf("%s is shooting %s with %d %s ", models.Game().SelectedPhaseUnit.Name, models.Game().SelectedTargetUnit.Name, weaponCount, weapon.Name)

	models.Game().DiceRoller.Roll("Rolling for Attack",
		models.DiceRollType{
			Dice:   weaponCount,
			Target: model.GetBallisticSkill(),
		}, func(success int, dice []int) {
			re.onAttackRolled(success, &model)
		})
}

func (re ShootingAttackPhase) onAttackRolled(success int, model *models.Model) {
	target := getWoundTarget(model.Strength, models.Game().SelectedTargetUnit.Models[0].Toughness)

	engine.WriteMessage("Wound target is " + strconv.Itoa(target))

	models.Game().DiceRoller.Roll("Rolling for Wounds", models.DiceRollType{
		Dice:   success,
		Target: target,
	}, func(success int, dice []int) {
		re.onWoundRolled(success, model)
	})
}

func (re ShootingAttackPhase) onWoundRolled(hits int, model *models.Model) {
	target := 0
	re.allocateAttacks(target, hits, model)
}

func (re ShootingAttackPhase) allocateAttacks(target int, hits int, model *models.Model) {

	ap := models.Game().SelectedWeapons[models.Game().SelectedWeaponIndex].ArmorPiercing

	save := (models.Game().SelectedTargetUnit.Models[target].GetIntSkill(models.Game().SelectedTargetUnit.Models[target].Save)) - ap

	models.Game().DiceRoller.Roll("Rolling for Save", models.DiceRollType{
		Dice:   1,
		Target: save,
	}, func(success int, dice []int) {
		re.allocateAttack(hits, success, target, model)
	})
}

func (re ShootingAttackPhase) allocateAttack(hits int, success int, target int, model *models.Model) {
	hits--

	if success == 1 {
		engine.WriteMessage("Model made Save!")
		re.nextWound(hits, model)
	} else {
		re.inflictWound(target, model, hits)
	}
}

func (re ShootingAttackPhase) inflictWound(target int, model *models.Model, hits int) {

	dmg := models.Game().SelectedWeapons[models.Game().SelectedWeaponIndex].Damage

	engine.WriteMessage(fmt.Sprintf("Model Saved Failed! %d wounds infliced!", dmg))

	models.Game().SelectedTargetUnit.InflictWounds(target, dmg)

	if len(models.Game().SelectedTargetUnit.Models) <= 0 {
		engine.WriteMessage(fmt.Sprintf("Unit %s wiped out!", models.Game().SelectedTargetUnit.Name))
		models.Game().SelectedTargetUnit.Destroyed = true
		re.ShootingWeaponPhase.Start()

	} else {
		re.nextWound(hits, model)
	}
}

func (re ShootingAttackPhase) nextWound(hits int, model *models.Model) {
	if hits > 0 {
		re.onWoundRolled(hits, model)
	} else {

		engine.WriteMessage(fmt.Sprintf("%s took %d casulties!", models.Game().SelectedTargetUnit.Name, len(models.Game().SelectedTargetUnit.DestroyedModels)))
		re.ShootingWeaponPhase.Start()
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

func (re ShootingAttackPhase) getModel(weapon models.Weapon) models.Model {
	model := models.Model{}

	for _, m := range models.Game().SelectedPhaseUnit.Models {
		for _, modelWeapon := range m.Weapons {
			if modelWeapon == weapon.Name {
				model = m
			}
		}
	}

	return model
}

func (re ShootingAttackPhase) getFiringWeaponCount() int {
	weaponCount := 0

	for _, model := range models.Game().SelectedPhaseUnit.Models {
		for _, weapon := range model.Weapons {
			if weapon == models.Game().SelectedWeaponName {
				weaponCount++
			}
		}
	}

	return weaponCount
}
