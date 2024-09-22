package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
)

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

type Point struct {
	x int
	y int
}

// empty -> head: 0
type Snake struct {
	next int // pointer to add next element, init to 0
	tail int // oldest element added, default to -1
	s    []Point
}

func MakeSnake(l int) Snake {
	return Snake{0, -1, make([]Point, 0)}
}

func (s Snake) Len() int {
	return s.tail - s.next + 1
}

func (s Snake) IsEmpty() bool {
	return s.Len() > 0
}

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	// Draw initial boxes
	// drawBox(s, 1, 1, 42, 7, boxStyle, "Click and drag to draw a box")
	// drawBox(s, 5, 9, 32, 14, boxStyle, "Press C to reset")

	// drawText(s, 10, 42, 20, 50, boxStyle, "Click and drag to draw a box")

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	// Here's how to get the screen size when you need it.
	xmax, ymax := s.Size()

	// Here's an example of how to inject a keystroke where it will
	// be picked up by the next PollEvent call.  Note that the
	// queue is LIFO, it has a limited length, and PostEvent() can
	// return an error.
	// s.PostEvent(tcell.NewEventKey(tcell.KeyRune, rune('a'), 0))

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	key_pressed := make(chan bool)

	// ?? lock!
	x, y := 20, 20
	up, down, left, right := true, false, false, false

	go func() {
		for {
			// Poll event
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					close(done)
				} else if ev.Key() == tcell.KeyCtrlL {
					s.Sync()
				} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
					s.Clear()
				} else if ev.Key() == tcell.KeyRight {
					right = true
					left = false
					up = false
					down = false
					key_pressed <- true
				} else if ev.Key() == tcell.KeyLeft {
					right = false
					left = true
					up = false
					down = false
					key_pressed <- true
				} else if ev.Key() == tcell.KeyUp {
					right = false
					left = false
					up = true
					down = false
					key_pressed <- true
				} else if ev.Key() == tcell.KeyDown {
					right = false
					left = false
					up = false
					down = true
					key_pressed <- true
				}
			}
		}
	}()

	// Move
	for {
		select {
		case <-done:
			return
		case <-key_pressed:
			s.SetContent(x, y, ' ', nil, boxStyle)
			if right {
				if x < xmax {
					x++
				} else {
					x = 0
				}
			} else if left {
				if x > 0 {
					x--
				} else {
					x = xmax
				}

			} else if up {
				if y > 0 {
					y--
				} else {
					y = ymax
				}
			} else if down {
				if y < ymax {
					y++
				} else {
					y = 0
				}
			}
		case <-ticker.C:
			s.SetContent(x, y, ' ', nil, boxStyle)
			if right {
				x++
			} else if left {
				x--
			} else if up {
				y--
			} else if down {
				y++
			}
		}
		s.Show()
	}
}
