package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

type player struct {
	x   float64
	y   float64
	img *ebiten.Image
}

var p player

// Update the game state
func (g *Game) Update() error {
	p.img.Fill(color.RGBA{0, 0xff, 0, 0xff})

	// Keyboard input
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

// Render to the screen
func (g *Game) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	// Add the Translate effect to the option struct
	opts.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.img, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("KRAK ORRRRRRK!")
	// Create the player
	p = player{64, 48, ebiten.NewImage(32, 32)}
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
