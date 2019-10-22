package controller

import (
	"cortago/controller"
	"net/http"
	"net/url"
)

type MainController struct {
	controller.Controller
}

type Person struct {
	Name   string
	Emails []string
}

func (c *MainController) Index(w http.ResponseWriter, r *http.Request, params url.Values) {
	c.Initiliaze(w, r)
	persons := struct {
		Persons []Person
	}{[]Person{
		Person{Name: "Drumil",
			Emails: []string{"drumilpatel720@gmail.com", "pnileshbhai@ee.iitr.ac.in"}},
		Person{Name: "Bismita",
			Emails: []string{"bguha@mt.iitr.ac.in"}},
	}}
	c.Directory = "views"
	c.RenderHTML("layout", persons)
	c.Terminate()
}
