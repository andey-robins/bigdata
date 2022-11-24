package ngram

import (
	"math"
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

// NSpacedRareGrams will return the i rarest ngrams that are at least spacing
// edit distance apart from each other
func (n *Ngram) NSpacedRareGrams(spacing, i int) []string {
	grams := make([]string, 0)

	idx := 0
	// iterate through our grams until we fill a list with sufficiently spaced ngrams
	for len(grams) < i {
		nextGram := n.NthRareGram(idx)
		goodGram := true

		for _, gram := range grams {
			if n.ngramDistance(nextGram, gram) < spacing {
				goodGram = false
			}
		}

		if goodGram {
			grams = append(grams, nextGram)
		}

		idx++
	}

	return grams
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

func (n *Ngram) Sentence() string {
	return n.sentence
}

// adapted for words and sentences from the code here:
// https://golangbyexample.com/edit-distance-two-strings-golang/
func ngramEditDistance(g1, g2 string) int {
	lenGram1 := len(g1)
	lenGram2 := len(g2)

	editDistanceMatrix := make([][]int, lenGram1+1)

	for i := range editDistanceMatrix {
		editDistanceMatrix[i] = make([]int, lenGram2+1)
	}

	for i := 1; i <= lenGram2; i++ {
		editDistanceMatrix[0][i] = i
	}

	for i := 1; i <= lenGram1; i++ {
		editDistanceMatrix[i][0] = i
	}
	for i := 1; i <= lenGram1; i++ {
		for j := 1; j <= lenGram2; j++ {

			if g1[i-1] == g2[j-1] {
				editDistanceMatrix[i][j] = editDistanceMatrix[i-1][j-1]
			} else {
				// modified here to not count a replace as a distance of 1
				editDistanceMatrix[i][j] = 1 + int(math.Min(float64(editDistanceMatrix[i-1][j]), float64(editDistanceMatrix[i][j-1])))
			}
		}
	}
	return editDistanceMatrix[lenGram1][lenGram2]
}

// ngramDistance will calculate the distance between two ngrams
// returns -1 if one of the ngrams isn't found (we never expect
// this to happen with current use cases, so something must have gone wrong)
// returns 0 if there is an overlap
func (n *Ngram) ngramDistance(g1, g2 string) int {
	g1Idx := strings.Index(n.sentence, g1)
	g2Idx := strings.Index(n.sentence, g2)

	if g1Idx == -1 || g2Idx == -1 {
		return -1
	}

	dist := g1Idx - g2Idx
	if int(math.Abs(float64(dist))) < len(g1) {
		return 0
	}

	return int(math.Abs(float64(dist))) - len(g1)
}
