package sentences

import (
	"bufio"
	"log"
	"os"

	"github.com/andey-robins/bigdata/similarity/hashtable"
	"github.com/andey-robins/bigdata/similarity/ngram"
)

// SentenceSimilarity is a wrapper object that is used to expose a sentence
// database for counting similar words with specific edit distances
// Duplicates is an exported int which is the number of edit distance 0 words in the input
// HashTable is the mapping of words to times they appear
type SentenceSimilarity struct {
	Duplicates int
	HashTable  *hashtable.Hashtable
}

// New creates a new sentence similarity object with a hashtable of size `size` and
// using a hashing algorithm of `hash`
func New(size int, hash func(gram *ngram.Ngram) [32]byte) *SentenceSimilarity {
	ht := hashtable.New(size, hash)
	return &SentenceSimilarity{
		Duplicates: 0,
		HashTable:  ht,
	}
}

// LoadFile takes in a filename as the argument fname and propogates the SentenceSimilarity
// datastructure with the unique sentenes. Also counts duplicates along the way in
// in linear time
func (ss *SentenceSimilarity) LoadFile(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		// TODO recovery
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if ss.HashTable.Exists(line) {
			val, err := ss.HashTable.Get(line)

			if err != nil {
				panic(err)
			}

			if val == 1 {
				ss.Duplicates++
			}

			ss.HashTable.Update(line, 2)
		} else {
			err := ss.HashTable.Insert(line, 1)
			if err != nil {
				// TODO recovery
				panic(err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning file: %v\n", err)
	}
}

// CountDupes returns the number of perfect duplicates (i.e. edit distance of 0)
// present within the hashtable
func (ss *SentenceSimilarity) CountDupes() int {
	log.Printf("collisions=%v", ss.HashTable.Collisions())
	return ss.Duplicates
}

// CountSimilar determines if a sentence is similar if any one deletion
// or one addition of a word creates a duplicate (i.e. edit distance of 1)
// and returns the count of the number of sentences within edit distance 1
// of another
func (ss *SentenceSimilarity) CountSimilar() int {
	count := 0
	for _, row := range ss.HashTable.GetSimilarSentences() {
		for i, s := range row {
			for j := i + 1; j < len(row); j++ {
				if isDistanceOne(s, row[j]) {
					log.Printf("%v ~ %v\n", s, row[j])
					count++
				}
			}
		}
	}
	log.Printf("collisions=%v", ss.HashTable.Collisions())
	return count
}
