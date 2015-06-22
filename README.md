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
