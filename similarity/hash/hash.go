package hash

import (
	"crypto/sha256"

	"github.com/andey-robins/bigdata/similarity/ngram"
)

func Sha256Wrapper(gram *ngram.Ngram) [32]byte {
	bytes := []byte(gram.Sentence())
	return sha256.Sum256(bytes[:])
}

func Campbell3(gram *ngram.Ngram) [32]byte {
	return gram.Bytes()
}
