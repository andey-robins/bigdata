package hash

import (
	"crypto/sha256"

	"github.com/andey-robins/bigdata/similarity/ngram"
)

func Campbell5(ngrams [3]ngram.Ngram, sentence string) [32]byte {
	hash := ""
	for _, e := range ngrams {
		hash += e.String()
	}

	hash += string(sha256.New().Sum([]byte(sentence)))
	return sha256.Sum256([]byte(hash))
}
