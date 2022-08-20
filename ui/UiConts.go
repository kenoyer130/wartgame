package ui

import "image/color"

const Margin = 10
const TileSize = 64

func GetTextColor() color.Color {
	return color.RGBA{255, 255, 255, 255}
}

func GetGridOutlineColor() color.Color {
	return color.RGBA{35, 31, 33, 255}
}

func GetBattleGroundBackgroundColor() color.Color {
	return color.RGBA{166, 142, 154, 1}
}

func GetWoundColor() color.Color {
	return color.RGBA{222, 0, 0, 255}
}

func GetMoveRangeColor() color.Color {
	return color.RGBA{231, 180, 206, 255}
}
