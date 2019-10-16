package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/cortago/cortago/router"
)

func main() {
<<<<<<< HEAD
	router.Port = 8000
	r := router.New()
	r.Handle("/", Root)
	r.Handle("/users/:name", UserShow)

	y := router.New()
	y.AppendRouter("/drumil",r)
	y.Handle("/",sd)
=======
	router.Port = 5123
	r := router.New(Root)
	r.Handle("/", Root)
	r.Handle("/users/:name", UserShow)
>>>>>>> bfd89eef95b6d4f691e70fa08d73878dbe5b5935

	y.StartApp()
}

// Root states root app
func Root(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprint(w, "Root!\n")
}

// Root states root app
func sd(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprint(w, "Madarchod!\n")
}

// UserShow jskldf
func UserShow(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprintf(w, "Hi %s", params["name"])
}
