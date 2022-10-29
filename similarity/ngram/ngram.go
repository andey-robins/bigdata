package ngram

type Ngram struct {
	n     int
	words []string
}

func New(n int, w []string) *Ngram {
	return &Ngram{
		n:     n,
		words: w,
	}
}

func (n *Ngram) String() string {
	s := ""
	for _, e := range n.words {
		s += e
	}
	return s
}
