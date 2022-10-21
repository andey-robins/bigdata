package sentences

import (
	"bufio"
	"crypto/sha256"
	"log"
	"math"
	"os"

	"github.com/andey-robins/bigdata/similarity/hashtable"
)

type SentenceSimilarity struct {
	Duplicates int
	HashTable  *hashtable.Hashtable
}

func New() *SentenceSimilarity {
	ht := hashtable.New(int(math.Pow(float64(2), float64(16))), sha256.Sum256)
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
			oldVal, err := ss.HashTable.Get(line)
			if err != nil {
				// TODO recovery
				panic(err)
			}
			ss.HashTable.Update(line, oldVal+1)
		} else {
			err := ss.HashTable.Insert(line, 1)
			if err != nil {
				// TODO recovery
				panic(err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %v\n", err)
	}
}

func (ss *SentenceSimilarity) CountDupes() int {
	dupes := 0

	for _, e := range ss.HashTable.Keys() {

		count, err := ss.HashTable.Get(e)
		if err != nil {
			// TODO recovery
			panic(err)
		}

		if count != 1 {
			dupes++
		}
	}

	return dupes
}
