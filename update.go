package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	elapsed := time.Since(g.startTime)
	g.rainIntensity = 1

    dt := 1.0 / 60.0

    // Spawn rain
    if g.rainEnabled {
		spawnCount := int(g.rainIntensity * 10)

		for i := 0; i < spawnCount; i++ {
			g.SpawnRain()
		}
	}

    g.UpdateRain(dt)
    g.UpdateSplashes(dt)
	targetVolume := 0.0

	if g.rainEnabled {
		targetVolume = 0.3
	}

	// Smooth fade speed
	fadeSpeed := 0.005

	if g.rainVolume < targetVolume {
		g.rainVolume += fadeSpeed

		if g.rainVolume > targetVolume {
			g.rainVolume = targetVolume
		}
	}

	if g.rainVolume > targetVolume {
		g.rainVolume -= fadeSpeed

		if g.rainVolume < targetVolume {
			g.rainVolume = targetVolume
		}
	}

	rainPlayer.SetVolume(g.rainVolume)

	if g.showMessage {
		if elapsed > 10*time.Second {
			// Fade out over 2 seconds
			fadeDuration := 2 * time.Second
			fadeProgress := (elapsed - 10*time.Second).Seconds() / fadeDuration.Seconds()

			if fadeProgress >= 1 {
				g.showMessage = false
				g.messageAlpha = 0
			} else {
				g.messageAlpha = 1.0 - fadeProgress
			}
		} else {
			g.messageAlpha = 1.0
		}
	}
	g.moving = false

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.camX += 2
		g.facingLeft = false
		g.moving = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.camX -= 2
		g.facingLeft = true
		g.moving = true
	}

	if g.camX < 0 {
		g.camX = 0
	}

	if g.camX > maxCamX {
		g.camX = maxCamX
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !g.spacePressed {
			g.nightMode = !g.nightMode // Change mode
			g.spacePressed = true      // Pressed
		}
	} else {
		g.spacePressed = false // vapautus
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		rainPlayer.Rewind()
		rainPlayer.Play()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		honkPlayer.Rewind()
		honkPlayer.Play()
	}

	if g.nightMode {

		// Wait before another lightning flash
		if g.lightningCoolDown > 0 {
			g.lightningCoolDown--
		}

		// Lightning currently active
		if g.lightningFrames > 0 {
			g.lightningFrames--

			// Trigger flash slightly later
			if g.lightningFrames == 1 {
				g.lightFlashFrames = 4
			}
		}

		if g.lightFlashFrames > 0 {
			g.lightFlashFrames--
		}

		// Randomly trigger lightning
		if g.lightningCoolDown <= 0 && g.lightningFrames <= 0 {

			// Small random chance each frame
			if rand.Intn(500) == 0 {

				g.lightningFrames = 6      // lightning image first
				g.lightFlashFrames = 10     // flash delayed
				g.lightningCoolDown = 180 + rand.Intn(300)
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		if !g.rainKeyPressed {
			g.rainEnabled = !g.rainEnabled
			g.rainKeyPressed = true
		}
	} else {
		g.rainKeyPressed = false
	}

	g.tick++
	if g.tick%8 == 0 {
		g.index = (g.index + 1) % 4
	}

	if g.tick%30 == 0 {
		fmt.Printf("camX: %.2f | busX: %.2f | busY: %.2f\n",
			g.camX,
			g.busX,
			g.busY,
		)
	}

	if g.tick%100 == 0 {
		
	}


	return nil
}