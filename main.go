package main

import (
	"github.com/andlabs/ui"
	"github.com/parnurzeal/gorequest"
)

func main() {
	err := ui.Main(func() {
		window := ui.NewWindow("Popcorn Time Go", 200, 100, true)

        wrapper := ui.NewVerticalBox()
       
		window.SetChild(wrapper)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()

		go counter(wrapper)
	})
	if err != nil {
		panic(err)
	}
}
func counter(wrapper *ui.Box) {
	request := gorequest.New()
	_, body, _:= request.Get("http://tv-v2.api-fetch.website/movies/1").End()


	ui.QueueMain(func() {
	
		box := ui.NewVerticalBox()
		box2 := ui.NewVerticalBox()
		name := ui.NewEntry()
        button := ui.NewButton("Greet")
        greeting := ui.NewLabel("")
    
	    box.Append(ui.NewLabel("Enter your name:"), false)

		name.SetText("number " + body)
		
        box.Append(name, false)
        box.Append(button, false)
        box.Append(greeting, false)
	
		wrapper.Append(box, false)
	})
}
