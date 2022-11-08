package ngram

import (
	"fmt"
	"sort"
)

// Ngram is a struct for holding an ngram defined by the New() function below
type Ngram struct {
	n        int
	sentence string
	// grams count is a mapping of gram to occurences; counts the frequency of the ngram
	grams_count map[string]int
	// a list of ngrams in the sentence. ordered if grams_ordered is true
	grams []string
	// a flag to identify if the grams slice is ordered
	grams_ordered bool
}

// New creates a new n, ngram over string s and fills in the appropriate data struct
func New(n int, s string) *Ngram {
	grams := make([]string, 0)
	grams_count := make(map[string]int, 0)
	return &Ngram{
		n:             n,
		sentence:      s,
		grams:         grams,
		grams_count:   grams_count,
		grams_ordered: false,
	}
}

// CalculateGrams will count and order the ngrams if they haven't been ordered
// yet, filling in the appropriate fields in the Ngram struct
func (n *Ngram) CalculateGrams() {
	// if we've already calculated Ngrams, return without re-computing
	if n.grams_ordered {
		return
	}

	// Count ngram occurence
	for i := 0; i <= len(n.sentence)-n.n; i++ {
		n.grams_count[n.sentence[i:i+n.n]] += 1
	}

	// get all unique ngrams
	keys := make([]string, 0)
	for key := range n.grams_count {
		keys = append(keys, key)
	}

	// sort based on associated values in grams_count
	sort.SliceStable(keys, func(i, j int) bool {
		if n.grams_count[keys[i]] == n.grams_count[keys[j]] {
			return keys[i] < keys[j]
		}
		return n.grams_count[keys[i]] > n.grams_count[keys[j]]
	})

	// transfer keys back into ngram struct
	n.grams = append(n.grams, keys...)

	// set ordered flag to prevent re-calculation accidentally
	n.grams_ordered = true
}

// NthFrequentGram will return the ith most frequent ngram. For instance if the most
// frequent ngrams were ['asdf', 'sdfg', 'dfgh', ...] then NthFrequentGram(0) == 'asdf'
func (n *Ngram) NthFrequentGram(i int) string {
	if !n.grams_ordered {
		n.CalculateGrams()
	}
	return n.grams[i]
}

// NthFrequentGram will return the ith least frequent ngram. For instance if the most
// frequent ngrams were [..., 'asdf', 'sdfg', 'dfgh'] then NthRareGram(0) == 'dfgh'
func (n *Ngram) NthRareGram(i int) string {
	if !n.grams_ordered {
		n.CalculateGrams()
	}
	return n.grams[len(n.grams)-1-i]
}

// String will get a string representation of the ngram hash
func (n *Ngram) String() string {
	s := ""
	for i := 0; i < 8 && i < len(n.grams); i++ {
		s += n.NthRareGram(i)
	}
	return s
}

// Bytes returns the byte represetation of the ngram hash
func (n *Ngram) Bytes() [32]byte {
	var byteSlice [32]byte
	copy(byteSlice[:], []byte(n.String()))
	fmt.Println(byteSlice)
	return byteSlice
}
