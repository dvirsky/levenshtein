package levenshtein

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSparseAutomaton(t *testing.T) {

	// The test doesn't test much, it just prints the results

	words := []string{"banana", "bananas"}
	//, "foobarbazfoobarbaz", "a", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", ""

	for n := 2; n < 3; n++ {

		for _, word := range words {
			a := NewSparseAutomaton(word, n)

			for _, query := range words {

				fmt.Printf("Testing query %s vs word %s, max distance %d\n\n", query, word, n)

				state := a.Start()
				for i, b := range query {

					state = a.Step(state, byte(b))
					canMatch, isMatch := a.CanMatch(state), a.IsMatch(state)

					fmt.Printf(" Query: %s, Match? %v, CanMatch? %v\n", query[:i+1], isMatch, canMatch)

					if isMatch && !canMatch {
						t.Errorf("IsMatch is true, canMatch must be true too")
					}
					if !canMatch {
						break
					}

				}
				fmt.Println("----\n")
			}
		}
	}

}

func TestTrie(t *testing.T) {

	//	t.SkipNow()

	trie := NewTrie()
	words := []string{"banana", "bananas", "bnaana", "world"}
	nonwords := []string{"sdfsdfsd", "hellos", "jeolls", "ello", "wrlds"}
	for _, word := range words {
		trie.Insert(word)
	}

	for _, word := range words {
		if !trie.Exists(word) {
			t.Error("Not found", word)
		}
	}

	for _, word := range nonwords {
		if trie.Exists(word) {
			t.Error("found", word)
		}
	}

	matchTest := map[string][]string{
		"banana": {"banana", "bananas", "bnaana"},
		"world":  {"world"},
		"wordl":  {"world"},
		"fordl":  {},
		"bnarna": {"banana", "bnaana"},
		"bananr": {"banana", "bananas"},
	}

	for k, expected := range matchTest {
		matches := trie.FuzzyMatches(k, 2)
		fmt.Println(k, matches)
		assert.EqualValues(t, expected, matches)
	}

}

func ExampleTrie() {

	trie := NewTrie()
	trie.Insert("banana")
	trie.Insert("bananas")
	trie.Insert("cabana")
	trie.Insert("cabasa")

	fmt.Println(trie.FuzzyMatches("banana", 2))
	// Output:
	// [banana bananas cabana]
}
