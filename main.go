package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)


func main() {
	loadRainSound()

	rainPlayer.SetVolume(0)
	rainPlayer.Play()
	loadHonkSound()

	loadBusSound()
	busPlayer.Play()
	
	game := NewGame()
	ebiten.SetWindowSize(800, 680)
	ebiten.SetWindowTitle("Turku")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
