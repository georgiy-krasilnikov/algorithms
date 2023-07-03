package algo

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
