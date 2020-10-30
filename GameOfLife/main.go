package main

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type cell struct {
	x      int
	y      int
	alive  bool
	target bool
}

var running bool = true

const squareSize = 4
const screenWidth, screenHeight = 1280, 800
const width, height = screenWidth / squareSize, screenHeight / squareSize

var cells [height][width]cell
var window sdl.Window
var surface sdl.Surface

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	rect := sdl.Rect{0, 0, screenWidth, screenHeight}

	surface.FillRect(&rect, 0xffffffff)

	var fps int64 = 60
	step := (int64)(1000000000 / fps)
	last := time.Now().UnixNano()
	fillCells()
	for true {
		if time.Now().UnixNano()-last >= step {
			update()
			for i := 0; i < len(cells); i++ {
				for j := 0; j < len(cells[0]); j++ {
					rect1 := sdl.Rect{int32(j * squareSize), int32(i * squareSize), squareSize, squareSize}
					if cells[i][j].alive {
						surface.FillRect(&rect1, 0xffffffff)
					} else {
						surface.FillRect(&rect1, 0x00000000)
					}
				}
			}
			window.UpdateSurface()
			last = time.Now().UnixNano()
		}
	}
	sdl.Delay(3000)
	sdl.Quit()
}
func run() {
	var fps int64 = 10
	step := (int64)(1000000000 / fps)
	last := time.Now().UnixNano()
	fillCells()
	for true {
		if time.Now().UnixNano()-last >= step {
			update()
			display()
			last = time.Now().UnixNano()
		}
	}
}

func update() {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if cells[i][j].alive {
				alive := getAliveNeighbors(i, j)
				if alive < 2 {
					cells[i][j].target = false
				} else if alive > 3 {
					cells[i][j].target = false
				}
			} else {
				alive := getAliveNeighbors(i, j)
				if alive == 3 {
					cells[i][j].target = true
				}
			}

		}

	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			cells[i][j].alive = cells[i][j].target
		}
	}
}
func fillCells() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			a := rand.Intn(101)
			if a <= 20 {
				cells[i][j] = cell{i, j, true, true}
			} else {
				cells[i][j] = cell{i, j, false, false}
			}
		}
	}
}
func getAliveNeighbors(x int, y int) int {
	alive := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i != x || j != y {
				if i >= 0 && i < height && j >= 0 && j < width {
					if cells[i][j].alive {
						alive++
					}
				}
			}
		}
	}
	return alive
}
func drawString() [height][width]string {
	var output [height][width]string
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if cells[i][j].alive {
				output[i][j] = "*"
			} else {
				output[i][j] = " "
			}
		}

	}
	return output
}
func display() {
	rect1 := sdl.Rect{0, 0, screenWidth / 2, screenHeight / 2}
	rect2 := sdl.Rect{screenWidth / 2, 0, screenWidth / 2, screenHeight / 2}
	rect3 := sdl.Rect{0, screenHeight / 2, screenWidth / 2, screenHeight / 2}
	rect4 := sdl.Rect{screenWidth / 2, screenHeight / 2, screenWidth / 2, screenHeight / 2}

	surface.FillRect(&rect1, 0xffff0000)
	surface.FillRect(&rect2, 0xffffff00)
	surface.FillRect(&rect3, 0xff00ff00)
	surface.FillRect(&rect4, 0xff0000ff)
	window.UpdateSurface()
}
