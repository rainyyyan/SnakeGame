package game

// from https://kylewbanks.com/blog/tutorial-opengl-with-golang-part-2-drawing-the-game-board
// makes the cells of the game board

// points to draw a square
var (
	square = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,

		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		0.5, -0.5, 0,
	}
)

type cell struct {
	drawable uint32

	x int
	y int
}

// make all cells on the board
func makeCells() [][]*cell {
	cells := make([][]*cell, boardSize, boardSize)
	for x := 0; x < boardSize; x++ {
		for y := 0; y < boardSize; y++ {
			c := newCell(x, y)
			cells[x] = append(cells[x], c)
		}
	}

	return cells
}

// creates and returns a new cell in row x column y
func newCell(x, y int) *cell {
	points := make([]float32, len(square), len(square))
	copy(points, square)

	for i := 0; i < len(points); i++ {
		var position float32
		var size float32

		//determine x and y position for each cell (case 0 and 1) in pixels
		switch i % 3 {
		case 0:
			size = 1.0 / float32(boardSize)
			position = float32(x) * size
		case 1:
			size = 1.0 / float32(boardSize)
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

	return &cell{
		drawable: makeVao(points),
		x:        x,
		y:        y,
	}
}
