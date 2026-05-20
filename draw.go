package main

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func drawImage(screen *ebiten.Image, img *ebiten.Image, scaleX, scaleY, tx, ty float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(tx, ty)
	screen.DrawImage(img, op)
}

func (g *Game) currentImage(day, night *ebiten.Image) *ebiten.Image {
	if g.nightMode {
		return night
	}
	return day
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawSky(screen)
	g.drawFarBackground(screen)
	g.drawRiver(screen)
	g.drawForest(screen)
	g.drawRoad(screen)
	g.drawBusStops(screen)
	g.drawRatikka(screen)
	g.drawBus(screen)
	g.drawUI(screen)
}

func (g *Game) drawSky(screen *ebiten.Image) {
	var sky color.RGBA

	if g.nightMode {
		sky = color.RGBA{50, 50, 80, 255}
	} else {
		sky = color.RGBA{135, 206, 250, 255}
	}

	topHeight := (2 * screenHeight) / 3

	screen.SubImage(
		image.Rect(0, 0, screenWidth, topHeight),
	).(*ebiten.Image).Fill(sky)
}

func (g *Game) drawFarBackground(screen *ebiten.Image) {
	X := g.camX

	// Parallax scrolling for distant background objects
	drawImage(screen, g.currentImage(g.mk, g.mkN), 1, 1, -X*0.2-150, 18)
	drawImage(screen, g.currentImage(g.mk, g.mkN), 1, 1, -X*0.2+160, 8)

	drawImage(screen, g.currentImage(g.posankka, g.posN), 1, 1, -X*0.2+300, 53)

	drawImage(screen, g.currentImage(g.mk, g.mkN), 1, 1, -X*0.2+200, 18)

	drawImage(screen, g.currentImage(g.background, g.backgroundN), 1, 1, -X*0.2+100, 18)

	// Night lights
	if g.nightMode {
		valoW := float64(g.valo.Bounds().Dx()) * 0.5
		margin := int(valoW * 20)

		for x := int(-math.Mod(X*0.2, valoW)) - int(valoW); x < screenWidth+margin; x += int(valoW) {
			drawImage(screen, g.valo, 0.5, 0.5, float64(x), 120)
		}
	}
}

func (g *Game) drawRiver(screen *ebiten.Image) {
	var river color.RGBA

	if g.nightMode {
		river = color.RGBA{60, 90, 120, 255}
	} else {
		river = color.RGBA{89, 121, 153, 255}
	}

	screen.SubImage(
		image.Rect(0, 126, screenWidth, screenHeight),
	).(*ebiten.Image).Fill(river)
}

func (g *Game) drawForest(screen *ebiten.Image) {
	X := g.camX

	// Infinite forest tiling
	metsaImg := g.currentImage(g.metsa, g.metsaN)
	metsaW := metsaImg.Bounds().Dx()

	start := int(-X) / metsaW * metsaW

	for x := start; x < screenWidth-int(-X)+metsaW; x += metsaW {
		drawImage(screen, metsaImg, 1, 1, float64(x+int(-X)), 95)
	}

	drawImage(screen, g.currentImage(g.Tietyokyltti, g.TietyokylttiN), 1, 1, -X+1300, 100)
	drawImage(screen, g.currentImage(g.Tulossa2, g.Tulossa2N), 1, 1, -X+1380, 97)
}

func (g *Game) drawRoad(screen *ebiten.Image) {
	road := color.RGBA{50, 50, 50, 255}
	road2 := color.RGBA{35, 35, 35, 255}

	screen.SubImage(
		image.Rect(0, 150, screenWidth, screenHeight),
	).(*ebiten.Image).Fill(road2)

	screen.SubImage(
		image.Rect(0, 155, screenWidth, screenHeight),
	).(*ebiten.Image).Fill(road)
}

func (g *Game) drawBusStops(screen *ebiten.Image) {
	X := g.camX

	// Bus stop 1
	drawImage(screen, g.currentImage(g.p1, g.p1N), 2, 2, -X+130, 64)

	// Street lamps
	lampImg := g.currentImage(g.l1, g.l1N)

	for x := int(-math.Mod(X, 200) - 200); x < screenWidth; x += 200 {
		drawImage(screen, lampImg, 2, 2, float64(x), 64)
	}

	// Bus stop 2
	drawImage(screen, g.currentImage(g.p2, g.p2N), 2, 2, -X+740, 64)
}

func (g *Game) drawRatikka(screen *ebiten.Image) {
	X := g.camX

	drawImage(screen, g.currentImage(g.cone, g.coneN), 2, 2, -X+1380, 75)
	drawImage(screen, g.currentImage(g.cone, g.coneN), 2, 2, -X+1395, 80)
	drawImage(screen, g.currentImage(g.cone, g.coneN), 2, 2, -X+1220, 73)
	drawImage(screen, g.currentImage(g.holes, g.holes), 2, 2, -X+1360, 81)

}

func (g *Game) drawBus(screen *ebiten.Image) {
	sx := g.width * g.index

	sheet := g.wr
	if g.facingLeft {
		sheet = g.wl
	}

	frame := sheet.SubImage(
		image.Rect(sx, 0, sx+g.width, g.height),
	).(*ebiten.Image)

	busScreenX := float64(screenWidth)/2 - float64(g.width)*0.25
	drawImage(screen, frame, 0.5, 0.5, float64(busScreenX), g.busY)

	// Bus headlights at night
	if g.nightMode {
		if g.facingLeft {
			drawImage(screen, g.valo, 2, 2, busScreenX+9, g.busY+54)
			drawImage(screen, g.valo, 2, 2, busScreenX+24, g.busY+54)
		} else {
			drawImage(screen, g.valo, 2, 2, busScreenX+70, g.busY+54)
			drawImage(screen, g.valo, 2, 2, busScreenX+85, g.busY+54)
		}
	}
	X := g.camX
	drawImage(screen, g.currentImage(g.cone, g.coneN), 2, 2, -X+1270, 86)
}

func (g *Game) drawUI(screen *ebiten.Image) {
	if !g.showMessage {
		return
	}

	op := &ebiten.DrawImageOptions{}

	// Apply fade
	op.ColorScale.ScaleAlpha(float32(g.messageAlpha))

	op.GeoM.Translate(1, 5)

	screen.DrawImage(g.mode, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
