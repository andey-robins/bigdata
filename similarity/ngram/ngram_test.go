package ngram

import (
	"fmt"
	"log"
	"testing"
)

func TestNgrams(t *testing.T) {
	ngrams := New(3, "asdfsdf")
	// asd, sdf, dfs, fsd, sdf
	if most := ngrams.NthFrequentGram(0); most != "sdf" {
		log.Println(ngrams.grams_count)
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
