package word2vec

import (
	"errors"
	"topic-tree/src/outsourced/word2vec/adapters"
	// "code.sajari.com/word2vec"
)

const (
	AdapterImplementationWego = iota
)

type W2VAdapter interface {
	SearchWord(query string, k int) (
		matches []string, scores []float64, err error)
}

func New(path string, implChoice int) (W2VAdapter, error) {

	var (
		adapter W2VAdapter
		err     error
	)

	switch implChoice {
	case AdapterImplementationWego:
		choice := adapters.Wego{}
		err = choice.Load(path)
		adapter = &choice
	default:
		err = errors.New("invalid implChoice")
	}
	return adapter, err
}
