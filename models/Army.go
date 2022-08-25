package models

type Army struct {
	ID             string
	Name           string
	Units          []*Unit
	DestroyedUnits []*Unit
}

func (re *Army) RemoveDestroyedUnits() {
	for i, unit := range re.Units {
		if len(unit.Models) == 0 {
			unit.Destroyed = true
			re.DestroyedUnits = append(re.DestroyedUnits, unit)
			re.Units[i] = nil
		}
	}
}
