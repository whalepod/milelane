package main

import (
	"log"

	"github.com/whalepod/milelane/app/presentation"
)

func main() {
	r := presentation.Router()
	if err := r.Run(); err != nil {
		log.Fatalf("main error: %s", err.Error())
	}
}
