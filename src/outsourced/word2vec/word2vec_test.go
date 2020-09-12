package word2vec

import (
	"testing"
)

var (
	// # 171 MB.
	modelPathA = "../../../models/glove/glove.6B.50d.txt"
	// # 347 MB.
	modelPathB = "../../../models/glove/glove.6B.300d.txt"
)

func sliceContains(search string, pool []string) bool {
	for i := 0; i < len(pool); i++ {
		if search == pool[i] {
			return true
		}
	}
	return false
}

func TestSearch(t *testing.T) {
	// # Search something which will have a predictable result.
	// # .. cats is often associated with dogs, so..
	var (
		search        = "cat"
		probableMatch = "dog"
	)
	// # Load model.
	s, err := New(modelPathA, AdapterImplementationWego)
	if err != nil {
		t.Error(err)
	}
	// # Lookup.
	matches, _, err := s.SearchWord(search, 5)
	if err != nil {
		t.Error(err)
	}
	// # Check result.
	if !sliceContains(probableMatch, matches) {
		t.Errorf("Prabable Model issue. \n\tSearched: %v, \n\tFound:%v",
			search, matches)
	}
}
