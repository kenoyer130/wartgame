package models

func InRange(id string, playerIndex int, r int, x int, y int) []Entity {

	entities := []Entity{}

	for i := x - r; i < x+r; i++ {
		for j := y - r; j < y+r; j++ {

			entity := Game().BattleGround.GetEntityAtLocation(Location{X: i, Y: j})

			if entity != nil && entity.GetID() != id && entity.GetPlayerIndex() != playerIndex {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}
