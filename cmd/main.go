package main

import (
	"ebiten-bunnymark/pkg/bunny"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 300
	screenHeight = 300
)

var sprites *ebiten.Image
var pressed = false

func init() {
	var err error
	sprites, _, err = ebitenutil.NewImageFromFile("assets/lineup.png")
	if err != nil {
		log.Fatal(err)
	}
}

var prevUpdateTime = time.Now()

var bunnies []*bunny.Bunny

const gravity = 0.05

type Game struct {
	pressedKeys []ebiten.Key
	touchIDs    []ebiten.TouchID
	touchID     ebiten.TouchID
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
func createRandomizeBunny(tint int) *bunny.Bunny {
	b := &bunny.Bunny{0, 0, 0.0, 0.0, 0}
	b.PositionX = rand.Float64()*(4.8-2.01) + 0.01
	b.PositionY = rand.Float64()*(4.8-2.01) + 0.01
	b.SpeedX = rand.Float64()*(0.8-0.01) + 0.01
	b.SpeedY = rand.Float64()*(0.8-0.01) + 0.01
	b.Theme = tint
	return b
}

func (g *Game) Update() error {
	timeDelta := float64(time.Since(prevUpdateTime).Milliseconds())
	prevUpdateTime = time.Now()
	theme := rand.Intn(11-0) + 0

	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])
	g.touchIDs = inpututil.AppendJustPressedTouchIDs(g.touchIDs[:0])
	for _, key := range g.pressedKeys {
		switch key.String() {
		case "Space":
			for i := 0; i < 10; i++ {
				bunnies = append(bunnies, createRandomizeBunny(theme))
			}

		}
	}
	if pressed {
		theme := rand.Intn(11-0) + 0
		for i := 0; i < 10; i++ {
			bunnies = append(bunnies, createRandomizeBunny(theme))
		}
	}
	for _, key := range g.touchIDs {
		if inpututil.TouchPressDuration(key) > 0 {
			pressed = true
			g.touchID = key
		}
	}
	if inpututil.IsTouchJustReleased(g.touchID) {
		pressed = false
	}

	for _, b := range bunnies {
		b.Update(screenWidth+36, screenHeight+36, timeDelta, gravity)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, b := range bunnies {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(36/2), -float64(36/2))
		op.GeoM.Translate(b.PositionX, b.PositionY)
		screen.DrawImage(sprites.SubImage(image.Rect(b.Theme*36, 0, b.Theme*36+36, 36)).(*ebiten.Image), op)
	}

	msg := fmt.Sprintf(
		`TPS: %0.2f
FPS: %0.2f
Num of Bunnies: %d
Press "Space" or "Tap" to add bunnies`, ebiten.CurrentTPS(), ebiten.CurrentFPS(), len(bunnies))
	ebitenutil.DebugPrint(screen, msg)
}
func main() {

	theme := rand.Intn(11-0) + 0
	for i := 0; i < 10; i++ {
		bunnies = append(bunnies, createRandomizeBunny(theme))
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ebiten Test")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
