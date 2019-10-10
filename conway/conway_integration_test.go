// +build integration

package conway_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/pallat/conway/conway"
)

func TestRun(t *testing.T) {
	d := conway.NewDimension(80, 30)
	gof := conway.New(d)
	cells := gof.Cells()

	for i := 0; i < 1000; i++ {
		cells.Next()
		time.Sleep(time.Second)
		Draw(cells, d)
	}
}

func Draw(cells conway.Cells, d conway.Dimension) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	i := 0
	for y := 0; y < d.Length; y++ {
		for x := 0; x < d.Width; x++ {
			if cells[i].Alive {
				fmt.Print("✸")
			} else {
				fmt.Print(" ")
			}
			i++
		}
		fmt.Println()
	}
}
