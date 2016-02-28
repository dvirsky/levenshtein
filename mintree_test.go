package levenshtein

import (
	"fmt"
	"sort"
	"testing"
)

func TestMinTreeFuzzySearch(t *testing.T) {
	var err error

	words := SampleEnglish()
	sort.Strings(words)
	mt, err = NewMinTree(words)
	if err != nil {
		t.Fatalf("Could not create MinTree: %q. Exiting.", err)
	}

	matches := mt.FuzzyMatches("danger", 2)
	fmt.Printf("Fuzzymatch count: %d.\n", len(matches))
	for _, match := range matches {
		fmt.Println(match)
	}
}

var mt *MinTree

func BenchmarkMinTree(b *testing.B) {
	if mt == nil {
		return
	}
	for i := 0; i < b.N; i++ {
		mt.FuzzyMatches("holocaust", 2)
	}
}
