package hash

import (
	"crypto/sha256"
	"fmt"

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
	fmt.Println(gram.Campbell4Hash())
	return gram.Campbell4Hash()
}

func Campbell5(gram *ngram.Ngram) [32]byte {
	// fmt.Println(gram.Campbell5Hash())
	return gram.Campbell5Hash()
}
