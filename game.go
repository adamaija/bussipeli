package main

import (
	"math/rand"
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
	RainPress   *ebiten.Image
	nightsky    *ebiten.Image
	lightning   *ebiten.Image
	blacksky    *ebiten.Image

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

	spacePressed bool

	messageAlpha float64 // 1.0 = fully visible, 0.0 = invisible

	RainDrops []RainDrop
    Splashes  []Splash

    SplashSheet *ebiten.Image

	rainEnabled bool
	rainKeyPressed bool
	rainVolume float64

	nightMode bool

	lightningFrames int
	lightningCoolDown int
	lightFlashFrames int

	rainIntensity float64
}

type RainDrop struct {
    X      float64
    Y      float64
    Speed  float64
    Length float64
	GroundY float64
    Alive  bool
}

type Splash struct {
    X      float64
    Y      float64

    Frame  int
    Timer  float64

    Alive  bool
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
		SplashSheet: loadImage("assets/grounddrops.png"),
		nightsky:    loadImage("assets/nightsky.png"),
		lightning:   loadImage("assets/lightning.png"),
		blacksky:    loadImage("assets/blacksky.png"),

		//yömode
		backgroundN: loadImage("assets/TuomioYoValmis.png"),
		metsaN:      loadImage("assets/MetsaN.png"),
		p1N:         loadImage("assets/Pysakki1N.png"),
		p2N:         loadImage("assets/Pysakki2N.png"),
		posN:        loadImage("assets/PosN.png"),
		mkN:         loadImage("assets/MetsakokoN.png"),
		l1N:         loadImage("assets/Lamppu1N.png"),
		RainPress:   loadImage("assets/RainPress.png"),

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

func (g *Game) SpawnRain() {
    d := RainDrop{
        X:       rand.Float64() * worldWidth,
        Y:       -10,
        Speed:   120 + rand.Float64()*80,
        Length:  4 + rand.Float64()*4,
        GroundY: 160 + rand.Float64()*20,
        Alive:   true,
    }

    g.RainDrops = append(g.RainDrops, d)
}

func (g *Game) SpawnSplash(x, y float64) {
    s := Splash{
        X:     x - 8,
        Y:     y - 14,
        Frame: 0,
        Timer: 0,
        Alive: true,
    }

    g.Splashes = append(g.Splashes, s)
}

func (g *Game) UpdateRain(dt float64) {
    for i := range g.RainDrops {
        d := &g.RainDrops[i]

        if !d.Alive {
            continue
        }
        d.Y += d.Speed * dt
        if d.Y >= d.GroundY {
            g.SpawnSplash(d.X, d.GroundY)
            d.Alive = false
        }
    }

    // Remove dead drops
    alive := g.RainDrops[:0]

    for _, d := range g.RainDrops {
        if d.Alive {
            alive = append(alive, d)
        }
    }

    g.RainDrops = alive
}

func (g *Game) UpdateSplashes(dt float64) {
    for i := range g.Splashes {
        s := &g.Splashes[i]
        if !s.Alive {
            continue
        }
        s.Timer += dt
        if s.Timer >= 0.05 {
            s.Timer = 0
            s.Frame++
            if s.Frame >= 4 {
                s.Alive = false
            }
        }
    }
    // Remove dead splashes
    alive := g.Splashes[:0]

    for _, s := range g.Splashes {
        if s.Alive {
            alive = append(alive, s)
        }
    }

    g.Splashes = alive
}