package set

type Set[T comparable] struct {
	Elements []T
}

func (s *Set[T]) Add(element T) {
	if s.IndexOf(element) < 0 {
		s.Elements = append(s.Elements, element)
	}
}

func (s *Set[T]) Remove(element T) {
	if index := s.IndexOf(element); index > 0 {
		if index == len(s.Elements)-1 {
			s.Elements = s.Elements[:index]
		} else if index == 0 {
			s.Elements = s.Elements[index+1:]
		} else {
			s.Elements = append(s.Elements[:index], s.Elements[index+1:]...)
		}
	}
}

func (s *Set[T]) IndexOf(element T) int {
	for i, v := range s.Elements {
		if v == element {
			return i
		}
	}
	return -1
}
