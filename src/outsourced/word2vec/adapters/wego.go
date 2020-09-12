package adapters

import (
	"os"

	"github.com/ynqa/wego/pkg/search"
)

// commented because circular ref
// var _ word2vec.W2VAdapter = (*wego)(nil)

// Wego implements word2vec.W2VAdapter. Functions mainly
// as a wrapper for github.com/ynqa/wego/pkg/search, which
// loads and searches models (Load() & SearchWord())
type Wego struct {
	searcher *search.Searcher
}

// Load attempts to lead a word2vec model from path.
func (w *Wego) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	searcher, err := search.NewForVectorFile(f)
	if err != nil {
		return err
	}
	w.searcher = searcher
	return nil
}

// SearchWord is the main API for struct, it wraps searching functionality of
// github.com/ynqa/wego/pkg/search by specifying query and result limit (k arg).
// Returned matches and scores are associated by index.
func (w *Wego) SearchWord(query string, k int) (matches []string, scores []float64, err error) {
	neighbors, err := w.searcher.InternalSearch(query, k)
	for _, neighbor := range neighbors {
		matches = append(matches, neighbor.Word)
		scores = append(scores, neighbor.Similarity)
	}
	return
}
