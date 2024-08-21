package set

// Set holds a set of elements.
type Set[E comparable] struct {
	m map[E]struct{}
}

// New returns a new [Set].
func New[E comparable]() *Set[E] {
	return &Set[E]{m: make(map[E]struct{})}
}

// Add adds an element to a set.
func (s *Set[E]) Add(v E) {
	s.m[v] = struct{}{}
}

// Contains reports whether an element is in a set.
func (s *Set[E]) Contains(v E) bool {
	_, ok := s.m[v]
	return ok
}
