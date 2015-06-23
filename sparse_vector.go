package levenshtein

type entry struct {
	idx, val int
}

// sparseVector is a crude implementation of a sparse vector for our needs
type sparseVector []*entry

// newSparseVector creates a new sparse vector with the initial values of the dense int slice given to it
func newSparseVector(values []int) sparseVector {
	v := make(sparseVector, len(values))

	for i := 0; i < len(values); i++ {
		v[i] = &entry{i, values[i]}
	}

	return v
}

// append appends another sparse vector entry with the given index and value. NOTE: We do not check
// that an entry with the same index is present in the vector
func (v sparseVector) append(index, value int) sparseVector {
	return append(v, &entry{idx: index, val: value})
}
