package sentences

import "testing"

func TestEditDistance(t *testing.T) {
	tests := []struct {
		s1  string
		s2  string
		exp int
	}{
		{
			"apples bananas cherries",
			"apples bananas cherries dates",
			1,
		},
		{
			"apples bananas cherries",
			"apples bananas",
			1,
		},
		{
			"apples bananas cherries",
			"apples bananas dates",
			2,
		},
		{
			"apples bananas cherries eggplant",
			"apples bananas dates",
			3,
		},
	}

	for i, test := range tests {
		if dist := EditDistance(test.s1, test.s2); dist != test.exp {
			t.Errorf("[Test %v] - wrong distance found. got=%v, exp=%v", i, dist, test.exp)
		}
	}
}
