package engine

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var messages []string

func WriteMessage(msg string) {
	messages = append(messages, msg)
	log.Println(msg)
}

func getMessagePanel() *ebiten.Image {

	panel := NewPanel(400, 400)

	panel.addTitle("Messages")

	var msgs []string

	if len(messages) > 4 {
		msgs = messages[len(messages)-4:]
	} else {
		msgs = messages
	}

	r := 2

	for _, msg := range msgs {
		panel.addMessage(msg, r)
		r++
	}

	return panel.Img
}
