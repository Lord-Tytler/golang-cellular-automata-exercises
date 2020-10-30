package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var running bool = true

type direction int

const (
	LEFT  direction = 0
	UP    direction = 1
	RIGHT direction = 2
	DOWN  direction = 3
)

const (
	BLACK = 0x00000000
	WHITE = 0xFFFFFFFF
	RED   = 0xFFFF0000
	GREEN = 0xFFFFA500
)

type ant struct {
	x   int
	y   int
	dir direction
}

const tileSize = 1
const screenWidth, screenHeight = 750, 750
const width, height = screenWidth / tileSize, screenHeight / tileSize

var tiles [width][height]int

var window sdl.Window
var surface sdl.Surface

func turnRight(a *ant) {
	switch a.dir {
	case UP:
		a.dir = RIGHT
	case RIGHT:
		a.dir = DOWN
	case DOWN:
		a.dir = LEFT
	case LEFT:
		a.dir = UP
	}
	/*
		a.dir++
		if a.dir > 3 {
			a.dir = 0
		} */
}
func turnLeft(a *ant) {

	switch a.dir {
	case UP:
		a.dir = LEFT
	case LEFT:
		a.dir = DOWN
	case DOWN:
		a.dir = RIGHT
	case RIGHT:
		a.dir = UP
	}
	/*a.dir--
	if a.dir < 0 {
		a.dir = 3
	} */
}
func flipTile(a *ant) {
	row := a.x
	col := a.y
	switch tiles[row][col] {
	case WHITE:
		tiles[row][col] = BLACK
	case BLACK:
		tiles[row][col] = WHITE
	}
}
func forward(a *ant) {
	switch a.dir {
	case UP:
		a.y--
	case DOWN:
		a.y++
	case LEFT:
		a.x--
	case RIGHT:
		a.x++
	}
}
func updateAnt(a *ant) {
	switch tiles[a.x][a.y] {
	case WHITE:
		turnRight(a)
	case BLACK:
		turnLeft(a)
	}
	flipTile(a)
	forward(a)

}
func fillTiles() {
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			tiles[row][col] = WHITE
		}
	}
}
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
	fillTiles()
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			rect := sdl.Rect{int32(row * tileSize), int32(col * tileSize), tileSize, tileSize}
			switch tiles[row][col] {
			case WHITE:
				surface.FillRect(&rect, WHITE)
			case BLACK:
				surface.FillRect(&rect, BLACK)
			}
		}
	}
	a := ant{width / 2, height / 2, DOWN}
	var fps int64 = 120000
	step := (int64)(1000000000 / fps)
	frames := 0
	last := time.Now().UnixNano()
	for true {
		if time.Now().UnixNano()-last >= step {
			frames++
			updateAnt(&a)
			for row := 0; row < height; row++ {
				for col := 0; col < width; col++ {
					rect := sdl.Rect{int32(row * tileSize), int32(col * tileSize), tileSize, tileSize}
					switch tiles[row][col] {
					case WHITE:

						surface.FillRect(&rect, WHITE)
					case BLACK:

						surface.FillRect(&rect, BLACK)
					}
					if a.x == row && a.y == col {
						surface.FillRect(&rect, RED)
					}
				}
			}

			window.UpdateSurface()
			last = time.Now().UnixNano()
		}
	}
	//sdl.Delay(3000)
	sdl.Quit()
}
