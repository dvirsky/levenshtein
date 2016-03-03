# levenshtein
--
    import "github.com/dvirsky/levenshtein"


## Usage

#### type SparseAutomaton

```go
type SparseAutomaton struct {
}
```

SparseAutomaton is a naive Go implementation of a levenshtein automaton using
sparse vectors, as described and implemented here:
http://julesjacobs.github.io/2015/06/17/disqus-levenshtein-simple-and-fast.html

#### func  NewSparseAutomaton

```go
func NewSparseAutomaton(s string, maxEdits int) *SparseAutomaton
```
NewSparseAutomaton creates a new automaton for the string s, with a given max
edit distance check

#### func (*SparseAutomaton) CanMatch

```go
func (a *SparseAutomaton) CanMatch(v sparseVector) bool
```
CanMatch returns true if there is a possibility that feeding the automaton with
more steps will yield a match. Once CanMatch is false there is no point in
continuing iteration

#### func (*SparseAutomaton) IsMatch

```go
func (a *SparseAutomaton) IsMatch(v sparseVector) bool
```
IsMatch returns true if the current state vector represents a string that is
within the max edit distance from the initial automaton string

#### func (*SparseAutomaton) Start

```go
func (a *SparseAutomaton) Start() sparseVector
```
Start initializes the automaton's state vector and returns it for further
iteration

#### func (*SparseAutomaton) Step

```go
func (a *SparseAutomaton) Step(state sparseVector, c byte) sparseVector
```
Step returns the next state of the automaton given a pervios state and a
character to check

#### func (*SparseAutomaton) Transitions

```go
func (a *SparseAutomaton) Transitions(v sparseVector) []byte
```

#### type Trie

```go
type Trie struct {
}
```

Trie holds a trie representation of a dictionary of words, for fuzzy matching
against it

#### func  NewTrie

```go
func NewTrie() *Trie
```
NewTrie creates a new empty trie

#### func (*Trie) Exists

```go
func (t *Trie) Exists(s string) bool
```
Exists returns true if a string exists as it is in the trie

#### func (*Trie) FuzzyMatches

```go
func (t *Trie) FuzzyMatches(s string, maxDist int) []string
```
FuzzyMatches returns all the words in the trie that are with maxDist edit
distance from s

#### func (*Trie) Insert

```go
func (t *Trie) Insert(s string)
```
Insert adds a string to the trie

#### type MinTree

```go
type MinTree struct {
	mafsa.MinTree
	root *MinTreeNode
}
```

#### func NewMinTree([]string) (*MinTree, error)

```go
func NewMinTree([]string) (*MinTree, error)
```
Creates a new MinTree from a sorted list of strings. The list must be
sorted because that is what the mafsa package expects.


#### func NewMinTreeWrite([]string, io.Writer) (*MinTree, error)

```go
func NewMinTreeWrite(words []string, wr io.Writer) (*MinTree, error) {
```
Creates a new MinTree from a sorted list of strings. The list must be
sorted because that is what the mafsa package expects.
After the MinTree has been successfully created, the function also
writes it to the io.Writer.


#### func LoadMinTree(io.Reader) (*MinTree, error)

```go
func LoadMinTree(r io.Reader) (*MinTree, error)
```
LoadMinTree loads a MinTree from an io.Reader.


#### func (mt *MinTree) FuzzyMatches(string, int) []string

```go
func (mt *MinTree) FuzzyMatches(s string, maxDist int) []string
```
FuzzyMatches returns all the words in the MinTree that are with
maxDist edit distance from s
