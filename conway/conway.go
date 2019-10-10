package conway

import (
	"math/rand"
	"time"
)

type Dimension struct {
	Width  int
	Length int
}

func NewDimension(w, l int) Dimension {
	return Dimension{
		Width:  w,
		Length: l,
	}
}

type cell struct {
	Alive       bool
	x           int
	y           int
	topleft     *cell
	top         *cell
	topright    *cell
	left        *cell
	right       *cell
	bottomleft  *cell
	bottom      *cell
	bottomright *cell
}

func (pixels Cells) born() {
	for i := range pixels {
		if rand.Intn(10) == 2 {
			pixels[i].Alive = true
		}
	}
}

func initial(dx, dy int) Cells {
	total := dx * dy
	pixels := make([]*cell, total, total)
	for i := range pixels {
		pixels[i] = &cell{x: i % dx, y: i / dx}
	}
	return pixels
}

type Cells []*cell

func (pixels Cells) connect(dx, dy int) {
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

func New(d Dimension) Space {
	rand.Seed(time.Now().UnixNano())

	pixels := initial(d.Width, d.Length)
	pixels.born()
	pixels.connect(d.Width, d.Length)

	return Space{
		dx:     d.Width,
		dy:     d.Length,
		Pixels: pixels,
	}
}

type Space struct {
	dx     int
	dy     int
	Pixels Cells
}

func (s *Space) Cells() Cells {
	return s.Pixels
}

func (cells Cells) Next() {
	current := make([]*cell, len(cells))
	for i := range cells {
		current[i] = &cell{}
		current[i].x = cells[i].x
		current[i].y = cells[i].y
		current[i].Alive = cells[i].Alive
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
			if v.topleft.Alive {
				neighbours++
			}
		}

		if v.top != nil {
			if v.top.Alive {
				neighbours++
			}
		}

		if v.topright != nil {
			if v.topright.Alive {
				neighbours++
			}
		}

		if v.left != nil {
			if v.left.Alive {
				neighbours++
			}
		}

		if v.right != nil {
			if v.right.Alive {
				neighbours++
			}
		}

		if v.bottomleft != nil {
			if v.bottomleft.Alive {
				neighbours++
			}
		}

		if v.bottom != nil {
			if v.bottom.Alive {
				neighbours++
			}
		}

		if v.bottomright != nil {
			if v.bottomright.Alive {
				neighbours++
			}
		}

		if current[i].Alive {
			if neighbours != 2 && neighbours != 3 {
				cells[i].Alive = false
			}
		} else {
			if neighbours == 3 {
				cells[i].Alive = true
			}
		}
	}
}
