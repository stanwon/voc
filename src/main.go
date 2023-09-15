package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

var fLog *os.File

type users struct {
	name []string
}

func main() {
	run()
}

func run() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	err = screen.Init()
	if err != nil {
		log.Fatal(err)
	}

	ui := newUI(screen)
	nav := newNav()
	voc := newVoc()
	app := newApp(ui, nav, voc)

	app.loop()

	screen.Fini()
	fLog.Close()
}

func init() {
	f, err := os.OpenFile("log", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}

	log.SetOutput(f)
}
