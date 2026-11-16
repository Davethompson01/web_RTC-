package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	setupApp()
	log.Println(http.ListenAndServe(":8080", nil))
}

func setupApp() *http.ServeMux {

	// creates new router
	mux := http.NewServeMux()

	// when user enter default url, GO returns them back to the browser
	mux.Handle("/", http.FileServer(http.Dir("/web/views")))

	// check for file like app.css or .js file from the html request
	output := http.FileServer(http.Dir("./web/static"))
	// removes /static from the requested path to aviod redudance
	mux.Handle("/static/", http.StripPrefix("/static/", output))

	ctx := context.Background()

	// mux
}
