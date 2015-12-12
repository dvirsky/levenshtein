package levenshtein

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSparseAutomaton(t *testing.T) {

	//	t.SkipNow()
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

	//t.SkipNow()

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

		assert.Equal(t, len(matches), len(expected))
		for _, m := range matches {
			assert.Contains(t, expected, m)
		}

	}

}

var trie *Trie

func BenchmarkTrie(b *testing.B) {

	//	words := []string{"banana", "bananas", "bnaana", "world"}
	//	for _, w := range words {
	//		trie.Insert(w)
	//	}

	for i := 0; i < b.N; i++ {
		trie.FuzzyMatches("holocaust", 2)
	}
}

func TestMain(m *testing.M) {

	trie = NewTrie()
	for _, word := range SampleEnglish() {
		trie.Insert(word)
	}

	rc := m.Run()

	os.Exit(rc)

}

func SampleEnglish() []string {
	var out []string
	file, err := os.Open("./big.txt")
	if err != nil {
		fmt.Println(err)
		return out
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	// Count the words.
	count := 0
	exp, _ := regexp.Compile("[a-zA-Z]+")

	for scanner.Scan() {
		words := exp.FindAll([]byte(scanner.Text()), -1)
		for _, word := range words {
			if len(word) > 1 {
				out = append(out, strings.ToLower(string(word)))
				count++
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println("Read", len(out), "words")
	return out
}

func ExampleTrie() {
	return

	trie := NewTrie()
	trie.Insert("banana")
	trie.Insert("bananas")
	trie.Insert("cabana")
	trie.Insert("cabasa")

	fmt.Println(trie.FuzzyMatches("banana", 2))
	// XOutput:
	// [banana bananas cabana]
}
