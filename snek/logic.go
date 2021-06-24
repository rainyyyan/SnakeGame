package snek

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"math/rand"
	"time"
)

var keyInput rune

// game logic to update game and end if applicable
func runGame(s *snake, cells [][]*cell, f *food, window *glfw.Window, program uint32) bool {
	t := time.Now()
	time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	s.move()
	if !s.isOutOfBounds() {
		s.eat(f)
		drawAll(s, cells, f, window, program)
		return true
	}
	return false
}

//set key callbacks
func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action,
	mods glfw.ModifierKey) {
	switch key {
	case glfw.KeyUp:
		keyInput = 'u'
	case glfw.KeyDown:
		keyInput = 'd'
	case glfw.KeyLeft:
		keyInput = 'l'
	case glfw.KeyRight:
		keyInput = 'r'
	case glfw.KeyEnter:
		keyInput = 'e'
	}
}

// generates random ints for location
func randomSpawn() (x, y int) {
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(factor-10) + 5
	y = rand.Intn(factor-10) + 5

	return
}
