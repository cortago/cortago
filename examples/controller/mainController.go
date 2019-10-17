package controller

import (
	"cortago/controller"
	"net/http"
	"net/url"
)

type MainController struct {
	controller.Controller
}

func (c *MainController) Index(w http.ResponseWriter, r *http.Request, params url.Values) {
	c.Initiliaze(w, r, "MainController")
	c.RenderText(c.Name + " is working")
}
