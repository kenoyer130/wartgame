package models

func InRange(id string, r int, x int, y int) []Entity {

	entities := []Entity{}

	for i := x - r; i < x+r; i++ {
		for j := y - r; j < y+r; j++ {

			entity := Game().BattleGround.GetEntityAtLocation(Location{X: i, Y: j})

			if entity != nil && entity.GetID() != id {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}
