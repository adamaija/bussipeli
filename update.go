package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	elapsed := time.Since(g.startTime)

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


	return nil
}