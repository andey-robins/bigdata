package sentences

import "testing"

func TestCountDupes(t *testing.T) {
	ss := New()
	ss.LoadFile("../sentence_files/tiny.txt") // this file has 2 duplicate sentences

	count := ss.CountDupes()
	if count != 2 {
		t.Errorf("[Test 1] - got=%v, exp=%v\n", count, 2)
	}

	ss = New()
	ss.LoadFile("../sentence_files/small.txt") // this file has 1 duplicate sentence

	count = ss.CountDupes()
	if count != 1 {
		t.Errorf("[Test 2] - got=%v, exp=%v\n", count, 1)
	}
}
