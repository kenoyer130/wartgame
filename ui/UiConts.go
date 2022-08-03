package ui

import "image/color"

const Margin = 10
const TileSize = 32

func GetTextColor() color.Color {
	return color.RGBA{255, 255, 255, 255}
}

func GetGridOutlineColor() color.Color {
	return color.RGBA{35, 31, 33, 255}
}

func GetBattleGroundBackgroundColor() color.Color {
	return color.RGBA{166, 142, 154, 1}
}
