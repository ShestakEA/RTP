package main

import (
	"flag"
	"net/http"
)

func App(rtp float64) http.Handler {

	// Хэндлер с генерацией
	router := http.NewServeMux()
	NewRandomHandler(router, rtp)

	return router
}

func main() {
	rtpInit := flag.Float64("rtp", 0.0, "init val")
	flag.Parse()

	app := App(*rtpInit)
	if err := http.ListenAndServe(":64333", app); err != nil {
		panic(err)
	}
}
