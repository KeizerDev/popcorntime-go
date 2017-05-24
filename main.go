package main

import (
	"github.com/andlabs/ui"
	"strconv"
	"time"
)

func main() {
	err := ui.Main(func() {
		window := ui.NewWindow("title", 200, 100, false)
		label := ui.NewLabel("text")
		window.SetChild(label)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
		go counter(label)
	})
	if err != nil {
		panic(err)
	}
}
func counter(label *ui.Label) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		ui.QueueMain(func() {
			label.SetText("number " + strconv.Itoa(i))
		})
	}
}
