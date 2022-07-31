package models

type SquadFormation int64

const (
	Standard SquadFormation = 0
)

func SetSquadFormation(squadFormation SquadFormation, squad *Squad, battleGround *BattleGround) {
	switch squadFormation {
	case Standard:
		setStandardFormation(squad, battleGround)
	}
}

func setStandardFormation(squad *Squad, battleGround *BattleGround) {
	// standard formation is any leaders in front in a square format
	leader, exists := getSquadLeader(squad)

	if exists && IsBattleGroundLocationFree(leader.Location, battleGround) {
		leader.Location = Location{X: squad.Location.X - 2, Y: squad.Location.Y}
		PlaceBattleGroundEntity(leader, battleGround)
	}
}

func getSquadLeader(squad *Squad) (Unit, bool) {
	for _, unit := range squad.Units {
		if unit.UnitType == Leader {
			return unit, true
		}
	}

	return Unit{}, false
}
