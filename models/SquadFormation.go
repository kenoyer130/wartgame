package models

type SquadFormation int64

const (
	StandardSquadFormation SquadFormation = 0
)

func SetSquadFormation(squadFormation SquadFormation, squad *Squad, battleGround *BattleGround) {
	switch squadFormation {
	case StandardSquadFormation:
		setStandardFormation(squad, battleGround)
	}
}

func setStandardFormation(squad *Squad, battleGround *BattleGround) {
	// standard formation is any leaders in front in a square format of 3 x ? until all models are placed
	// there is one grid space between each model in this formation
	leader, exists := getSquadLeader(squad)

	if exists && IsBattleGroundLocationFree(leader.Location, battleGround) {
		leader.Location = Location{X: squad.Location.X, Y: squad.Location.Y + 2}
		PlaceBattleGroundEntity(leader, battleGround)
	}

	// loop through all remaining units in a 3 x ? pattern until all units placed

	rank := squad.Location.X

	unitX := rank
	unitY := squad.Location.Y

	for _, unit := range squad.Units {

		// already placed
		if unit.UnitType == LeaderUnitType {
			continue
		}

		placed := false

		for !placed {
			testLocation := Location{X: unitX, Y: unitY}

			// TODO: need to handle infinite loop if unable to place

			if IsBattleGroundLocationFree(testLocation, battleGround) {
				unit.Location = testLocation
				PlaceBattleGroundEntity(unit, battleGround)
				placed = true
			} else {

				unitY = unitY + 2

				if unitY > squad.Location.Y+4 {
					rank = rank + 2
					unitX = rank
					unitY = squad.Location.Y
				}
			}
		}
	}
}

func getSquadLeader(squad *Squad) (*Unit, bool) {
	for _, unit := range squad.Units {
		if unit.UnitType == LeaderUnitType {
			return &unit, true
		}
	}

	return &Unit{}, false
}
