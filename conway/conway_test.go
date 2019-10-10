package conway

import "testing"

func TestInitial(t *testing.T) {
	px := initial(10, 10)

	if len(px) != 100 {
		t.Errorf("it should initial number of member to %d but got %d\n", 100, len(px))
	}
}
