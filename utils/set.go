package utils

func NewSet[T comparable](vals ...T) Set[T] {
	set := Set[T]{}
	set.Add(vals...)
	return set
}

type Set[T comparable] map[T]struct{}

// Add a value to the set.
func (s Set[T]) Add(vals ...T) {
	for _, val := range vals {
		s[val] = struct{}{}
	}
}

// Remove a value from the set.
func (s Set[T]) Remove(vals ...T) {
	for _, val := range vals {
		delete(s, val)
	}
}

// Values returns all the unique values in the set.
func (s Set[T]) Values() []T {
	var values []T
	for val := range s {
		values = append(values, val)
	}
	return values
}

// Contains returns if a value is in the set.
func (s Set[T]) Contains(val T) bool {
	_, found := s[val]
	return found
}

// ContainsAny returns if any provided value is in the set.
func (s Set[T]) ContainsAny(val ...T) bool {
	for _, x := range val {
		if _, found := s[x]; found {
			return true
		}
	}
	return false
}

// IntersectionWith returns the set of values in both sets.
func (s Set[T]) IntersectionWith(other Set[T]) Set[T] {
	intersection := NewSet[T]()
	for val := range s {
		if other.Contains(val) {
			intersection.Add(val)
		}
	}
	return intersection
}

// IdenticalTo returns true if all values exist in both sets.
func (s Set[T]) IdenticalTo(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for val := range s {
		if !other.Contains(val) {
			return false
		}
	}
	return true
}
