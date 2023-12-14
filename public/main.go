// go:build js && wasm

package main

import (
	"math/rand"
	"syscall/js"

	"snake.ham.nfaschool.org/public/lib"
)

const GRIDWIDTH int = 16
const GRIDHEIGHT int = 16

type SnakePoint struct {
	X, Y float64
}

func (s *SnakePoint) GetGridCell() (int, int) {
	return int(s.X / float64(GRIDWIDTH)), int(s.Y / float64(GRIDHEIGHT))
}

type Snake struct {
	Body  []SnakePoint
	Speed float64
}

func (s *Snake) Contains(pos int) bool {
	var x, y int = pos % GRIDWIDTH, pos / GRIDWIDTH

	for _, point := range s.Body {
		var px, py int = int(point.X) / GRIDWIDTH, int(point.Y) / GRIDHEIGHT

		if px == x && py == y {
			return true
		}
	}

	return false
}

var fruitGrid []bool = make([]bool, GRIDWIDTH*GRIDHEIGHT)

func randomFruitPos(snake *Snake) int {
	var pos int = -1

	for pos == -1 || fruitGrid[pos] || snake.Contains(pos) {
		pos = rand.Intn(GRIDWIDTH * GRIDHEIGHT)
	}

	return pos
}

func main() {
	var keepAliveChannel chan bool = make(chan bool)
	var canvas *lib.CanvasStruct = lib.NewCanvas(js.Global().Get("document").Call("querySelector", "canvas"))

	var snake *Snake = &Snake{
		Body: make([]SnakePoint, 1),
	}

	snake.Body[0] = SnakePoint{
		X: float64(GRIDWIDTH / 2),
		Y: float64(GRIDHEIGHT / 2),
	}

	var fruitPos int = randomFruitPos(snake)

	lib.SetupRenderingFrame(canvas, func(width, height float64) {
		if fruitGrid[fruitPos] == false {
			fruitPos = randomFruitPos(snake)
			fruitGrid[fruitPos] = true
		}

		var scale float64 = lib.UiScale(canvas)

		canvas.ClearRect(0, 0, width, height)

		canvas.Save()
		canvas.Scale(scale, scale)

		// Get scale of grid
		var gridSize float64 = width / 2 / float64(GRIDWIDTH)

		// Draw grid
		canvas.SetStroke("black")
		canvas.BeginPath()
		for i := 0; i <= GRIDWIDTH; i++ {
			canvas.MoveTo(float64(i)*gridSize, 0)
			canvas.LineTo(float64(i)*gridSize, float64(GRIDHEIGHT)*gridSize)
		}
		for i := 0; i <= GRIDHEIGHT; i++ {
			canvas.MoveTo(0, float64(i)*gridSize)
			canvas.LineTo(float64(GRIDWIDTH)*gridSize, float64(i)*gridSize)
		}

		canvas.Stroke()

		// Draw fruit
		canvas.SetFill("red")

		for i := 0; i < GRIDWIDTH*GRIDHEIGHT; i++ {
			if fruitGrid[i] {
				var x, y int = i % GRIDWIDTH, i / GRIDWIDTH
				canvas.FillRect(float64(x)*gridSize, float64(y)*gridSize, gridSize, gridSize)
			}
		}

		canvas.Restore()
	})

	<-keepAliveChannel
}
