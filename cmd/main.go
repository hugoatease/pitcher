package main

import (
	"log"

	"github.com/hugoatease/pitcher"
)

func main() {
	app, err := pitcher.NewApp()
	if err != nil {
		log.Fatal(err)
		return
	}

	app.Serve()
}
