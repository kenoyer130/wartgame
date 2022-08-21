package models

import "math/rand"

type Stack struct {
	vals []interface{}
}

func (s *Stack) Array() []interface{}{
	return s.vals
}

func (s *Stack) Push(val interface{}) {
	s.vals = append(s.vals, val)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.vals) == 0 {
		return nil, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s *Stack) Peek() (interface{}, bool) {
	if len(s.vals) == 0 {
		return nil, false
	}
	top := s.vals[len(s.vals)-1]
	return top, true
}

func (s *Stack) Count() int {
	return len(s.vals)
}

func (s *Stack) Randomize() {
	currentIndex := s.Count()
	size := s.Count()

	// While there remain elements to shuffle.
	for currentIndex != 0 {

		// Pick a remaining element.
		randomIndex := rand.Intn(size)

		// And swap it with the current element.
		s.Swap(currentIndex-1, randomIndex)

		currentIndex--
	}
}

func (s *Stack) Swap(i int, x int) {
	val1 := s.vals[i]
	val2 := s.vals[x]

	s.vals[i] = val2
	s.vals[x] = val1
}
