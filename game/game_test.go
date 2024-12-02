package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameZero(t *testing.T) {
	g := NewGame(0, 0, 10, 10, 5, 5)
	g.MoveRight()
	g.MoveLeft()
	g.MoveUp()
	g.MoveDown()
	assert.Equal(t, 0, g.x)
	assert.Equal(t, 0, g.y)
}

func TestGameMoveNoWrap(t *testing.T) {
	g := NewGame(0, 0, 10, 10, 5, 5)
	g.MoveDown()
	g.MoveRight()
	assert.Equal(t, 1, g.x)
	assert.Equal(t, 1, g.y)
}

func TestGameMoveWrap(t *testing.T) {
	g := NewGame(0, 0, 10, 10, 5, 5)
	g.MoveUp()
	g.MoveLeft()
	assert.Equal(t, 10, g.x)
	assert.Equal(t, 10, g.y)
}
