package conway_test

import (
	"testing"
	"time"

	"github.com/pallat/conway/conway"
)

func TestRun(t *testing.T) {
	d := conway.NewDimension(80, 30)
	s := conway.New(d)
	s.Draw()

	for i := 0; i < 1000; i++ {
		s.Pixels.Next()
		time.Sleep(time.Second)
		s.Draw()
	}
}
