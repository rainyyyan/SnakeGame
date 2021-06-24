package snek

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type snake struct {
	length    int
	body      [][]int
	direction rune
}

var keyInput rune

var dirOffset = map[rune][]int{
	'l': {-1, 0},
	'r': {1, 0},
	'u': {0, 1},
	'd': {0, -1},
	'n': {0, 0},
}

func makeSnake() *snake {
	x, y := randomSpawn()

	return &snake{
		length:    1,
		body:      [][]int{{x, y}},
		direction: 'n',
	}
}

func (s *snake) move() {
	if _, err := dirOffset[keyInput]; err == true {
		s.direction = keyInput
	}
	s.moveInDirection(s.direction)
}

func (s *snake) isOutOfBounds() bool {
	if s.body[0][0] >= factor || s.body[0][1] >= factor ||
		s.body[0][0] < 0 || s.body[0][1] < 0 {
		return true
	}
	return false
}

func (s *snake) moveInDirection(direction rune) {
	move := dirOffset[direction]
	newBody := [][]int{{s.body[0][0] + move[0], s.body[0][1] + move[1]}}
	for i, v := range s.body {
		if i < len(s.body)-1 {
			newBody = append(newBody, v)
		}
	}
	s.body = newBody
}

func (s *snake) eat(f *food) {
	head := s.body[0]
	if head[0] == f.x && head[1] == f.y {
		f.move()
		s.grow()
	}
}

func (s *snake) grow() {
	tail := s.body[len(s.body)-1]
	nextSegment := []int{tail[0] + -1*dirOffset[s.direction][0], tail[1] + -1*dirOffset[s.direction][1]}
	s.body = append(s.body, nextSegment)
}

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
