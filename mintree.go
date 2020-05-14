package levenshtein

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/rubenv/mafsa"
)

type MinTree struct {
	mafsa.MinTree
	root *MinTreeNode
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

// Creates a new MinTree from a sorted list of strings. The list must be
// sorted because that is what the mafsa package expects.
func NewMinTree(words []string) (*MinTree, error) {
	bt := mafsa.New()
	for _, w := range words {
		err := bt.Insert(w)
		if err != nil {
			return nil, fmt.Errorf("Could not insert value into MinTree: %q\n", err)
		}
	}

	bt.Finish()
	me := mafsa.Encoder{}
	bytes, err := me.Encode(bt)
	if err != nil {
		return nil, err
	}

	de := mafsa.Decoder{}
	mmt, err := de.Decode(bytes)
	if err != nil {
		return nil, err
	}

	return &MinTree{*mmt, &MinTreeNode{*mmt.Root, rune(0)}}, nil
}

// Creates a new MinTree from a sorted list of strings. The list must be
// sorted because that is what the mafsa package expects.
// After the MinTree has been successfully created, the function also
// writes it to the io.Writer.
func NewMinTreeWrite(words []string, wr io.Writer) (*MinTree, error) {
	bt := mafsa.New()
	for _, w := range words {
		bt.Insert(w)
	}

	bt.Finish()
	me := mafsa.Encoder{}
	bs, err := me.Encode(bt)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(wr, bytes.NewReader(bs))
	if err != nil {
		return nil, err
	}

	de := mafsa.Decoder{}
	mmt, err := de.Decode(bs)
	if err != nil {
		return nil, err
	}

	return &MinTree{*mmt, &MinTreeNode{*mmt.Root, rune(0)}}, nil
}

// LoadMinTree loads a MinTree from an io.Reader.
func LoadMinTree(r io.Reader) (*MinTree, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	mt, err := new(mafsa.Decoder).Decode(data)
	if err != nil {
		return nil, err
	}

	return &MinTree{*mt, &MinTreeNode{*mt.Root, rune(0)}}, nil
}

func (n *MinTreeNode) traverse(a *SparseAutomatonRune, vec sparseVector) []string {
	ret := []string{}

	stack := make([]*mtstackNode, len(n.Edges))
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

// FuzzyMatches returns all the words in the MinTree that are with
// maxDist edit distance from s
func (mt *MinTree) FuzzyMatches(s string, maxDist int) []string {
	a := NewSparseAutomatonRune(s, maxDist)

	state := a.Start()
	return mt.root.traverse(a, state)
}
