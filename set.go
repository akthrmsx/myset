package myset

type Set[T comparable] interface {
	Clear()
	Has(element T) bool
	Add(element T)
	IsEmpty() bool
	Len() int
	Remove(element T)
	Values() []T
}

type set[T comparable] struct {
	elements map[T]struct{}
}

func New[T comparable](values ...T) Set[T] {
	elements := make(map[T]struct{}, len(values))
	for _, value := range values {
		elements[value] = struct{}{}
	}
	return &set[T]{elements: elements}
}

func (s *set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

func (s *set[T]) Has(element T) bool {
	_, ok := s.elements[element]
	return ok
}

func (s *set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

func (s *set[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *set[T]) Len() int {
	return len(s.elements)
}

func (s *set[T]) Remove(element T) {
	delete(s.elements, element)
}

func (s *set[T]) Values() []T {
	values := make([]T, 0, s.Len())
	for element := range s.elements {
		values = append(values, element)
	}
	return values
}
