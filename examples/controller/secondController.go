package controller

import (
	"cortago/controller"
	"net/http"
	"net/url"
)

type SecondController struct {
	controller.Controller
}

func (c *SecondController) Index(w http.ResponseWriter, r *http.Request, params url.Values) {
	c.Initiliaze(w, r, "SecondController")

	c.RenderText(c.Name + " index is working")
}

func (c *SecondController) List(w http.ResponseWriter, r *http.Request, params url.Values) {
	c.Initiliaze(w, r, "SecondController")

	c.RenderText(c.Name + " list is working")
}
