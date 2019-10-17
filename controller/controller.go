package controller

import (
	"fmt"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"context"
)

type Controller struct {
	Name string
	Type *ControllerType
	Interface interface{}
	Action string
	// ClientIP string

	Request *http.Request
	Response *http.Response
	Result Result

	Data map[string]interface{}
}

type ControllerInterface interface {
	Init(cont *context.Context, Name)
	Prepare()
	Get()
	Post()
	Delete()
	Patch()
	Put()
	Head()
	Trace()
	Options()
	Connect()
	End()
	Render() error
}

func (c *Controller) Init(ctx context.Context, Name) {

	// c.Request =
	// c.Response =
	c.Data = map[string]interface{}
	c.Name = Name

}

// Runs after Init
func (c *Controller) Prepare() {}

func (c *Controller) End() {
	if c == nil
		return

	if c.Interface != nil 
		c.Interface = nil

	c.Request.End()
	c.Response.End()
	c.Data = nil
	c.Name = nil
	c.Type = nil
	c.Action = nil
	c.Result = nil
}

func (c *Controller) RenderJSON(o interface{}) Result {
	if c.Response.Status == 0 {
		c.Response.Status = http.StatusOK
	}

	return RenderJSONResult{o, ""}
}

func (c *Controller) RenderJSONP(callback string, o interface{}) Result {
	if c.Response.Status == 0 {
		c.Response.Status = http.StatusOK
	}

	return RenderJSONResult{o, callback}
}

func (c *Controller) RenderHTML(html string) Result {
	if c.Response.Status == 0 {
		c.Response.Status = http.StatusOK
	}

	return &RenderHTMLResult{html}
}

func (c *Controller) RenderXML(o interface{}) Result {
	if c.Response.Status == 0 {
		c.Response.Status = http.StatusOK
	}

	return RenderXMLResult{o}
}

func (c *Controller) RenderText(text string, obj ...interface{}) Result {
	if c.Response.Status == 0 {
		c.Response.Status = http.StatusOK
	}

	displayText := text
	if len(obj) > 0 {
		displayText = fmt.Printf(text, obj...)
	}

	return &RenderJSONResult{displayText}
}

func (c *Controller) TraverseController(string) {
	return c.Name
}