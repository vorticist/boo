package main

import (
	"fmt"
	"github.com/vorticist/boo/ui"
	"os"

	"github.com/getlantern/systray"
)

var (
	exit chan bool
)

func main() {
	exit = make(chan bool)
	go systray.Run(onReady, onExit)
	<-exit
}

func onReady() {
	go func() {
		systray.SetIcon(getIcon("assets/auto_stories_black_24dp.svg"))
		systray.SetTitle("Book of Omens")
		systray.SetTooltip("Book of Omens")
		openApp := systray.AddMenuItem("Open App", "Open App")
		systray.AddSeparator()
		quit := systray.AddMenuItem("Quit", "Quit")

		for {
			select {
			case <-openApp.ClickedCh:
				fmt.Println("Open App")
				go ui.StartApp()
			case <-quit.ClickedCh:
				fmt.Println("Quit")
				exit <- true
				systray.Quit()
			}

		}
	}()
}

func onExit() {
	// Cleaning up stuff here.
}

func getIcon(s string) []byte {
	b, err := os.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
