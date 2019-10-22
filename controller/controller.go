package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

// Controller is basic
type Controller struct {
	Name string

	Response http.ResponseWriter
	Request  *http.Request

	Directory string
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

// RenderHTML is used to render render HTML files
func (c *Controller) RenderHTML(name string, t interface{}) {
	tmpl, err := template.ParseGlob(path.Join(c.Directory, "*"))
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(c.Response, name, t)
	if err != nil {
		fmt.Println(err)
	}

}

// Terminate a specific controller
func (c *Controller) Terminate() {
	c.Request = nil
	c.Response = nil
}
