package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/pallat/conway/conway"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Conway's Game Of Life",
		Bounds: pixel.R(0, 0, 640, 480),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	d := conway.NewDimension(64, 48)
	gof := conway.New(d)
	cells := gof.Cells()

	for !win.Closed() {
		buffer := []*imdraw.IMDraw{}
		win.Clear(colornames.Whitesmoke)
		i := 0
		for y := 0; y < d.Length; y++ {
			for x := 0; x < d.Width; x++ {
				if cells[i].Alive {
					imd := imdraw.New(nil)
					imd.Color = colornames.Black
					imd.EndShape = imdraw.NoEndShape
					imd.Push(pixel.V(float64(x*10), float64(y*10)), pixel.V(float64((x*10)+9), float64(y*10)+9))
					imd.Rectangle(0)
					imd.Draw(win)
					buffer = append(buffer, imd)
				}
				i++
			}
		}

		win.Update()

		time.Sleep(time.Millisecond * 300)
		for i := range buffer {
			buffer[i].Clear()
		}

		cells.Next()
	}
}

func main() {
	pixelgl.Run(run)
}
