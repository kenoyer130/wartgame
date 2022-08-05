package engine

import "github.com/hajimehoshi/ebiten/v2"

var KeyBoardRegistry = make(map[ebiten.Key]func())