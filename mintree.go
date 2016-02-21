package levenshtein

import (
	"github.com/smartystreets/mafsa"
)

type MinTree struct {
	Root MinTreeNode
}

type MinTreeNode struct {
	mafsa.MinTreeNode
	r rune
}

type mtstackNode struct {
	vec  sparseVector
	str  string
	node *MinTreeNode
}

func (n *MinTreeNode) traverse(a *SparseAutomaton, vec sparseVector) []string {
	ret := []string{}

	stack := make([]*mtstackNode, 0, len(n.Edges))
	var i int

	for r, mt := range n.Edges {
		nmt := MinTreeNode{*mt, r}
		stack[i] = &mtstackNode{vec, "", &nmt}
		i++
	}

	var top *mtstackNode
	newVec := vec

	for len(stack) > 0 {
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		n = top.node
		if n.r != 0 {
			newVec = a.Step(top.vec, n.r)
		}

		// if this is a terminal node - just check if we have
		// a match and add it to the results
		if n.Final && len(newVec) > 0 && a.IsMatch(newVec) {
			ret = append(ret, top.str+string(n.r))
		}

		if n.Edges != nil && a.CanMatch(newVec) {
			if n.r != 0 {
				top.str += string(n.r)
			}

			for r, child := range n.Edges {
				stack = append(stack, &mtstackNode{newVec, top.str, &MinTreeNode{*child, r}})
			}
		}
	}

	return ret
}

// FuzzyMatches returns all the words in the MT that are with maxDist
// edit distance from s
func (mt *MinTree) FuzzyMatches(s string, maxDist int) []string {
	a := NewSparseAutomaton(s, maxDist)

	state := a.Start()
	return mt.Root.traverse(a, state)
}
