package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var messages []string

func WriteMessage(msg string) {
	messages = append(messages, msg)
}

func getMessagePanel(g *Game) *ebiten.Image {

	panel := NewPanel(400, 400)

	panel.addTitle("Messages")

	var msgs []string

	if len(messages) > 4 {
		msgs = messages[len(messages)-4:]
	} else {
		msgs = messages
	}
	
	for _, msg := range msgs {
		panel.addMessage(msg)
	}

	return panel.Img
}

