package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/cortago/cortago/router"
)

func main() {
	router.Port = 5123
	r := router.New(Root)
	r.Handle("/", Root)
	r.Handle("/users/:name", UserShow)

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
