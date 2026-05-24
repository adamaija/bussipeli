package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

const sampleRate = 44100

var audioContext = audio.NewContext(sampleRate)

var rainPlayer *audio.Player
var honkPlayer *audio.Player
var busPlayer  *audio.Player

func loadRainSound() {
	f, err := os.Open("assets/sounds/rainsound.mp3")
	if err != nil {
		log.Fatal(err)
	}

	stream, err := mp3.DecodeWithSampleRate(sampleRate, f)
	if err != nil {
		log.Fatal(err)
	}

	loop := audio.NewInfiniteLoop(stream, stream.Length())

	rainPlayer, err = audioContext.NewPlayer(loop)
	if err != nil {
		log.Fatal(err)
	}

	rainPlayer.SetVolume(0.3)


}


func loadHonkSound() {
	f, err := os.Open("assets/sounds/honk.mp3")
	if err != nil {
		log.Fatal(err)
	}

	stream, err := mp3.DecodeWithSampleRate(sampleRate, f)
	if err != nil {
		log.Fatal(err)
	}

	honkPlayer, err = audioContext.NewPlayer(stream)
	if err != nil {
		log.Fatal(err)
	}

	honkPlayer.SetVolume(0.7)
}

func loadBusSound() {
	f, err := os.Open("assets/sounds/bussound.mp3")
	if err != nil {
		log.Fatal(err)
	}

	stream, err := mp3.DecodeWithSampleRate(sampleRate, f)
	if err != nil {
		log.Fatal(err)
	}

	loop := audio.NewInfiniteLoop(stream, stream.Length())

	busPlayer, err = audioContext.NewPlayer(loop)
	if err != nil {
		log.Fatal(err)
	}

	busPlayer.SetVolume(0.1)
}