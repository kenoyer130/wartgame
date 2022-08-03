package ui

type Rect struct {
	X int
	Y int
	W int
	H int
}

func (re Rect) InPixelBounds(x int, y int) bool {
	x1 := re.X * TileSize
	y1 := re.Y * TileSize

	w := x1 + (re.W+1)*TileSize
	h := y1 + (re.H+1)*TileSize

	if x > x1 && x < w && y > y1 && y < h {
		return true
	} else {
		return false
	}
}
