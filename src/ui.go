package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

type ui struct {
	screen  tcell.Screen
	tevChan chan tcell.Event
	win     *win
}

func newUI(screen tcell.Screen) *ui {
	wtot, htot := screen.Size()
	ui := &ui{
		screen:  screen,
		tevChan: make(chan tcell.Event, 1000),
		win:     newWin(wtot, htot, 0, 0),
	}

	go ui.pollEvent()

	return ui
}

func (ui *ui) pollEvent() {
	for {
		ui.tevChan <- ui.screen.PollEvent()
	}
}

func (ui *ui) draw(nav *nav, voc *voc) {
	if nav.curPage < 0 {
		return
	}
	// log.Println(nav.pages[nav.curPage].lines[0])

	for y := ui.win.y; y < ui.win.h; y++ {
		for x := ui.win.x; x < ui.win.w; x++ {
			ui.screen.SetContent(x, y, ' ', nil, tcell.StyleDefault)
		}
	}

	if nav.curPage == 0 || nav.curPage == 1 {
		for i, line := range nav.pages[nav.curPage].lines {
			var pos int
			st := tcell.StyleDefault
			for _, s := range line {
				if i == nav.pages[nav.curPage].cur {
					st = st.Reverse(true)
				}
				ui.screen.SetContent(pos, i, s, nil, st)
				pos += runewidth.RuneWidth(s)
			}
		}
	} else if nav.curPage == 2 {
		for i, line := range voc.en[nav.pages[0].cur] {
			var pos int
			st := tcell.StyleDefault
			for _, s := range line {
				if i == nav.pages[nav.curPage].cur {
					st = st.Reverse(true)
				}
				ui.screen.SetContent(pos, i, s, nil, st)
				pos += runewidth.RuneWidth(s)
			}
		}
	}

	ui.screen.Show()
}
