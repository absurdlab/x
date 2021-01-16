package stringz

import "encoding/json"

// NewSet returns a new empty Set.
func NewSet() *Set {
	return &Set{
		list:   []string{},
		lookup: map[string]struct{}{},
	}
}

// NewSetOf is a convenient wrapper for NewSet and Add
func NewSetOf(values ...string) *Set {
	return NewSet().Add(values...)
}

// Set is a set data structure that guarantees stable iteration in the
// order of insertion.
//
// Insertion by Add is O(1). Deletion by Remove is O(N). Element lookup by Contains is O(1).
// Equality comparison by Equals is O(N).
type Set struct {
	list   []string
	lookup map[string]struct{}
}

// Len returns the length of the Set. If set is nil, returns 0.
func (s *Set) Len() int {
	if s == nil {
		return 0
	}

	return len(s.list)
}

// Add adds all values into the Set, and returns the Set for fluent access.
func (s *Set) Add(values ...string) *Set {
	for _, it := range values {
		s.addOne(it)
	}
	return s
}

func (s *Set) addOne(value string) {
	if s.Contains(value) {
		return
	}
	s.lookup[value] = struct{}{}
	s.list = append(s.list, value)
}

// Remove removes all values from the Set, and returns the Set for fluent access.
func (s *Set) Remove(values ...string) *Set {
	for _, it := range values {
		s.removeOne(it)
	}
	return s
}

func (s *Set) removeOne(value string) {
	if !s.Contains(value) {
		return
	}

	for i, it := range s.list {
		if it == value {
			s.list = append(s.list[:i], s.list[i+1:]...)
			break
		}
	}

	delete(s.lookup, value)
}

// Contains returns true if the value is in the Set.
func (s *Set) Contains(value string) bool {
	_, ok := s.lookup[value]
	return ok
}

// Equals returns true if this Set contains the same set of elements from
// the provided Set. Any nil set break equality.
func (s *Set) Equals(t *Set) bool {
	if s == nil || t == nil || s.Len() != t.Len() {
		return false
	}

	for each := range s.lookup {
		if !t.Contains(each) {
			return false
		}
	}

	return true
}

// Array exports a copy of the elements in the set to a string slice.
func (s *Set) Array() []string {
	if s == nil {
		return []string{}
	}

	dup := make([]string, len(s.list))
	copy(dup, s.list)
	return dup
}

// ForEach iterates elements in the Set in place.
func (s *Set) ForEach(elemFunc func(index int, it string)) {
	for i, elem := range s.list {
		elemFunc(i, elem)
	}
}

// All checks if all elements in the Set meets the criteria.
func (s *Set) All(criteria func(it string) bool) bool {
	for _, elem := range s.list {
		if !criteria(elem) {
			return false
		}
	}
	return true
}

// Any checks if any elements in the Set meets the criteria.
func (s *Set) Any(criteria func(it string) bool) bool {
	for _, elem := range s.list {
		if criteria(elem) {
			return true
		}
	}
	return false
}

func (s *Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.list)
}

func (s *Set) UnmarshalJSON(bytes []byte) error {
	var list []string
	if err := json.Unmarshal(bytes, &list); err != nil {
		return err
	}

	set := NewSet().Add(list...)
	*s = *set

	return nil
}
