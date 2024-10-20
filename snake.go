package main

type Point struct {
	x int
	y int
}

const (
	MAX_SIZE = 10000
)

// - - - tail e1 e2 e3 head ---
type Snake struct {
	head int // pointer the head
	tail int // pointer to the tail
	v    [MAX_SIZE]Point
}

func MakeSnake() Snake {
	return Snake{head: 0, tail: 0}
}

func (s *Snake) relocate() {
	t := [MAX_SIZE]Point{}
	for i, j := s.tail, 0; i <= s.head; i, j = i+1, j+1 {
		t[j] = s.v[i]
	}
	s.v = t
	s.head = s.head - s.tail
	s.tail = 0
}

func (s *Snake) AddHead(p Point) {
	if s.head >= MAX_SIZE-1 {
		s.relocate()
	}
	s.head++
	s.v[s.head] = p
}

func (s *Snake) RemoveTail() {
	s.tail++
}

func (s Snake) Len() int {
	return s.head - s.tail
}

func (s Snake) IsEmpty() bool {
	return s.Len() == 0
}

func (s Snake) Head() Point {
	return s.v[s.head]
}

func (s Snake) Tail() Point {
	return s.v[s.tail]
}
