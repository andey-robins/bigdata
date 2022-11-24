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

func Campbell4(gram *ngram.Ngram) [32]byte {
	var byteSlice []byte
	var byteArray [32]byte
	grams := gram.NSpacedRareGrams(3, 8)
	for _, gram := range grams {
		gramBytes := []byte(gram)
		byteSlice = append(byteSlice, gramBytes...)
	}
	copy(byteArray[:], byteSlice)
	return byteArray
}
