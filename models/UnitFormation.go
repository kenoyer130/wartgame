package models

type UnitFormation int64

const (
	StandardUnitFormation UnitFormation = 0
)

func SetUnitFormation(UnitFormation UnitFormation, Unit *Unit, battleGround *BattleGround) {
	switch UnitFormation {
	case StandardUnitFormation:
		setStandardFormation(Unit, battleGround)
	}
}

func setStandardFormation(Unit *Unit, battleGround *BattleGround) {
	// standard formation is any leaders in front in a square format of 3 x ? until all models are placed
	// there is one grid space between each model in this formation
	leader, exists := getUnitLeader(Unit)

	if exists && IsBattleGroundLocationFree(leader.Location, battleGround) {
		leader.Location = Location{X: Unit.Location.X, Y: Unit.Location.Y + 2}
		PlaceBattleGroundEntity(leader, battleGround)
	}

	// loop through all remaining Models in a 3 x ? pattern until all Models placed

	rank := Unit.Location.X

	ModelX := rank
	ModelY := Unit.Location.Y

	for _, Model := range Unit.Models {

		// already placed
		if Model.ModelType == LeaderModelType {
			continue
		}

		placed := false

		for !placed {
			testLocation := Location{X: ModelX, Y: ModelY}

			// TODO: need to handle infinite loop if unable to place

			if IsBattleGroundLocationFree(testLocation, battleGround) {
				Model.Location = testLocation
				PlaceBattleGroundEntity(Model, battleGround)
				placed = true
			} else {

				ModelY = ModelY + 2

				if ModelY > Unit.Location.Y+4 {
					rank = rank + 2
					ModelX = rank
					ModelY = Unit.Location.Y
				}
			}
		}
	}
}

func getUnitLeader(Unit *Unit) (*Model, bool) {
	for _, Model := range Unit.Models {
		if Model.ModelType == LeaderModelType {
			return &Model, true
		}
	}

	return &Model{}, false
}
