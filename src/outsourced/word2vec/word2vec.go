package word2vec

import (
	"errors"
	"topic-tree/src/outsourced/word2vec/adapters"
	// "code.sajari.com/word2vec"
)

const (
	// AdapterImplementationWego chooses adapter wrapping
	// github.com/ynqa/wego.
	AdapterImplementationWego = iota
)

// W2VAdapter serves as a word2vec adapter for this
// project (topic-tree). Only API is SearchWord.
type W2VAdapter interface {
	// Searches a model by query and limits results to k.
	// Returned matches and scores are associated by index.
	SearchWord(query string, k int) (
		matches []string, scores []float64, err error)
}

// New accepts a path to a word2vec model and an int for implementation
// choice. Returns W2VAdapter which can be used for w2v queries.
// Note: Implementation choice constants are defined in this pkg,
// starting with the prefix 'AdapterImplementation'
func New(path string, implChoice int) (adapter W2VAdapter, err error) {

	switch implChoice {
	case AdapterImplementationWego:
		choice := adapters.Wego{}
		err = choice.Load(path)
		adapter = &choice
	default:
		err = errors.New("invalid implChoice")
	}
	return
}
