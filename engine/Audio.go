package engine

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const SampleRate = 48000
var audioContext = audio.NewContext(SampleRate)

func PlaySound(id string) {

	f, err := os.Open("assets/audio/" + id + ".wav")
	if err != nil {
		log.Fatal(err)
	}

	d, err := wav.DecodeWithSampleRate(SampleRate, f)
	if err != nil {
		log.Fatal(err)
	}

	player, err := audioContext.NewPlayer(d)
	if err != nil {
		log.Fatal(err)
	}

	player.Play()
}
