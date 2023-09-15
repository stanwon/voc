package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gdamore/tcell/v2"
)

type app struct {
	ui       *ui
	nav      *nav
	voc      *voc
	quitChan chan struct{}
}

func newApp(ui *ui, nav *nav, voc *voc) *app {
	quitChan := make(chan struct{}, 1)
	app := &app{
		ui:       ui,
		nav:      nav,
		voc:      voc,
		quitChan: quitChan,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		switch <-sigChan {
		case os.Interrupt:
			return
		case syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT:
			os.Exit(3)
			return
		}
	}()

	return app
}

func (app *app) loop() {
	app.nav.readPages()
	app.voc.readVoc(app.nav)

	for {
		select {
		case <-app.quitChan:
			return
		case tev := <-app.ui.tevChan:
			switch ev := tev.(type) {
			case *tcell.EventResize:
				app.ui.draw(app.nav, app.voc)
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyCtrlC:
					return
				case tcell.KeyEsc:
					app.nav.prevPage(app)
					app.ui.draw(app.nav, app.voc)
				case tcell.KeyEnter:
					app.nav.nextPage(app)
					app.ui.draw(app.nav, app.voc)
				case tcell.KeyRune:
					switch ev.Rune() {
					case 'q':
						return
					case 'j':
						app.nav.up()
						app.ui.draw(app.nav, app.voc)
					case 'k':
						app.nav.down()
						app.ui.draw(app.nav, app.voc)
					}
				}
			}
		}
	}

}
