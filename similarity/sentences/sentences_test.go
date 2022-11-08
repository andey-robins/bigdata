package sentences

import (
	"crypto/sha256"
	"os"
	"testing"
)

func TestCountDupes(t *testing.T) {
	file, err := os.Stat("../sentence_files/tiny.txt")
	if err != nil {
		t.Error(err)
	}
	ss := New(int(file.Size()/32), sha256.Sum256)
	ss.LoadFile("../sentence_files/tiny.txt") // this file has 2 duplicate sentences

	count := ss.CountDupes()
	if count != 2 {
		t.Errorf("[Test 1] - got=%v, exp=%v\n", count, 2)
	}

	file, err = os.Stat("../sentence_files/small.txt")
	if err != nil {
		t.Error(err)
	}
	ss = New(int(file.Size()/32), sha256.Sum256)
	ss.LoadFile("../sentence_files/small.txt") // this file has 1 duplicate sentence

	count = ss.CountDupes()
	if count != 1 {
		t.Errorf("[Test 2] - got=%v, exp=%v\n", count, 1)
	}

	ss = New(1000, sha256.Sum256)
	ss.LoadFile("../sentence_files/1k.txt") // this file has 1 duplicate sentence

	count = ss.CountDupes()
	if count != 34 {
		t.Errorf("[Test 2] - got=%v, exp=%v\n", count, 34)
	}
}
