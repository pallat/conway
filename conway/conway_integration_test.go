package conway_test

import (
	"testing"
	"time"

	"github.com/pallat/conway/conway"
)

func TestRun(t *testing.T) {
	s := conway.New(80, 30)
	s.Draw()

	for i := 0; i < 1000; i++ {
		s.Pixels.Next()
		time.Sleep(time.Second)
		s.Draw()
	}
}
