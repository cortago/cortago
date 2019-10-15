package main

import (
	"Cortago/router"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	r := router.New(Root)
	r.Handle("GET", "/", Root)
	r.Handle("GET", "/users/:name", UserShow)

	r.StartApp()
}

// Root states root app
func Root(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprint(w, "Root!\n")
}

// UserShow jskldf
func UserShow(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprintf(w, "Hi %s", params["name"])
}
