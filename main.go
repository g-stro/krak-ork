package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

type player struct {
	x      float64
	y      float64
	radius float64
	img    *ebiten.Image
	clr    color.Color
}

var p player

// Update the game state
func (g *Game) Update() error {
	// Keyboard input for player movement
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.y--
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.y++
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.x--
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.x++
	}

	return nil
}

// drawCircle draws a circle to the screen
func drawCircle(screen *ebiten.Image, p player) {
	angle := math.Acos(1 - 1/p.radius)
	// iterate over each angle until the circle is complete
	for a := float64(0); a <= 360; a += angle {
		// Delta x and y values
		xd := p.radius * math.Cos(a)
		yd := p.radius * math.Sin(a)
		// draw the coordinates to the screen
		x := int(math.Round(p.x + xd))
		y := int(math.Round(p.y + yd))
		screen.Set(x, y, p.clr)
	}
}

// Draw renders graphics to the screen
func (g *Game) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct
	opts.GeoM.Translate(p.x, p.y)
	drawCircle(screen, p)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("KRAK ORRRRRRK!")
	// Create the player
	p = player{64, 48, 16, ebiten.NewImage(32, 32), color.RGBA{0, 0xff, 0, 0xff}}
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
