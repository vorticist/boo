package main

import (
	"fmt"
	"github.com/vorticist/boo/client"
	"github.com/vorticist/boo/crypt"
	"github.com/vorticist/boo/subs"
	"github.com/vorticist/boo/ui"
	"gitlab.com/vorticist/logger"
	"os"

	"github.com/getlantern/systray"
)

var (
	exit    chan bool
	appOpen = false
)

func main() {
	exit = make(chan bool)
	subs.NewEventListener()
	booClient := client.New()
	key, err := booClient.GetCryptKey()
	if err != nil {
		logger.Errorf("error getting crypto key: %v", err)
		return
	}
	crypt.New(key)
	subscriptions(booClient)
	go systray.Run(onReady, onExit)
	<-exit
}

func onReady() {
	go func() {
		systray.SetIcon(getIcon("/home/vorticist/boo/assets/auto_stories_black_24dp.svg"))
		systray.SetTitle("Book of Omens")
		systray.SetTooltip("Book of Omens")
		openApp := systray.AddMenuItem("Open App", "Open App")
		systray.AddSeparator()
		quit := systray.AddMenuItem("Quit", "Quit")

		for {
			select {
			case <-openApp.ClickedCh:
				logger.Info("Open App")
				if !appOpen {
					// TODO: need a way to detect when the app window is closed and update this value again
					appOpen = true
				}
				go ui.StartApp()
			case <-quit.ClickedCh:
				logger.Info("Quit")
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
