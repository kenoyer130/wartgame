package engine

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

const SampleRate = 48000

var audioContext = audio.NewContext(SampleRate)

func PlaySound(id string) {

	path := "assets/audio/" + id + "."

	f, err := os.Open(path)
	if err != nil {
		log.Println("warning ", id, " ", path, " file missing")
		return
	}

	d, err := mp3.DecodeWithSampleRate(SampleRate, f)
	if err != nil {
		log.Fatal(err)
	}

	player, err := audioContext.NewPlayer(d)
	if err != nil {
		log.Fatal(err)
	}

	player.Play()
}
