package models

import "github.com/google/uuid"

type GameStateUpdater struct {
}

func (re GameStateUpdater) Init() {

	for i := 0; i < len(Game().Players); i++ {
		for j := 0; j < len(Game().Players[i].Army.Units); j++ {
			if Game().Players[i].Army.Units[j].ID == "" {
				Game().Players[i].Army.Units[j].ID = uuid.New().String()
			}

			Game().Players[i].Army.Units[j].PlayerIndex = i

			for x := 0; x < len(Game().Players[i].Army.Units[j].Models); x++ {
				if Game().Players[i].Army.Units[j].Models[x].ID == "" {
					Game().Players[i].Army.Units[j].Models[x].ID = uuid.New().String()
				}

				Game().Players[i].Army.Units[j].Models[x].PlayerIndex = i
			}
		}
	}
}

func (re GameStateUpdater) UpdateUnit(player int, unit *Unit) {
	for i, thisUnit := range Game().Players[player].Army.Units {
		if thisUnit.ID == unit.ID {
			Game().Players[player].Army.Units[i] = unit
			return
		}
	}
}

func (re GameStateUpdater) UpdateModel(player int, model *Model) {
	for i, thisUnit := range Game().Players[player].Army.Units {
		for j, thisModel := range thisUnit.Models {
			if thisModel.ID == model.ID {
				Game().Players[player].Army.Units[i].Models[j] = model
			}
		}
	}
}
