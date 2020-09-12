package adapters

import (
	"os"

	"github.com/ynqa/wego/pkg/search"
)

// wego implements word2vec.W2VAdapter
// var _ word2vec.W2VAdapter = (*wego)(nil)

type Wego struct {
	searcher *search.Searcher
}

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

func (w *Wego) SearchWord(query string, k int) (matches []string, scores []float64, err error) {
	neighbors, err := w.searcher.InternalSearch(query, k)
	for _, neighbor := range neighbors {
		matches = append(matches, neighbor.Word)
		scores = append(scores, neighbor.Similarity)
	}
	return
}
