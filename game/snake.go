package game

type snake struct {
	length    int
	body      [][]int
	direction rune
}

var dirOffset = map[rune][]int{
	'l': {-1, 0},
	'r': {1, 0},
	'u': {0, 1},
	'd': {0, -1},
	'n': {0, 0},
}

// creates new snake at random location
func makeSnake() *snake {
	x, y := randomSpawn()
	keyInput = ' '

	return &snake{
		length:    1,
		body:      [][]int{{x, y}},
		direction: 'n',
	}
}

// checks if snake is going out of bounds
func (s *snake) isOutOfBounds() bool {
	if s.body[0][0] >= boardSize || s.body[0][1] >= boardSize ||
		s.body[0][0] < 0 || s.body[0][1] < 0 {
		return true
	}
	return false
}

// moves snake in specified direction
func (s *snake) move() {
	changeSnakeDirection(s)
	move := dirOffset[s.direction]
	newBody := [][]int{{s.body[0][0] + move[0], s.body[0][1] + move[1]}}
	for i, v := range s.body {
		if i < len(s.body)-1 {
			newBody = append(newBody, v)
		}
	}
	s.body = newBody
}

// determines when snake collides with food
// snake grows and food moves if collision happens
func (s *snake) eat(f *food) {
	head := s.body[0]
	if head[0] == f.x && head[1] == f.y {
		f.changeFoodLocation(s)
		s.grow()
	}
}

// adds new segment onto the end of snake body with
// respect to the current direction the snake is moving in
func (s *snake) grow() {
	tail := s.body[len(s.body)-1]
	// the next segment is behind the current tail, in the opposite direction the snake is moving in
	nextSegment := []int{tail[0] + -1*dirOffset[s.direction][0], tail[1] + -1*dirOffset[s.direction][1]}
	s.body = append(s.body, nextSegment)
}
