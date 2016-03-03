package levenshtein

// SparseAutomatonRune is almost identical to SparseAutomaton except
// that it operates on runes instead of bytes
type SparseAutomatonRune struct {
	SparseAutomaton
	runes []rune
}

// NewSparseAutomatonRune creates a new automaton for the string s,
// with a given max edit distance check
func NewSparseAutomatonRune(s string, maxEdits int) *SparseAutomatonRune {
	return &SparseAutomatonRune{
		SparseAutomaton{max: maxEdits},
		[]rune(s),
	}
}

func (a *SparseAutomatonRune) Transitions(v sparseVector) []rune {
	set := map[rune]struct{}{}

	for _, entry := range v {
		if entry.idx < len(a.runes) {
			set[a.runes[entry.idx]] = struct{}{}
		}
	}

	ret := make([]rune, 0, len(set))
	for r, _ := range set {
		ret = append(ret, r)
	}

	return ret
}

// StepRune returns the next state of the automaton given a previous
// state and a rune to check
func (a *SparseAutomatonRune) Step(state sparseVector, r rune) sparseVector {
	newVec := make(sparseVector, 0)

	if len(state) > 0 && state[0].idx == 0 && state[0].val < a.max {
		newVec = newVec.append(0, state[0].val+1)
	}

	for j, entry := range state {
		if entry.idx == len(a.runes) {
			break
		}

		cost := 0
		if a.runes[entry.idx] != r {
			cost = 1
		}

		val := state[j].val + cost
		if len(newVec) != 0 && newVec[len(newVec)-1].idx == entry.idx {
			val = min(val, newVec[len(newVec)-1].val+1)
		}

		if len(state) > j+1 && state[j+1].idx == entry.idx+1 {
			val = min(val, state[j+1].val+1)
		}

		if val <= a.max {
			newVec = newVec.append(entry.idx+1, val)
		}
	}

	return newVec
}

// IsMatch returns true if the current state vector represents a string
// of runes that is within the max edit distance from the initial
// automaton string
func (a *SparseAutomatonRune) IsMatch(v sparseVector) bool {
	return len(v) != 0 && v[len(v)-1].idx == len(a.runes)
}
