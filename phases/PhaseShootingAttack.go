package phases

import (
	"log"
	"strconv"

	"github.com/kenoyer130/wartgame/engine"
	"github.com/kenoyer130/wartgame/models"
)

func StartPhaseShootingAttack(g *engine.Game) {

	// figure out how many weapons are making the attack
	weaponCount := getFiringWeaponCount(g)

	weapon := g.Assets.Weapons[g.SelectedWeapon]

	model := getModel(g, weapon)

	if model.Name == "" {
		log.Fatal("no model found for " + weapon.Name)
	}

	engine.RollDice("Rolling for Attack", weaponCount, model.GetBallisticSkill(), func(success int, dice []int) {
		onAttackRolled(success, &model, g)
	}, g)
}

func onAttackRolled(success int, model *models.Model, g *engine.Game) {
	target := getWoundTarget(model.Strength, g.SelectedTargetUnit.Models[0].Toughness)

	engine.WriteMessage("Wound target is " + strconv.Itoa(target))

	engine.RollDice("Rolling for Wounds", success, target, func(success int, dice []int) {
		onWoundRolled(success, model, g)
	}, g)
}

func onWoundRolled(success int, model *models.Model, g *engine.Game) {
	index :=0
	target := 0

	save := (g.SelectedTargetUnit.Models[target].GetIntSkill(g.SelectedTargetUnit.Models[target].Save) - model.SpecificWeapon.ArmorPiercing)

			engine.RollDice("Rolling for Save", 1, save, func(success int, dice []int) {
				if(success == 0) {
					engine.WriteMessage("Unit Saved!")
					
				} else {

				}
			}, g)
		}
	
}

func getWoundTarget(str int, toughness int) int {
	if str*2 > toughness {
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

	if (str / 2) < toughness {
		return 6
	}

	return 4
}

func getModel(g *engine.Game, weapon models.Weapon) models.Model {
	model := models.Model{}

	for _, m := range g.SelectedPhaseUnit.Models {
		for _, modelWeapon := range m.Weapons {
			if modelWeapon == weapon.Name {
				model = m
			}
		}
	}

	return model
}

func getFiringWeaponCount(g *engine.Game) int {
	weaponCount := 0

	for _, model := range g.SelectedPhaseUnit.Models {
		for _, weapon := range model.Weapons {
			if weapon == g.SelectedWeapon {
				weaponCount++
			}
		}
	}

	return weaponCount
}
