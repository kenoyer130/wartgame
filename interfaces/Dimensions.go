package interfaces

import "math"

type Location struct {
	X int
	Y int
}

func (re Location) Subtract(l Location) Location {
	return Location{
		X: int(math.Abs(float64(re.X) - float64(l.X))),
		Y: int(math.Abs(float64(re.Y) - float64(l.Y))),
	}
}

type Rectangle struct {
	X      int
	Y      int
	Width  int
	Height int
}