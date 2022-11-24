package ngram

import (
	"fmt"
	"testing"
)

func TestNgrams(t *testing.T) {
	ngrams := New(3, "asdfsdf")
	// asd, sdf, dfs, fsd, sdf
	if most := ngrams.NthFrequentGram(0); most != "sdf" {
		t.Errorf("Most Frequent 3-gram is not 'sdf'. got=%v\n", most)
	}

	rareTests := []struct {
		i    int
		gram string
	}{
		{0, "fsd"},
		{1, "dfs"},
		{2, "asd"},
	}

	for i, test := range rareTests {
		if rare := ngrams.NthRareGram(test.i); rare != test.gram {
			t.Errorf("[Test %v] - Wrong ngram. exp=%v got=%v\n", i, test.gram, rare)
		}
	}

	if ngrams.String() != "fsddfsasdsdf" {
		t.Errorf("Error in calculating ngram hash string. got=%v exp=%v\n", ngrams.String(), "fsddfsasdsdf")
	}

	if fmt.Sprintf("%v", ngrams.Bytes()) != "[102 115 100 100 102 115 97 115 100 115 100 102 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]" {
		t.Errorf("Wrong byte slice calculated for ngram")
	}
}

func TestNgramEditDistance(t *testing.T) {
	grams := []struct {
		g1   string
		g2   string
		dist int
	}{
		{"asdf", "asd", 1},
		{"asdf", "asfg", 2},
		{"abc", "xyz", 6},
		{"abcd", "cdba", 4},
	}

	for i, gram := range grams {
		calculatedDist := ngramEditDistance(gram.g1, gram.g2)
		if calculatedDist != gram.dist {
			t.Errorf("[Test %v] - ngramDistance failed. got=%v exp=%v\n", i, calculatedDist, gram.dist)
		}
	}
}

func TestNgramDistance(t *testing.T) {
	grams := New(3, "abcdefghijkl")
	tests := []struct {
		g1   string
		g2   string
		dist int
	}{
		{"abc", "efg", 1},
		{"abc", "bcd", 0},
		{"abc", "def", 0},
		{"abc", "jkl", 6},
		{"abc", "xyz", -1},
	}

	for i, test := range tests {
		if val := grams.ngramDistance(test.g1, test.g2); val != test.dist {
			t.Errorf("[Test %v] - ngram distance between grams was wrong. got=%v exp=%v", i, val, test.dist)
		}
		if val := grams.ngramDistance(test.g2, test.g1); val != test.dist {
			t.Errorf("[Test %v'] - ngram distance between grams was wrong. got=%v exp=%v", i, val, test.dist)
		}
	}
}
