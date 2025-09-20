package myset

import (
	"iter"
	"maps"
)

type Set[T any] interface {
	Add(element T)
	Clear()
	DeepCopy() Set[T]
	Difference(other Set[T]) Set[T]
	Has(element T) bool
	Intersection(other Set[T]) Set[T]
	IsEmpty() bool
	IsSubset(other Set[T]) bool
	IsSuperset(other Set[T]) bool
	Iter() iter.Seq[T]
	Len() int
	Remove(element T)
	Union(other Set[T]) Set[T]
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

func (s *set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

func (s *set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

func (s *set[T]) DeepCopy() Set[T] {
	elements := make(map[T]struct{}, s.Len())
	maps.Copy(elements, s.elements)
	return &set[T]{elements: elements}
}

func (s *set[T]) Difference(other Set[T]) Set[T] {
	elements := make(map[T]struct{})
	for element := range s.Iter() {
		if !other.Has(element) {
			elements[element] = struct{}{}
		}
	}
	for element := range other.Iter() {
		if !s.Has(element) {
			elements[element] = struct{}{}
		}
	}
	return &set[T]{elements: elements}
}

func (s *set[T]) Has(element T) bool {
	_, ok := s.elements[element]
	return ok
}

func (s *set[T]) Intersection(other Set[T]) Set[T] {
	elements := make(map[T]struct{})
	for element := range s.Iter() {
		if other.Has(element) {
			elements[element] = struct{}{}
		}
	}
	return &set[T]{elements: elements}
}

func (s *set[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *set[T]) IsSubset(other Set[T]) bool {
	for element := range s.Iter() {
		if !other.Has(element) {
			return false
		}
	}
	return true
}

func (s *set[T]) IsSuperset(other Set[T]) bool {
	for element := range other.Iter() {
		if !s.Has(element) {
			return false
		}
	}
	return true
}

func (s *set[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for element := range s.elements {
			if !yield(element) {
				return
			}
		}
	}
}

func (s *set[T]) Len() int {
	return len(s.elements)
}

func (s *set[T]) Remove(element T) {
	delete(s.elements, element)
}

func (s *set[T]) Union(other Set[T]) Set[T] {
	elements := make(map[T]struct{})
	for element := range s.Iter() {
		elements[element] = struct{}{}
	}
	for element := range other.Iter() {
		elements[element] = struct{}{}
	}
	return &set[T]{elements: elements}
}

func (s *set[T]) Values() []T {
	values := make([]T, 0, s.Len())
	for element := range s.Iter() {
		values = append(values, element)
	}
	return values
}
