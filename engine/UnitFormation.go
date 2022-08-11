package engine

import (
	"math"

	"github.com/kenoyer130/wartgame/models"
	"github.com/kenoyer130/wartgame/ui"
)

type UnitFormation int64

const (
	StandardUnitFormation UnitFormation = 0
)

func SetUnitFormation(UnitFormation UnitFormation, Unit *models.Unit, battleGround *models.BattleGround) *models.Unit {
	switch UnitFormation {
	case StandardUnitFormation:
		setStandardFormation(Unit, battleGround)
	}

	return Unit
}

func setStandardFormation(Unit *models.Unit, battleGround *models.BattleGround) {
	// standard formation is any leaders in front in a square format of 3 x ? until all models are placed
	// there is one grid space between each model in this formation
	leader, exists := getUnitLeader(Unit)

	if exists && models.IsBattleGroundLocationFree(leader.Location, battleGround) {
		leader.Location = models.Location{X: Unit.Location.X, Y: Unit.Location.Y + 2}
		models.PlaceBattleGroundEntity(leader, battleGround)
	}

	// loop through all remaining Models in a 3 x ? pattern until all Models placed

	rank := Unit.Location.X

	row := 0
	col := 0

	ModelX := rank
	ModelY := Unit.Location.Y

	height := 0
	width := 0

	for _, Model := range Unit.Models {

		// already placed
		if Model.ModelType == models.LeaderModelType {
			continue
		}

		placed := false

		for !placed {
			testLocation := models.Location{X: ModelX, Y: ModelY}

			// TODO: need to handle infinite loop if unable to place

			if models.IsBattleGroundLocationFree(testLocation, battleGround) {

				placeModel:= *Model

				placeModel.Location = testLocation
				models.PlaceBattleGroundEntity(&placeModel, battleGround)
				placed = true

				models.Game().GameStateUpdater.UpdateModel(models.Game().CurrentPlayerIndex, &placeModel)

				width = int(math.Max(float64(width), float64(row)))
				height = int(math.Max(float64(height), float64(col)))

			} else {

				ModelY = ModelY + 2	
				
				col = col + 2

				if ModelY > Unit.Location.Y+4 {
					rank = rank + 2

					row = row + 2

					col = 0

					ModelX = rank
					ModelY = Unit.Location.Y
				}
			}
		}
	}

	Unit.Rect = ui.Rect{X: Unit.Location.X, Y: Unit.Location.Y, W: width, H: height}
}

func getUnitLeader(Unit *models.Unit) (*models.Model, bool) {
	for _, Model := range Unit.Models {
		if Model.ModelType == models.LeaderModelType {
			return Model, true
		}
	}

	return &models.Model{}, false
}
