package snek

import (
	"math"
	"math/rand"
	"time"
)

type food struct {
	drawable uint32

	x int
	y int
}

var circle = make([]float32, 0)

func (f *food) move() {
	rand.Seed(time.Now().UnixNano())
	newX := rand.Intn(factor - 10) + 5
	newY := rand.Intn(factor - 10) + 5

	f.x = newX
	f.y = newY
}

func makeCircle() {

	radius := 1 / float64(factor)
	numPoints := 30
	var x, y float32
	var theta float64 = 0
	var z float32 = 1

	for theta < 360 {
		x = float32(radius * math.Cos(theta * math.Pi / 180))
		y = float32(radius * math.Sin(theta * math.Pi / 180))

		circle = append(circle, x)
		circle = append(circle, y)
		circle = append(circle, z)

		theta = theta + (360 / float64(numPoints))
	}

}

func initializeFood() *food {
	x, y := randomSpawn()
	return newFood(x, y)
}

func newFood(x, y int) *food {
	makeCircle()

	points := make([]float32, len(circle), len(circle))
	copy(points, circle)

	for i := 0; i < len(points); i++ {
		var position float32
		var size float32
		switch i % 3 {
		case 0:
			size = 1 / float32(factor)
			position = float32(x) * size
		case 1:
			size = 1 / float32(factor)
			position = float32(y) * size
		default:
			continue
		}

		if points[i] < 0 {
			points[i] = (position * 2) - 1
		} else {
			points[i] = ((position + size) * 2) - 1
		}
	}

	return &food{
		drawable: makeVao(circle),
		x: x,
		y: y,
	}
}

func randomSpawn() (x, y int) {
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(factor - 10) + 5
	y = rand.Intn(factor - 10) + 5

	return
}