package levenshtein

import (
	"fmt"
	"testing"
)

func TestSparseAutomaton(t *testing.T) {

	// The test doesn't test much, it just prints the results

	words := []string{"banana", "bananas", "cabana"}
	//, "foobarbazfoobarbaz", "a", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", ""

	for n := 0; n < 3; n++ {

		for _, word := range words {
			a := NewSparseAutomaton(word, n)

			for _, query := range words {

				fmt.Printf("Testing %s <==> %s, max distance %d\n\n", query, word, n)

				state := a.Start()
				for i, b := range query {

					state = a.Step(state, byte(b))
					canMatch, isMatch := a.CanMatch(state), a.IsMatch(state)

					fmt.Printf("Query: %s, Match? %v, CanMatch? %v\n", query[:i+1], isMatch, canMatch)

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
