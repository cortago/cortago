package cortago

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
	ClientIP string

	Request *Request
	Response *Response
	Result Result

	Data map[string]interface{}
}

type ControllerInterface interface {
	Init()
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

func (c *Controller) Init(ctx contextContext) {

	// c.Request.SetRequest(context.GetRequest())
}

// Runs after Init
func (c *Controller) Prepare() {}

func (c *Controller) End() {}

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