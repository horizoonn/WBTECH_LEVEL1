package main

import "fmt"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](values ...T) Set[T] {
	s := make(Set[T], len(values))
	for _, v := range values {
		s[v] = struct{}{}
	}
	return s
}

func (s Set[T]) Add(values ...T) {
	for _, v := range values {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Has(val T) bool {
	_, ok := s[val]
	return ok
}

func (s Set[T]) Delete(val T) {
	delete(s, val)
}

func (s Set[T]) Values() []T {
	values := make([]T, 0, len(s))
	for v := range s {
		values = append(values, v)
	}
	return values
}

func Intersection[T comparable](a, b Set[T]) Set[T] {
	if len(a) > len(b) {
		a, b = b, a
	}
	res := make(Set[T], len(a))
	for v := range a {
		if _, ok := b[v]; ok {
			res[v] = struct{}{}
		}
	}

	return res
} 

func Union[T comparable](a, b Set[T]) Set[T] {
	res := make(Set[T], len(a)+len(b))
	for v := range a {
		res[v] = struct{}{}
	}
	for v := range b {
		res[v] = struct{}{}
	}
	return res
}

func (s Set[T]) String() string {
    b := make([]byte, 0, 64)
    b = append(b, '{')
    first := true
    for v := range s {
        if !first {
            b = append(b, ',', ' ')
        }
        first = false
        b = append(b, fmt.Sprintf("%v", v)...)
    }
    b = append(b, '}')
    return string(b)
}

func main() {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(2, 3, 4)

	res := Intersection(s1, s2)
	fmt.Println(res)
}