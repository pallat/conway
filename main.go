package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type cell struct {
	x           int
	y           int
	alive       bool
	topleft     *cell
	top         *cell
	topright    *cell
	left        *cell
	right       *cell
	bottomleft  *cell
	bottom      *cell
	bottomright *cell
}

func main() {
	s := New(80, 30)
	s.draw()

	for i := 0; i < 1000; i++ {
		s.pixels.next()
		time.Sleep(time.Second)
		s.draw()
	}
}

func (pixels conways) born() {
	for i := range pixels {
		if rand.Intn(20) == 2 {
			pixels[i].alive = true
		}
	}
}

func initial(dx, dy int) conways {
	total := dx * dy
	pixels := make([]*cell, total, total)
	for i := range pixels {
		pixels[i] = &cell{x: i % dx, y: i / dx}
	}
	return pixels
}

type conways []*cell

func (pixels conways) connect(dx, dy int) {
	total := dx * dy

	for i := range pixels {
		topleft := i - dx - 1
		if topleft > 0 {
			pixels[i].topleft = pixels[topleft]
		}

		top := i - dx
		if top > 0 {
			pixels[i].top = pixels[top]
		}

		topright := i - dx + 1
		if topright > 0 {
			pixels[i].topright = pixels[topright]
		}

		left := i - 1
		if left > 0 {
			pixels[i].left = pixels[left]
		}

		right := i + 1
		if right < total {
			pixels[i].right = pixels[right]
		}

		bottomleft := i + dx - 1
		if bottomleft < total {
			pixels[i].bottomleft = pixels[bottomleft]
		}

		bottom := i + dx
		if bottom < total {
			pixels[i].bottom = pixels[bottom]
		}

		bottomright := i + dx + 1
		if bottomright < total {
			pixels[i].bottomright = pixels[bottomright]
		}

	}
}

func New(dx, dy int) space {
	rand.Seed(time.Now().UnixNano())

	pixels := initial(dx, dy)
	pixels.born()
	pixels.connect(dx, dy)

	return space{
		dx:     dx,
		dy:     dy,
		pixels: pixels,
	}
}

type space struct {
	dx     int
	dy     int
	pixels conways
}

func (s *space) draw() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	i := 0
	for y := 0; y < s.dy; y++ {
		for x := 0; x < s.dx; x++ {
			if s.pixels[i].alive {
				fmt.Print("0")
			} else {
				fmt.Print(" ")
			}
			i++
		}
		fmt.Println()
	}
}

func (cells conways) next() {
	current := make([]*cell, len(cells))
	for i := range cells {
		current[i] = &cell{}
		current[i].x = cells[i].x
		current[i].y = cells[i].y
		current[i].alive = cells[i].alive
		current[i].topleft = &cell{}
		if cells[i].topleft != nil {
			*current[i].topleft = *cells[i].topleft
		}
		current[i].top = &cell{}
		if cells[i].top != nil {
			*current[i].top = *cells[i].top
		}
		current[i].topright = &cell{}
		if cells[i].topright != nil {
			*current[i].topright = *cells[i].topright
		}
		current[i].left = &cell{}
		if cells[i].left != nil {
			*current[i].left = *cells[i].left
		}
		current[i].right = &cell{}
		if cells[i].right != nil {
			*current[i].right = *cells[i].right
		}
		current[i].bottomleft = &cell{}
		if cells[i].bottomleft != nil {
			*current[i].bottomleft = *cells[i].bottomleft
		}
		current[i].bottom = &cell{}
		if cells[i].bottom != nil {
			*current[i].bottom = *cells[i].bottom
		}
		current[i].bottomright = &cell{}
		if cells[i].bottomright != nil {
			*current[i].bottomright = *cells[i].bottomright
		}
	}

	for i, v := range current {
		neighbours := 0

		if v.topleft != nil {
			if v.topleft.alive {
				// fmt.Println("topleft alive of:", i)
				neighbours++
			}
		}

		if v.top != nil {
			if v.top.alive {
				// fmt.Println("top alive of:", i)
				neighbours++
			}
		}

		if v.topright != nil {
			if v.topright.alive {
				// fmt.Println("topright alive of:", i)
				neighbours++
			}
		}

		if v.left != nil {
			if v.left.alive {
				// fmt.Println("left alive of:", i)
				neighbours++
			}
		}

		if v.right != nil {
			if v.right.alive {
				// fmt.Println("right alive of:", i)
				neighbours++
			}
		}

		if v.bottomleft != nil {
			if v.bottomleft.alive {
				// fmt.Println("bottomleft alive of:", i)
				neighbours++
			}
		}

		if v.bottom != nil {
			if v.bottom.alive {
				// fmt.Println("bottom alive of:", i)
				neighbours++
			}
		}

		if v.bottomright != nil {
			if v.bottomright.alive {
				// fmt.Println("bottomright alive of:", i)
				neighbours++
			}
		}

		if current[i].alive {
			if neighbours != 2 && neighbours != 3 {
				cells[i].alive = false
			}
		} else {
			if neighbours == 3 {
				cells[i].alive = true
			}
		}
	}
}
