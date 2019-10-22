package controller

import (
	"net/http"
	"net/url"

	"github.com/cortago/cortago/controller"
)

type SecondController struct {
	controller.Controller
}

func (c *SecondController) Index(w http.ResponseWriter, r *http.Request, params url.Values) {
	c.Initiliaze(w, r)
	c.RenderText(c.Name + " index is working")
	c.Terminate()
}

func (c *SecondController) List(w http.ResponseWriter, r *http.Request, params url.Values) {
	c.Initiliaze(w, r)
	c.RenderText(c.Name + " list is working")
	c.Terminate()
}
