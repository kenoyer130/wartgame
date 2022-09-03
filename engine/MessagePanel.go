package engine

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kenoyer130/wartgame/models"
)

var messages []string

const messageCount = 10

func WriteMessage(msg string) {
	messages = append(messages, msg)
	log.Println(msg)
}

func WriteStatusMessage(msg string) {
	models.Game().StatusMessage.Messsage = msg
	WriteMessage(msg)
}

func WriteStatusKeys(msg string) {
	models.Game().StatusMessage.Keys = msg
	WriteMessage(msg)
}

func getMessagePanel() *ebiten.Image {

	panel := NewPanel(400, 800)

	panel.addTitle("Messages")

	var msgs []string

	if len(messages) > messageCount {
		msgs = messages[len(messages)-messageCount:]
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
