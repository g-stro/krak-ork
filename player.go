package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type player struct {
	x      float64
	y      float64
	radius float64
	img    *ebiten.Image
	clr    color.Color
}

func (p *player) update() {
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

	// Collision detection
	if p.colliding(op) {
		p.resolveCollision(op)
		log.Println("COLLIDING")
	}
}

// colliding checks whether a player is colliding with another player
func (p player) colliding(op player) bool {
	distX := (p.x + p.radius) - (op.x + op.radius)
	distY := (p.y + p.radius) - (op.y + op.radius)
	dist := math.Sqrt(distX*distX + distY*distY)

	return dist < p.radius+op.radius
}

// resolveCollision updates a player's x and y values
func (p *player) resolveCollision(op player) {
	distX := (p.x + p.radius) - (op.x + op.radius)
	distY := (p.y + p.radius) - (op.y + op.radius)
	dist := math.Sqrt(distX*distX + distY*distY)
	radii := p.radius + op.radius

	opX := distX / dist
	opY := distY / dist

	p.x = op.x + (radii+1)*opX
	p.y = op.y + (radii+1)*opY
}
