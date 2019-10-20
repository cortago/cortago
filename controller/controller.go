package controller

import (
	"fmt"
	"net/http"
)

// Controller is basic
type Controller struct {
	Name string

	Response http.ResponseWriter
	Request  *http.Request
}

// Initiliaze values for specific controller
func (c *Controller) Initiliaze(response http.ResponseWriter, request *http.Request) {
	c.Request = request
	c.Response = response
}

// RenderText is used to render a text of value
func (c *Controller) RenderText(value string) {
	fmt.Fprint(c.Response, value)
}

// Terminate a specific controller
func (c *Controller) Terminate() {
	c.Request = nil
	c.Response = nil
}
