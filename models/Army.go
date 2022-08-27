package models

type Army struct {
	ID             string
	Name           string
	Units          []*Unit
	DestroyedUnits []*Unit
}

func (re *Army) RemoveDestroyedUnits() {

	for i := 0; i < len(re.Units); {
		if re.Units[i] != nil {
			i++
			continue
		}

		if i < len(re.Units)-1 {
			copy(re.Units[i:], re.Units[i+1:])
		}

		re.Units[len(re.Units)-1] = nil
		re.Units = re.Units[:len(re.Units)-1]
	}

	for i, unit := range re.Units {
		if len(unit.Models) == 0 {
			unit.Destroyed = true
			re.DestroyedUnits = append(re.DestroyedUnits, unit)
			re.Units[i] = nil
		}
	}
}
