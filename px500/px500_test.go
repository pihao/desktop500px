package px500

import (
	"testing"
)

func Test_randPhoto(t *testing.T) {
	photos := Photos{
		Photos: []Photo{
			Photo{
				Width:  1,
				Height: 2,
			},
			Photo{
				Width:  3,
				Height: 4,
			},
			Photo{
				Width:  6,
				Height: 5,
			},
		},
	}
	p := randPhoto(&photos)
	if p.Width != 6 || p.Height != 5 {
		t.Error("The width must larger than the height.")
	}
}
