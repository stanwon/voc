package main

import (
	"log"
	"os"
	"sort"
	"strings"
)

type nav struct {
	pages   []*page
	curPage int
	pageNum int
}

type page struct {
	init  bool
	lines []string
	cur   int
	nums  int
}

func newNav() *nav {
	return &nav{}
}

func (nav *nav) readPages() {
	f, err := os.Open("english-vocabulary")
	if err != nil {
		log.Println(err)
	}
	names, err := f.Readdirnames(-1)
	if err != nil {
		log.Println(err)
	}

	var lineDict []string
	for _, name := range names {
		if strings.HasSuffix(name, ".txt") {
			lineDict = append(lineDict, name)
		}
	}
	sort.Strings(lineDict)
	nav.curPage = 0

	nav.pageNum++
	nav.pages = append(nav.pages, &page{
		lines: lineDict,
		cur:   0,
		nums:  len(lineDict),
		init:  true,
	})

	var lineMode []string
	lineMode = lineMode[:0]
	lineMode = append(lineMode, "list")
	lineMode = append(lineMode, "zh to en")

	nav.pageNum++
	nav.pages = append(nav.pages, &page{
		lines: lineMode,
		cur:   0,
		nums:  len(lineMode),
		init:  true,
	})
}

func (nav *nav) up() {
  if nav.curPage == 2 {
    return
  }
	page := nav.pages[nav.curPage]
	if page.init {
		page.cur = (page.cur + 1) % page.nums
		nav.pages[nav.curPage] = page
	}
}

func (nav *nav) down() {
  if nav.curPage == 2 {
    return
  }
	page := nav.pages[nav.curPage]
	if page.init {
		page.cur = (page.cur - 1)
		if page.cur < 0 {
			page.cur = page.nums - 1
		}
		nav.pages[nav.curPage] = page
	}
}

func (nav *nav) prevPage(app *app) {
	nav.curPage -= 1
	if nav.curPage < 0 {
		app.quitChan <- struct{}{}
	}
}

func (nav *nav) nextPage(app *app) {
	nav.curPage++
	if nav.curPage >= nav.pageNum {
		nav.prevPage(app)
	}
}
