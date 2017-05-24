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
	//request := gorequest.New()
	//_, body, _:= request.Get("http://tv-v2.api-fetch.website/movies/1").End()

	resp, err := api.Movies(1, core.SORT_TRENDING, core.ORDER_NEWEST, "")
	if err != nil {
		panic(err)
	}
	
	fmt.Println(resp)

	//for i := range body {
	//	fmt.Println(body[i])
	//}
	// ui.QueueMain(func() {
	
	// 	box := ui.NewVerticalBox()
	// 	box2 := ui.NewVerticalBox()
	// 	name := ui.NewEntry()
    //     button := ui.NewButton("Greet")
    //     greeting := ui.NewLabel("")
    
	//     box.Append(ui.NewLabel("Enter your name:"), false)

	// 	name.SetText("number " + body)
		
    //     box.Append(name, false)
    //     box.Append(button, false)
    //     box.Append(greeting, false)
	
	// 	wrapper.Append(box, false)
	// })
}


//func ParseMovies(body string) (*StationsResponse, error) {
//    var s = new(StationsResponse)
//    err := json.Unmarshal(body, &s)
//    if(err != nil){
//        fmt.Println("whoops:", err)
//    }
//    return s, err
//}