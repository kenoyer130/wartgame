package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawLine(length int) *ebiten.Image {

	line := ebiten.NewImage(length, length)
	color := GetGridOutlineColor()	
	line.Fill(color)
	return line
}