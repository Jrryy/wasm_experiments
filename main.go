// A basic HTTP server.
// By default, it serves the current working directory on port 8080.
package main

import (
	"github.com/Jrryy/wasm_experiments/internal/calculator"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"log"
	"net/http"
)

func main() {
	app.Route("/", &calculator.Calculator{})
	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name:        "wasm experiments",
		Description: "Some wasm experiments",
	})
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
