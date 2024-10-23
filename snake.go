package main

import "github.com/gammazero/deque"

type Point struct {
	x int
	y int
}

const (
	MAX_SIZE = 10000
)

// - - - tail e1 e2 e3 head ---
type Snake struct {
	v *deque.Deque[Point]
}

func MakeSnake() Snake {
	return Snake{v: deque.New[Point]()}
}

func (s *Snake) AddHead(p Point) {
	s.v.PushFront(p)
}

func (s *Snake) RemoveTail() {
	s.v.PopBack()
}

func (s Snake) Len() int {
	return s.v.Len()
}

func (s Snake) IsEmpty() bool {
	return s.v.Len() == 0
}

func (s Snake) Head() Point {
	return s.v.Front()
}

func (s Snake) Tail() Point {
	return s.v.Back()
}
