package snek

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"runtime"
	"time"
)

const (
	width  = 700
	height = 700
	factor = 50
	fps    = 10

	vertexShaderSource = `
		#version 410
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(1, 1, 1, 1.0);
		}
	` + "\x00"
)

var isInPlay = true
var gameStart = true

func Play() {
	runtime.LockOSThread()

	window := initGlfw()
	window.SetKeyCallback(keyCallback)

	defer glfw.Terminate()
	program := initOpenGL()

	cells := makeCells()
	food := initializeFood()
	snake := makeSnake()

	for !window.ShouldClose() {
		if keyInput == 'e' {
			isInPlay = true
		}
		if isInPlay {
			isInPlay = playGame(snake, cells, food, window, program)
		} else {
			drawWindow(window, program)
			food = initializeFood()
			snake = makeSnake()
		}
	}
}

func playGame(s *snake, cells [][]*cell, f *food, window *glfw.Window, program uint32) bool {
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

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Snake", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}
