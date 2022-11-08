package ngram

import (
	"sort"
	"strings"
)

// Ngram is a struct for holding an ngram defined by the New() function below
type Ngram struct {
	n        int
	sentence string
	// grams count is a mapping of gram to occurences; counts the frequency of the ngram
	gramsCount map[string]int
	// a list of ngrams in the sentence. ordered if grams_ordered is true
	grams []string
	// a flag to identify if the grams slice is ordered
	gramsOrdered bool
}

// New creates a new n, ngram over string s and fills in the appropriate data struct
func New(n int, s string) *Ngram {
	grams := make([]string, 0)
	gramsCount := make(map[string]int, 0)
	return &Ngram{
		n:            n,
		sentence:     strings.ReplaceAll(s, " ", ""),
		grams:        grams,
		gramsCount:   gramsCount,
		gramsOrdered: false,
	}
}

// CalculateGrams will count and order the ngrams if they haven't been ordered
// yet, filling in the appropriate fields in the Ngram struct
func (n *Ngram) CalculateGrams() {
	// if we've already calculated Ngrams, return without re-computing
	if n.gramsOrdered {
		return
	}

	// Count ngram occurence
	for i := 0; i <= len(n.sentence)-n.n; i++ {
		n.gramsCount[n.sentence[i:i+n.n]] += 1
	}

	// get all unique ngrams
	keys := make([]string, 0)
	for key := range n.gramsCount {
		keys = append(keys, key)
	}

	// sort based on associated values in grams_count
	sort.SliceStable(keys, func(i, j int) bool {
		if n.gramsCount[keys[i]] == n.gramsCount[keys[j]] {
			return keys[i] < keys[j]
		}
		return n.gramsCount[keys[i]] > n.gramsCount[keys[j]]
	})

	// transfer keys back into ngram struct
	n.grams = append(n.grams, keys...)

	// set ordered flag to prevent re-calculation accidentally
	n.gramsOrdered = true
}

// NthFrequentGram will return the ith most frequent ngram. For instance if the most
// frequent ngrams were ['asdf', 'sdfg', 'dfgh', ...] then NthFrequentGram(0) == 'asdf'
func (n *Ngram) NthFrequentGram(i int) string {
	// memoize the frequency calculations
	if !n.gramsOrdered {
		n.CalculateGrams()
	}
	return n.grams[i]
}

// NthFrequentGram will return the ith least frequent ngram. For instance if the most
// frequent ngrams were [..., 'asdf', 'sdfg', 'dfgh'] then NthRareGram(0) == 'dfgh'
func (n *Ngram) NthRareGram(i int) string {
	// memoize the frequency calculations
	if !n.gramsOrdered {
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
	return byteSlice
}
