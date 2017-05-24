package main

import (
	"github.com/andlabs/ui"
	"github.com/KeizerDev/popcorntime-go/core"
	"fmt"
)

func main() {
	c := core.Core{Endpoint: core.ENDPOINT}
	api := core.API{
		Core: c,
	}

	err := ui.Main(func() {
		window := ui.NewWindow("Popcorn Time Go", 200, 100, true)

        wrapper := ui.NewVerticalBox()

		window.SetChild(wrapper)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()

		go counter(wrapper, api)
	})
	if err != nil {
		panic(err)
	}
}
func counter(wrapper *ui.Box, api core.API) {

	resp, err := api.Movies(1, core.SORT_TRENDING, core.ORDER_NEWEST, "")
	if err != nil {
		panic(err)
	}


	ui.QueueMain(func() {
		for _, value := range resp {
			fmt.Println(value.Title)

			box := ui.NewVerticalBox()
			name := ui.NewEntry()

			name.SetText(value.Title)

			box.Append(name, false)

			wrapper.Append(box, false)
		}
	 })
}