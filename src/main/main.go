package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type Canvas struct {
	screen       tcell.Screen
	defaultStyle tcell.Style
	xMax         int
	yMax         int
	curLine      int
}

func main() {
  testTcell()
	/* canvas := newCanvas()
	canvas.init()
	var buf [][]rune
	buf = append(buf, []rune("-----choose dictionary-----"))
	buf = append(buf, []rune("1: 初中-乱序"))
	buf = append(buf, []rune("2: 高中-乱序"))
	buf = append(buf, []rune("3: 四级-乱序"))
	buf = append(buf, []rune("4: 六级-乱序"))
	buf = append(buf, []rune("5: 考研-乱序"))
	buf = append(buf, []rune("6: 托福-乱序"))
	buf = append(buf, []rune("7: SAT-乱序"))
	buf = append(buf, []rune("q: quit"))
	buf = append(buf, []rune("---------------------------"))
	buf = append(buf, []rune("choose: \n"))
	for i := 0; i < len(buf); i++ {
    xNext := 0
		for j := 0; j < len(buf[i]); j++ {
			log.Println(runewidth.RuneWidth(buf[i][j]))
			canvas.screen.SetContent(xNext, i, buf[i][j], nil, canvas.defaultStyle)
      xNext += runewidth.RuneWidth(buf[i][j])
		}
	}
	// voc(canvas)
	canvas.screen.Show()
	time.Sleep(time.Second * 5)
	canvas.finish() */
}

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

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}

func testTcell(){
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
	drawBox(s, 1, 1, 42, 7, boxStyle, "Click and drag to draw a box")
	drawBox(s, 5, 9, 32, 14, boxStyle, "Press C to reset")

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
	// xmax, ymax := s.Size()

	// Here's an example of how to inject a keystroke where it will
	// be picked up by the next PollEvent call.  Note that the
	// queue is LIFO, it has a limited length, and PostEvent() can
	// return an error.
	// s.PostEvent(tcell.NewEventKey(tcell.KeyRune, rune('a'), 0))

	// Event loop
	ox, oy := -1, -1
	for {
		// Update screen
		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				s.Clear()
			}
		case *tcell.EventMouse:
			x, y := ev.Position()

			switch ev.Buttons() {
			case tcell.Button1, tcell.Button2:
				if ox < 0 {
					ox, oy = x, y // record location when click started
				}

			case tcell.ButtonNone:
				if ox >= 0 {
					label := fmt.Sprintf("%d,%d to %d,%d", ox, oy, x, y)
					drawBox(s, ox, oy, x, y, boxStyle, label)
					ox, oy = -1, -1
				}
			}
		}
	}
}

func (c *Canvas) init() {
	if err := c.screen.Init(); err != nil {
		log.Fatal(err)
	}
	c.screen.Clear()
}
func (c *Canvas) finish() {
	c.screen.Fini()
}

func newCanvas() *Canvas {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}

	defaultStyle := tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorRed)
	screen.SetStyle(defaultStyle)
	x, y := screen.Size()

	return &Canvas{
		screen:       screen,
		defaultStyle: defaultStyle,
		xMax:         x,
		yMax:         y,
	}
}

func voc(canvas *Canvas) {
	for {
		path := chooseVoc(canvas)
		if path == "" {
			return
		}
		zh, en := readRaw1(path)
		switch readAction() {
		case "1":
			practice(zh, en)
		case "2":
			zhToEn(zh, en)
		case "q":
			// continue
		default:
			fmt.Println("unknown action.")
		}
	}
}

func (c *Canvas) printline() {

}

func chooseVoc(canvas *Canvas) (path string) {
	for {
		fmt.Println("-----choose dictionary-----")
		fmt.Println("1: 初中-乱序")
		fmt.Println("2: 高中-乱序")
		fmt.Println("3: 四级-乱序")
		fmt.Println("4: 六级-乱序")
		fmt.Println("5: 考研-乱序")
		fmt.Println("6: 托福-乱序")
		fmt.Println("7: SAT-乱序")
		fmt.Println("q: quit")
		fmt.Println("---------------------------")
		fmt.Print("choose: ")
		sIn := bufio.NewScanner(os.Stdin)
		sIn.Scan()
		choose := sIn.Text()
		switch choose {
		case "1":
			path = "english-vocabulary/1.初中-乱序.txt"
			return
		case "2":
			path = "english-vocabulary/2.高中-乱序.txt"
			return
		case "3":
			path = "english-vocabulary/3.四级-乱序.txt"
			return
		case "4":
			path = "english-vocabulary/4.六级-乱序.txt"
			return
		case "5":
			path = "english-vocabulary/5.考研-乱序.txt"
			return
		case "6":
			path = "english-vocabulary/6.托福-乱序.txt"
			return
		case "7":
			path = "english-vocabulary/7.SAT-乱序.txt"
			return
		case "q":
			path = ""
			return
		}
	}
}

func readAction() (action string) {
	for {
		fmt.Println("---------------------------")
		fmt.Println("1: show vocabulary")
		fmt.Println("2: zh to en")
		fmt.Println("q: back to choose dictionary")
		fmt.Println("---------------------------")
		fmt.Print("choose: ")
		// fmt.Println("3: en to zh")
		sIn := bufio.NewScanner(os.Stdin)
		sIn.Scan()
		action = sIn.Text()
		if "q" == action || "1" == action || "2" == action {
			break
		}
	}
	return
}

func readRaw1(path string) (zh, en []string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scaner := bufio.NewScanner(f)

	count := 1
	for scaner.Scan() {
		if err = scaner.Err(); err != nil {
			log.Fatal(err)
		}
		raw := scaner.Text()
		before, after, ok := strings.Cut(raw, "\t")
		if ok {
			en = append(en, before)
			zh = append(zh, strings.Trim(after, " "))
		}
		count++
	}
	fmt.Println("the number of word: ", count)
	return
}

func readRaw() (zh, en []string) {
	enPath := "./en.txt"
	zhPath := "./zh.txt"
	fEn, err := os.Open(enPath)
	if err != nil {
		log.Fatal(err)
	}
	fZh, err := os.Open(zhPath)
	if err != nil {
		log.Fatal(err)
	}
	defer fEn.Close()
	defer fZh.Close()
	sEn := bufio.NewScanner(fEn)
	sZh := bufio.NewScanner(fZh)

	for sEn.Scan() {
		if err = sEn.Err(); err != nil {
			log.Fatal(err)
		}
		en = append(en, sEn.Text())
	}

	for sZh.Scan() {
		if err = sZh.Err(); err != nil {
			log.Fatal(err)
		}
		zh = append(zh, sZh.Text())
	}
	return
}

var indexs = make(map[int]int)

func checkIdx(i int) bool {
	if _, ok := indexs[i]; ok {
		return true
	}
	indexs[i] = i
	return false
}

func practice(zh, en []string) {
	for {
		idx := rand.Intn(len(zh))
		if checkIdx(idx) {
			continue
		}
		fmt.Println(en[idx])
		fmt.Println(zh[idx])
		sIn := bufio.NewScanner(os.Stdin)
		sIn.Scan()
		sIn.Text()
	}
}

func zhToEn(zh, en []string) {
	for {
		idx := rand.Intn(len(zh))
		sIn := bufio.NewScanner(os.Stdin)
		count := 1
		fmt.Println("-------------------------------------")
		fmt.Println(zh[idx])
		for {
			fmt.Print("answer: ")
			sIn.Scan()
			result := sIn.Text()
			if result == "j" {
				fmt.Println(en[idx])
				break
			}

			if result == en[idx] {
				break
			} else if count <= len(en[idx]) {
				fmt.Print("tips: ")
				fmt.Println(en[idx][:count])
				count++
			} else if count > len(en[idx]) {
				fmt.Println("!!!笨蛋!!!")
				break
			}
		}
	}
}

func enToZh(zh, en []string) {
	for {
		idx := rand.Intn(len(en))
		fmt.Println(en[idx])
		sIn := bufio.NewScanner(os.Stdin)
		for {
			sIn.Scan()
			result := sIn.Text()
			if result == "j" {
				fmt.Println(zh[idx])
				break
			}

			fmt.Println(len(en[idx]))
			if result == zh[idx] {
				break
			}
		}
	}
}
