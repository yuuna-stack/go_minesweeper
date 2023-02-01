package main

import (
	"math/rand"
	"path"
	"runtime"
	"time"

	"github.com/yuuna-stack/go_minesweeper/wrapper"

	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

const resourcesDir = "images"

type Point struct {
	x int
	y int
}

func init() { runtime.LockOSThread() }

func fullname(filename string) string {
	return path.Join(resourcesDir, filename)
}

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	resources := wrapper.Resources{}

	const gameWidth = 400
	const gameHeight = 400

	option := uint(window.SfResize | window.SfClose)
	wnd := wrapper.CreateWindow(gameWidth, gameHeight, "Minesweeper!", option, 60)

	w := 32
	grid := [12][12]int{}
	sgrid := [12][12]int{}

	s, err := wrapper.FileToSprite(fullname("tiles.jpg"), &resources)
	if err != nil {
		panic("Couldn't load tiles.jpg")
	}

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			sgrid[i][j] = 10
			if r1.Int()%5 == 0 {
				grid[i][j] = 9
			} else {
				grid[i][j] = 0
			}
		}
	}

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			n := 0
			if grid[i][j] == 9 {
				continue
			}
			if grid[i+1][j] == 9 {
				n++
			}
			if grid[i][j+1] == 9 {
				n++
			}
			if grid[i-1][j] == 9 {
				n++
			}
			if grid[i][j-1] == 9 {
				n++
			}
			if grid[i+1][j+1] == 9 {
				n++
			}
			if grid[i-1][j-1] == 9 {
				n++
			}
			if grid[i+1][j-1] == 9 {
				n++
			}
			if grid[i-1][j+1] == 9 {
				n++
			}
			grid[i][j] = n
		}
	}

	for wnd.IsOpen() {
		vec := graphics.SfMouse_getPositionRenderWindow(wnd.Get_Window())
		x := vec.GetX() / w
		y := vec.GetY() / w

		for wnd.Poll_Event() {
			if wnd.Close_Window() {
				return
			}
			if wnd.Mouse_ButtonPressed() {
				if wnd.Mouse_ButtonIs(window.SfMouseLeft) {
					sgrid[x][y] = grid[x][y]
				} else if wnd.Mouse_ButtonIs(window.SfMouseRight) {
					sgrid[x][y] = 11
				}
			}
		}

		for i := 1; i <= 10; i++ {
			for j := 1; j <= 10; j++ {
				if sgrid[i][j] == 9 {
					sgrid[i][j] = grid[i][j]
				}
				s.SetTextureRect(sgrid[i][j]*w, 0, w, w)
				s.SetPosition(float32(i*w), float32(j*w))
				s.Draw(wnd.Get_Window())
			}
		}

		graphics.SfRenderWindow_display(wnd.Get_Window())
	}

	resources.Clear()
	wnd.Clear()
}
