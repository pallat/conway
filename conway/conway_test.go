package conway

import "testing"

func TestInitial(t *testing.T) {
	px := initial(10, 10)

	if len(px) != 100 {
		t.Errorf("it should initial number of member to %d but got %d\n", 100, len(px))
	}
}

func TestConnectMember(t *testing.T) {
	px := initial(3, 3)
	px.connect(3, 3)

	t.Run("the right side is next index", func(t *testing.T) {
		for i := 0; i < 8; i++ {
			if px[i].right != px[i+1] {
				t.Errorf("member number %d from 3x3 should be the right side of member number %d but it was not", i+2, i+1)
			}
		}
	})
	t.Run("the left side is before index", func(t *testing.T) {
		for i := 0; i < 8; i++ {
			if px[i+1].left != px[i] {
				t.Errorf("member number %d from 3x3 should be the left side of member number %d but it was not", i+1, i+2)
			}
		}
	})
	t.Run("the top side", func(t *testing.T) {
		for i := 3; i < 9; i++ {
			if px[i].top != px[i-3] {
				t.Errorf("member number %d from 3x3 should be the top side of member number %d but it was not", i-2, i+1)
			}
		}
	})
	t.Run("the bottom side", func(t *testing.T) {
		for i := 0; i < 6; i++ {
			if px[i].bottom != px[i+3] {
				t.Errorf("member number %d from 3x3 should be the bottom side of member number %d but it was not", i+4, i+1)
			}
		}
	})

	t.Run("Diagonale", func(t *testing.T) {
		t.Run("the top left side", func(t *testing.T) {
			if px[4].topleft != px[0] {
				t.Errorf("member number %d from 3x3 should be the top left side of member number %d but it was not", 1, 5)
			}
			if px[5].topleft != px[1] {
				t.Errorf("member number %d from 3x3 should be the top left side of member number %d but it was not", 2, 6)
			}
			if px[7].topleft != px[3] {
				t.Errorf("member number %d from 3x3 should be the top left side of member number %d but it was not", 4, 8)
			}
			if px[8].topleft != px[4] {
				t.Errorf("member number %d from 3x3 should be the top left side of member number %d but it was not", 5, 9)
			}
		})
		t.Run("the top right side", func(t *testing.T) {
			if px[3].topright != px[1] {
				t.Errorf("member number %d from 3x3 should be the top right side of member number %d but it was not", 2, 4)
			}
			if px[4].topright != px[2] {
				t.Errorf("member number %d from 3x3 should be the top right side of member number %d but it was not", 3, 5)
			}
			if px[6].topright != px[4] {
				t.Errorf("member number %d from 3x3 should be the top right side of member number %d but it was not", 5, 7)
			}
			if px[7].topright != px[5] {
				t.Errorf("member number %d from 3x3 should be the top right side of member number %d but it was not", 6, 8)
			}
		})
		t.Run("the bottom left side", func(t *testing.T) {
			if px[1].bottomleft != px[3] {
				t.Errorf("member number %d from 3x3 should be the bottom left side of member number %d but it was not", 4, 2)
			}
			if px[2].bottomleft != px[4] {
				t.Errorf("member number %d from 3x3 should be the bottom left side of member number %d but it was not", 5, 3)
			}
			if px[4].bottomleft != px[6] {
				t.Errorf("member number %d from 3x3 should be the bottom left side of member number %d but it was not", 7, 5)
			}
			if px[5].bottomleft != px[7] {
				t.Errorf("member number %d from 3x3 should be the bottom left side of member number %d but it was not", 8, 6)
			}
		})
		t.Run("the bottom right side", func(t *testing.T) {
			if px[0].bottomright != px[4] {
				t.Errorf("member number %d from 3x3 should be the bottom right side of member number %d but it was not", 5, 1)
			}
			if px[1].bottomright != px[5] {
				t.Errorf("member number %d from 3x3 should be the bottom right side of member number %d but it was not", 6, 2)
			}
			if px[3].bottomright != px[7] {
				t.Errorf("member number %d from 3x3 should be the bottom right side of member number %d but it was not", 8, 4)
			}
			if px[4].bottomright != px[8] {
				t.Errorf("member number %d from 3x3 should be the bottom right side of member number %d but it was not", 9, 5)
			}
		})
	})
}
