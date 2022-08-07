package phases

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func StartPhaseShootingAttack() {

	// figure out how many weapons are making the attack
	weaponCount := getFiringWeaponCount()

	weapon := models.Game().Assets.Weapons[models.Game().SelectedWeaponName]

	for i := 0; i < weaponCount; i++ {
		thisWeapon := weapon
		models.Game().SelectedWeapons = append(models.Game().SelectedWeapons, thisWeapon)
	}

	model := getModel(weapon)

	if model.Name == "" {
		log.Fatal("no model found for " + weapon.Name)
	}

	models.Game().StatusMessage.Messsage = fmt.Sprintf("%s is shooting %s with %d %s ", models.Game().SelectedPhaseUnit.Name, models.Game().SelectedTargetUnit.Name, weaponCount, weapon.Name)

	engine.RollDice("Rolling for Attack",
		engine.DiceRollType{
			Dice:   weaponCount,
			Target: model.GetBallisticSkill(),
		}, func(success int, dice []int) {
			onAttackRolled(success, &model)
		})
}

func onAttackRolled(success int, model *models.Model) {
	target := getWoundTarget(model.Strength, models.Game().SelectedTargetUnit.Models[0].Toughness)

	engine.WriteMessage("Wound target is " + strconv.Itoa(target))

	engine.RollDice("Rolling for Wounds", engine.DiceRollType{
		Dice:   success,
		Target: target,
	}, func(success int, dice []int) {
		onWoundRolled(success, model)
	})
}

func onWoundRolled(hits int, model *models.Model) {
	target := 0
	allocateAttacks(target, hits, model)
}

func allocateAttacks(target int, hits int, model *models.Model) {

	ap := models.Game().SelectedWeapons[models.Game().SelectedWeaponIndex].ArmorPiercing

	save := (models.Game().SelectedTargetUnit.Models[target].GetIntSkill(models.Game().SelectedTargetUnit.Models[target].Save)) - ap

	engine.RollDice("Rolling for Save", engine.DiceRollType{
		Dice:   1,
		Target: save,	
	}, func(success int, dice []int) {
		allocateAttack(hits, success, target, model)
	})
}

func allocateAttack(hits int, success int, target int, model *models.Model) {
	hits--

	if success == 1 {
		engine.WriteMessage("Model made Save!")
		nextWound(hits, model)
	} else {
		inflictWound(target, model, hits)
	}
}

func inflictWound(target int, model *models.Model, hits int) {

	dmg := models.Game().SelectedWeapons[models.Game().SelectedWeaponIndex].Damage

	models.Game().SelectedTargetUnit.InflictWounds(target, dmg)

	engine.WriteMessage(fmt.Sprintf("Model Saved Failed! %d wounds infliced!", dmg))

	if len(models.Game().SelectedTargetUnit.Models) <= 0 {
		engine.WriteMessage(fmt.Sprintf("Unit %s wiped out!", models.Game().SelectedTargetUnit.Name))
		models.Game().SelectedTargetUnit.Destroyed = true

	} else {
		nextWound(hits, model)
	}
}

func nextWound(hits int, model *models.Model) {
	if hits > 0 {
		onWoundRolled(hits, model)
	} else {

		engine.WriteMessage(fmt.Sprintf("%s took %d casulties!", models.Game().SelectedTargetUnit.Name, len(models.Game().SelectedTargetUnit.DestroyedModels)))

		StartPhaseShootingWeapons()
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

func getModel(weapon models.Weapon) models.Model {
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

func getFiringWeaponCount() int {
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
