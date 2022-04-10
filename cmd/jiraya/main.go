package main

import "github.com/juicyluv/jiraya/internal/jiraya/app"

func main() {
	a := app.New()

	if err := a.Start(); err != nil {
		panic(err)
	}
}
