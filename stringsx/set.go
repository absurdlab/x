package stringsx

// Create a new Set that contains the given values.
func NewSet(values ...string) Set {
	s := Set{}
	for _, value := range values {
		s[value] = struct{}{}
	}
	return s
}

// Set is unique a collection of strings. The internal order is not guaranteed.
type Set map[string]struct{}

// Len returns the size of the Set
func (s Set) Len() int {
	return len(s)
}

// First returns the first item in a non-empty Set. If invoked on an empty Set, the
// method would panic. Because Set does not guarantee internal iteration order, this
// method in fact returns any element. The effect is only consistent when the Set
// contains only one item.
func (s Set) First() string {
	if s.Len() == 0 {
		panic("calling First on empty set")
	}

	for e := range s {
		return e
	}

	panic("unreachable code")
}

// Contains returns true if the Set contains the given value.
func (s Set) Contains(value string) bool {
	_, contains := s[value]
	return contains
}

// All returns true if all Set elements fulfill the criteria.
func (s Set) All(criteria func(element string) bool) bool {
	for e := range s {
		if ok := criteria(e); !ok {
			return false
		}
	}
	return true
}

// Any returns true if any Set element fulfills the criteria.
func (s Set) Any(criteria func(element string) bool) bool {
	for e := range s {
		if ok := criteria(e); ok {
			return true
		}
	}
	return false
}

// ContainsAll returns true if this Set is a super set of the given Set.
func (s Set) ContainsAll(t Set) bool {
	return t.All(s.Contains)
}

// ContainsAny returns true if this Set intersects with the given Set.
func (s Set) ContainsAny(t Set) bool {
	return t.Any(s.Contains)
}

// Equals returns true if the two Set contains exactly the same elements.
func (s Set) Equals(t Set) bool {
	return s.Len() == t.Len() && s.ContainsAll(t)
}

// Add adds a single element to the Set.
func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

// AddAll adds multiple elements to the Set.
func (s Set) AddAll(elem ...string) {
	for _, each := range elem {
		s.Add(each)
	}
}

// Remove removes a single element from the Set.
func (s Set) Remove(elem string) {
	delete(s, elem)
}

// Array returns the elements of this Set in a slice. This method is not stable, because the
// internal iteration order is not guaranteed.
func (s Set) Array() []string {
	var arr []string
	for elem := range s {
		arr = append(arr, elem)
	}
	return arr
}
