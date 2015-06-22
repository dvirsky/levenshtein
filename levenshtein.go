package levenshtein

type SparseAutomaton struct {
	str string
	max int
}

type sparseVector struct {
	values  []int
	indices []int
}

func newSparseVector(values ...int) *sparseVector {
	v := new(sparseVector)
	v.values = values
	v.indices = make([]int, len(values))
	for i := 0; i < len(v.values); i++ {
		v.indices[i] = i
	}

	return v
}

func (v *sparseVector) set(index, value int) {
	v.values, v.indices = append(v.values, value), append(v.indices, index)
}

func (v *sparseVector) len() int {
	return len(v.indices)
}

func NewSparseAutomaton(s string, maxEdits int) *SparseAutomaton {
	return &SparseAutomaton{
		str: s,
		max: maxEdits,
	}

}

func (a *SparseAutomaton) Start() *sparseVector {

	vals := make([]int, len(a.str)+1)
	for i := 0; i < a.max+1; i++ {
		vals[i] = i
	}

	return newSparseVector(vals...)

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (a *SparseAutomaton) Step(state *sparseVector, c byte) *sparseVector {

	var newVec *sparseVector
	if state.len() > 0 && state.indices[0] == 0 && state.values[0] < a.max {
		newVec = newSparseVector(state.values[0] + 1)
	} else {
		newVec = newSparseVector()
	}

	for j, i := range state.indices {

		if i == len(a.str) {
			break
		}

		cost := 1
		if a.str[i] == c {
			cost = 0
		}

		val := state.values[j] + cost
		if newVec.len() > 0 && newVec.indices[len(newVec.indices)-1] == i {
			val = min(val, newVec.values[len(newVec.values)-1]+1)
		}

		if j+1 < len(newVec.indices) && newVec.indices[j+1] == i+1 {
			val = min(val, newVec.values[j+1]+1)
		}
		if val <= a.max {
			newVec.set(i+1, val)
		}

	}
	return newVec
}

func (a *SparseAutomaton) IsMatch(v *sparseVector) bool {
	return v.len() > 0 && v.indices[len(v.indices)-1] == len(a.str)
}

func (a *SparseAutomaton) CanMatch(v *sparseVector) bool {
	return v.len() > 0
}

func (a *SparseAutomaton) Transitions(indices, values []int) []byte {

	set := map[byte]struct{}{}
	for _, i := range indices {
		if i < len(a.str) {
			set[a.str[i]] = struct{}{}
		}
	}

	ret := make([]byte, 0, len(set))
	for b, _ := range set {
		ret = append(ret, b)
	}

	return ret
}

//        return set(self.string[i] for i in indices if i < len(self.string))

//    def can_match(self, (indices, values)):
//        return bool(indices)

//    def step(self, (indices, values), c):
//        if indices and indices[0] == 0 and values[0] < self.max_edits:
//            new_indices = [0]
//            new_values = [values[0] + 1]
//        else:
//            new_indices = []
//            new_values = []

//        for j,i in enumerate(indices):
//            if i == len(self.string): break
//            cost = 0 if self.string[i] == c else 1
//            val = values[j] + cost
//            if new_indices and new_indices[-1] == i:
//                val = min(val, new_values[-1] + 1)
//            if j+1 < len(indices) and indices[j+1] == i+1:
//                val = min(val, values[j+1] + 1)
//            if val <= self.max_edits:
//                new_indices.append(i+1)
//                new_values.append(val)

//        return (new_indices, new_values)

//    def is_match(self, (indices, values)):
//        return bool(indices) and indices[-1] == len(self.string)

//    def can_match(self, (indices, values)):
//        return bool(indices)

//    def transitions(self, (indices, values)):
//        return set(self.string[i] for i in indices if i < len(self.string))
