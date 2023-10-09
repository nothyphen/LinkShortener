package main

import "linkshortner/routers"

func main() {

	route := routers.Urls()
	err := route.Run()
	if err != nil {
		panic(err)
	}
}