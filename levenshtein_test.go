package levenshtein

import (
	"fmt"
	"testing"
)

func TestSparseAutomaton(t *testing.T) {

	words := []string{"banana", "bananas", "cabana"}
	//, "foobarbazfoobarbaz", "a", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", ""

	for n := 0; n < 3; n++ {

		for _, word := range words {
			a := NewSparseAutomaton(word, n)

			for _, query := range words {

				state := a.Start()

				for i, b := range query {

					state = a.Step(state, byte(b))
					fmt.Println(word, n, "=>", query[:i+1], a.IsMatch(state), a.CanMatch(state))
					//assert dense.is_match(s_dense) == sparse.is_match(s_sparse)
					//assert dense.can_match(s_dense) == sparse.can_match(s_sparse)
				}
			}
		}
	}

}
