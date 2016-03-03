package levenshtein

import (
	"fmt"
	"sort"
	"testing"
)

func TestMinTreeFuzzySearch(t *testing.T) {
	var err error
	uniq := make(map[string]struct{}, len(testwords)/2)

	for _, w := range testwords {
		uniq[w] = struct{}{}
	}
	uniqwords := make([]string, 0, len(uniq))
	for k, _ := range uniq {
		uniqwords = append(uniqwords, k)
	}
	fmt.Printf("Number of unique words in the MinTree: %d\n", len(uniqwords))

	sort.Strings(uniqwords)
	mt, err = NewMinTree(uniqwords)
	if err != nil {
		t.Fatalf("Could not create MinTree: %q. Exiting.", err)
	}

	matches := mt.FuzzyMatches("danger", 2)
	fmt.Printf("Fuzzymatch count for \"danger\" with distance two: %d\n", len(matches))
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
