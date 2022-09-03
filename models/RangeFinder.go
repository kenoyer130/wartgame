package models

import "github.com/kenoyer130/wartgame/interfaces"

func InRange(id string, playerIndex int, r int, x int, y int) []interfaces.Entity {

	entities := []interfaces.Entity{}

	for i := x - r; i < x+r; i++ {
		for j := y - r; j < y+r; j++ {

			entity := Game().BattleGround.GetEntityAtLocation(interfaces.Location{X: i, Y: j})

			if entity != nil && entity.GetID() != id && entity.GetPlayerIndex() != playerIndex {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}
