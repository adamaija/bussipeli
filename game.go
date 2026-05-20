package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	startTime   time.Time
	showMessage bool
	wr          *ebiten.Image // right
	wl          *ebiten.Image // left
	background  *ebiten.Image
	metsa       *ebiten.Image
	p1          *ebiten.Image
	p2          *ebiten.Image
	puska       *ebiten.Image
	posankka    *ebiten.Image
	mk          *ebiten.Image
	l1          *ebiten.Image
	p22         *ebiten.Image
	valo        *ebiten.Image
	mode        *ebiten.Image

	backgroundN *ebiten.Image
	metsaN      *ebiten.Image
	p1N         *ebiten.Image
	p2N         *ebiten.Image
	posN        *ebiten.Image
	mkN         *ebiten.Image
	l1N         *ebiten.Image

	Tulossa1L  *ebiten.Image
	Tulossa1LN *ebiten.Image

	Tulossa2  *ebiten.Image
	Tulossa2N *ebiten.Image

	Tietyokyltti  *ebiten.Image
	TietyokylttiN *ebiten.Image

	cone  *ebiten.Image
	coneN *ebiten.Image

	holes  *ebiten.Image

	width      int
	height     int
	index      int
	facingLeft bool
	moving     bool
	tick       int

	camX	 float64
	camY 	float64
	busX 	float64
	busY 	float64

	nightMode    bool
	spacePressed bool

	messageAlpha float64 // 1.0 = fully visible, 0.0 = invisible
}

func NewGame() *Game {
	return &Game{
		startTime:   time.Now(),
		showMessage: true,
		wr:          loadImage("assets/Busright.png"),
		wl:          loadImage("assets/Busleft.png"),
		background:  loadImage("assets/Tuomio2.png"),
		metsa:       loadImage("assets/Metsa.png"),
		p1:          loadImage("assets/Pysakki1.png"),
		p2:          loadImage("assets/Pysakki2 (1).png"),
		p22:         loadImage("assets/Pysakki22.png"),
		puska:       loadImage("assets/Puska.png"),
		posankka:    loadImage("assets/Pos.png"),
		mk:          loadImage("assets/Metsakoko.png"),
		l1:          loadImage("assets/Lamppu1.png"),
		valo:        loadImage("assets/Valo.png"),
		mode:        loadImage("assets/Mode.png"),

		//yömode
		backgroundN: loadImage("assets/Tuomio2 (1).png"),
		metsaN:      loadImage("assets/MetsaN.png"),
		p1N:         loadImage("assets/Pysakki1N.png"),
		p2N:         loadImage("assets/Pysakki2N.png"),
		posN:        loadImage("assets/PosN.png"),
		mkN:         loadImage("assets/MetsakokoN.png"),
		l1N:         loadImage("assets/Lamppu1N.png"),

		Tulossa1L: loadImage("assets/Tulossa1L.png"),
		Tulossa1LN: loadImage("assets/Tulossa1LN.png"),

		Tulossa2: loadImage("assets/Tulossa2.png"),
		Tulossa2N: loadImage("assets/Tulossa2N.png"),

		Tietyokyltti: loadImage("assets/Tietyokyltti.png"),
		TietyokylttiN: loadImage("assets/TietyokylttiN.png"),

		cone: loadImage("assets/cone.png"),
		coneN: loadImage("assets/coneN.png"),

		holes: loadImage("assets/holes.png"),

		width:  		spriteWidth,
		height: 		spriteHeight,

		camX: 			0,
		camY: 			0,

		busX: 			busStartX,
		busY: 			busStartY,

		messageAlpha: 	1,
	}
}