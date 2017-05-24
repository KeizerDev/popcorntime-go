package main

import (
	"github.com/andlabs/ui"
	"github.com/parnurzeal/gorequest"
)

func main() {
	err := ui.Main(func() {
		window := ui.NewWindow("Popcorn Time Go", 200, 100, false)
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
	request := gorequest.New()
	_, body, _:= request.Get("http://tv-v2.api-fetch.website/movies/1").End()

	ui.QueueMain(func() {
		label.SetText("number " + body)
	})
}
