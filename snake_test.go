package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptySnake(t *testing.T) {
	s := MakeSnake()
	assert.Equal(t, 0, s.Len())
}

func TestAdd(t *testing.T) {
	s := MakeSnake()
	s.AddHead(Point{1, 3})
	assert.Equal(t, 1, s.Len())
}

func TestRemove(t *testing.T) {
	s := MakeSnake()
	s.AddHead(Point{1, 3})
	s.RemoveTail()
	assert.Equal(t, 0, s.Len())
}
