package sentences

import (
	"bufio"
	"log"
	"os"

	"github.com/andey-robins/bigdata/similarity/hashtable"
)

type SentenceSimilarity struct {
	Duplicates int
	HashTable  *hashtable.Hashtable
}

func New(size int, hash func() [32]byte) *SentenceSimilarity {
	ht := hashtable.New(size, hash)
	return &SentenceSimilarity{
		Duplicates: 0,
		HashTable:  ht,
	}
}

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

func (ss *SentenceSimilarity) CountDupes() int {
	return ss.Duplicates
}

// A Sentence is similar if any one deletion or one addition of a word creates a duplicate
func (ss *SentenceSimilarity) CountSimilar() int {
	return 0
}
