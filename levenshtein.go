package levenshtein

// SparseAutomaton is a naive Go implementation of a levenshtein automaton using sparse vectors, as described
// and implemented here: http://julesjacobs.github.io/2015/06/17/disqus-levenshtein-simple-and-fast.html
type SparseAutomaton struct {
	str string
	max int
}

// NewSparseAutomaton creates a new automaton for the string s, with a given max edit distance check
func NewSparseAutomaton(s string, maxEdits int) *SparseAutomaton {
	return &SparseAutomaton{
		str: s,
		max: maxEdits,
	}

}

// Start initializes the automaton's state vector and returns it for further iteration
func (a *SparseAutomaton) Start() sparseVector {

	vals := make([]int, a.max+1)
	for i := 0; i < a.max+1; i++ {
		vals[i] = i
	}

	return newSparseVector(vals)

}

// just a utility min function for ints, CUZ GO AINT GOT NO GENERICS
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Step returns the next state of the automaton given a pervios state and a character to check
func (a *SparseAutomaton) Step(state sparseVector, c byte) sparseVector {

	var newVec sparseVector

	if len(state) > 0 && state[0].idx == 0 && state[0].val < a.max {
		newVec = newSparseVector([]int{state[0].val + 1})
	} else {
		newVec = sparseVector{}
	}

	for j, entry := range state {

		if entry.idx == len(a.str) {
			break
		}

		cost := 0
		if a.str[entry.idx] != c {
			cost = 1
		}

		val := state[j].val + cost

		if len(newVec) > 0 && newVec[len(newVec)-1].idx == entry.idx {
			val = min(val, newVec[len(newVec)-1].val+1)
		}

		if j+1 < len(state) && state[j+1].idx == entry.idx+1 {
			val = min(val, state[j+1].val+1)
		}

		if val <= a.max {
			newVec = newVec.append(entry.idx+1, val)
		}

	}
	return newVec
}

// IsMatch returns true if the current state vector represents a string that is within the max
// edit distance from the initial automaton string
func (a *SparseAutomaton) IsMatch(v sparseVector) bool {
	return len(v) > 0 && v[len(v)-1].idx == len(a.str)
}

// CanMatch returns true if there is a possibility that feeding the automaton with more steps will
// yield a match. Once CanMatch is false there is no point in continuing iteration
func (a *SparseAutomaton) CanMatch(v sparseVector) bool {
	return len(v) > 0
}

func (a *SparseAutomaton) Transitions(v sparseVector) []byte {

	set := map[byte]struct{}{}
	for _, entry := range v {

		if entry.idx < len(a.str) {
			set[a.str[entry.idx]] = struct{}{}
		}
	}

	ret := make([]byte, 0, len(set))
	for b, _ := range set {
		ret = append(ret, b)
	}

	return ret
}
