package main

import (
	"math"
	"testing"

	"github.com/andey-robins/bigdata/similarity/hash"
	"github.com/andey-robins/bigdata/similarity/sentences"
)

func TestDuplicates(t *testing.T) {
	tests := []struct {
		fname string
		count int
	}{
		{"sentence_files/trivial.txt", 1},
		{"sentence_files/tiny.txt", 2},
		{"sentence_files/small.txt", 1},
	}

	for i, test := range tests {
		ss := sentences.New(int(math.Pow(2, 8)), hash.Sha256Wrapper)
		ss.LoadFile(test.fname)
		count := ss.CountDupes()
		if count != test.count {
			t.Errorf("[Test %v] - failed for file=%v. got=%v, exp=%v", i, test.fname, count, test.count)
		}
	}
}

func TestSimilar(t *testing.T) {
	tests := []struct {
		fname string
		count int
	}{
		{"sentence_files/trivial.txt", 1},
		{"sentence_files/tiny.txt", 0},
		{"sentence_files/small.txt", 0},
	}

	for i, test := range tests {
		ss := sentences.New(int(math.Pow(2, 8)), hash.Campbell3)
		ss.LoadFile(test.fname)
		count := ss.CountSimilar()
		if count != test.count {
			t.Errorf("[Test %v] - failed for file=%v. got=%v, exp=%v", i, test.fname, count, test.count)
		}
	}
}
