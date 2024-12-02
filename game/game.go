package game

import "sync"

type Game struct {
	x, y       int
	xmax, ymax int
	xgem, ygem int
	l          sync.Mutex
}

//func (g *Game) locked(fn func()) {
//	g.l.Lock()
//	defer g.l.Unlock()
//	fn()
//}

func NewGame(x, y, xmax, ymax, xgem, ygem int) *Game {
	return &Game{x: x, y: y, xmax: xmax, ymax: ymax, xgem: xgem, ygem: ygem}
}

func (g *Game) MoveRight() {
	g.l.Lock()
	defer g.l.Unlock()
	if g.x < g.xmax {
		g.x++
	} else {
		g.x = 0
	}
}

func (g *Game) MoveLeft() {
	g.l.Lock()
	defer g.l.Unlock()
	if g.x > 0 {
		g.x--
	} else {
		g.x = g.xmax
	}
}

func (g *Game) MoveUp() {
	g.l.Lock()
	defer g.l.Unlock()
	if g.y > 0 {
		g.y--
	} else {
		g.y = g.ymax
	}
}

func (g *Game) MoveDown() {
	g.l.Lock()
	defer g.l.Unlock()
	if g.y < g.ymax {
		g.y++
	} else {
		g.y = 0
	}
}

func (g *Game) UpdateGem(x, y int) {
	g.l.Lock()
	defer g.l.Unlock()
	g.xgem = x
	g.ygem = y
}

func (g *Game) GetGemX() int {
	g.l.Lock()
	defer g.l.Unlock()
	return g.xgem
}

func (g *Game) GetGemY() int {
	g.l.Lock()
	defer g.l.Unlock()
	return g.ygem
}

func (g *Game) GetX() int {
	g.l.Lock()
	defer g.l.Unlock()
	return g.x
}

func (g *Game) GetY() int {
	g.l.Lock()
	defer g.l.Unlock()
	return g.y
}

func (g *Game) HasReachedGem() bool {
	g.l.Lock()
	defer g.l.Unlock()
	return g.x == g.xgem && g.y == g.ygem
}
