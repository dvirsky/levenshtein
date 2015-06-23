package levenshtein

type node struct {
	b        byte
	children []*node
	terminal bool
}

func newNode(c byte) *node {
	return &node{
		b: c,
	}
}

func (n *node) child(c byte) *node {

	if n.children == nil {
		return nil
	}
	for _, child := range n.children {
		if child.b == c {
			return child
		}
	}
	return nil
}

func (n *node) addChild(c byte) *node {

	child := newNode(c)
	if n.children == nil {
		n.children = []*node{child}
	} else {
		n.children = append(n.children, child)
	}
	return child
}

///insert a new record into the index
func (n *node) add(key string) {

	current := n
	if n.b == key[0] {
		key = key[1:]
	}

	//find or create the node to put this record on
	for pos := 0; pos < len(key); pos++ {

		next := current.child(key[pos])

		//we're iterating an existing node here
		if next != nil {
			current = next
		} else { //nothing for this prefix - create a new node
			current = current.addChild(key[pos])
		}

		if pos == len(key)-1 {
			current.terminal = true
		}
	}

}

type stackNode struct {
	vec  sparseVector
	str  string
	node *node
}

func (n *node) traverse(a *SparseAutomaton, vec sparseVector) []string {

	ret := []string{}

	stack := make([]*stackNode, 1, 20)
	stack[0] = &stackNode{vec, "", n}

	var top *stackNode
	newVec := vec
	for len(stack) > 0 {

		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		n = top.node
		if n.b != 0 {
			newVec = a.Step(top.vec, n.b)
		}
		// if this is a terminal node - just check if we have a match and add it to the results
		if n.terminal && len(newVec) > 0 && a.IsMatch(newVec) {
			ret = append(ret, top.str+string(n.b))
		}

		if n.children != nil && a.CanMatch(newVec) {

			if n.b != 0 {
				top.str += string(n.b)
			}

			for _, child := range n.children {

				stack = append(stack, &stackNode{newVec, top.str, child})
			}
		}

	}

	return ret

}

// Trie holds a trie representation of a dictionary of words, for fuzzy matching against it
type Trie struct {
	root *node
}

// NewTrie creates a new empty trie
func NewTrie() *Trie {
	return &Trie{
		root: newNode(0),
	}
}

// Insert adds a string to the trie
func (t *Trie) Insert(s string) {
	t.root.add(s)
}

// Exists returns true if a string exists as it is in the trie
func (t *Trie) Exists(s string) bool {

	current := t.root
	for i := 0; i < len(s); i++ {

		child := current.child(s[i])
		if child == nil {
			return false
		}
		current = child

	}

	return true

}

// FuzzyMatches returns all the words in the trie that are with maxDist edit distance from s
func (t *Trie) FuzzyMatches(s string, maxDist int) []string {

	a := NewSparseAutomaton(s, maxDist)

	state := a.Start()
	return t.root.traverse(a, state)

	//	for _, child := range t.root.children {

	//		matches := child.traverse(a, state)
	//		if len(matches) > 0 {
	//			ret = append(ret, matches...)
	//		}

	//	}
	//	return ret
}
