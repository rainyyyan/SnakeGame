package snek

type food struct {
	x int
	y int
}

// creates new food object
func initializeFood() *food {
	x, y := randomSpawn()
	return &food{
		x: x,
		y: y,
	}
}

// moves food to random location but not on top of snake
func (f *food) move(s *snake) {
	x, y := randomSpawn()
	for _, s := range s.body {
		if s[0] == x && s[1] == y {
			x, y = randomSpawn()
		}
	}

	f.x = x
	f.y = y
}
