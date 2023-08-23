package algo

import "math"

//"strings"
//"unicode/utf8"

type Set[T comparable] struct {
	Map map[T]bool
}

func NewSet[T comparable](size int) *Set[T] {
	return &Set[T]{make(map[T]bool, size)}
}

func (s *Set[T]) Add(d T) bool {
	if !s.Map[d] {
		s.Map[d] = true
	}

	return s.Map[d]
}

func Has[T comparable](slice []T, el T) bool {
	for _, v := range slice {
		if v == el {
			return true
		}
	}

	return false
}

func min(T, S []int) int {
	min, max := -1, math.Inf(0)
	for i, t := range T {
		if float64(t) < max && !Has(S, i) {
			max, min = float64(t), i
		}
	}

	return min
}
