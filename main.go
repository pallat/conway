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

	for i := 0; i < 10000; i++ {
		next(s.pixels)
		s.draw()
		time.Sleep(time.Second)
	}
}

func New(dx, dy int) space {
	rand.Seed(time.Now().UnixNano())

	total := dx * dy
	pixels := make([]*cell, total, total)

	for i := range pixels {
		pixels[i] = &cell{x: i % dx, y: i / dx}
		if rand.Intn(20) == 2 {
			pixels[i].alive = true
		}
	}

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

	return space{
		dx:     dx,
		dy:     dy,
		pixels: pixels,
	}
}

type space struct {
	dx     int
	dy     int
	pixels []*cell
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

func next(cells []*cell) {
	current := make([]*cell, len(cells))
	for i := range cells {
		current[i] = &cell{}
		*current[i] = *cells[i]
	}

	for i, v := range current {
		neighbours := 0

		if v.topleft != nil {
			if v.topleft.alive {
				neighbours++
			}
		}

		if v.top != nil {
			if v.top.alive {
				neighbours++
			}
		}

		if v.topright != nil {
			if v.topright.alive {
				neighbours++
			}
		}

		if v.left != nil {
			if v.left.alive {
				neighbours++
			}
		}

		if v.right != nil {
			if v.right.alive {
				neighbours++
			}
		}

		if v.bottomleft != nil {
			if v.bottomleft.alive {
				neighbours++
			}
		}

		if v.bottom != nil {
			if v.bottom.alive {
				neighbours++
			}
		}

		if v.bottomright != nil {
			if v.bottomright.alive {
				neighbours++
			}
		}

		if cells[i].alive {
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
