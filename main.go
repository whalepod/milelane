package main

import (
	"log"
	"time"

	"github.com/whalepod/milelane/app/presentation"
)

func init() {
	time.Local = time.FixedZone("JST", 9*60*60)
}

func main() {
	r := presentation.Router()
	if err := r.Run(); err != nil {
		log.Fatalf("main error: %s", err.Error())
	}
}
