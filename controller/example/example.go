// package main
package controller

import (
	"github.com/cortago/cortago/controller"
)

// A controller named first is defined
type first struct { 
	cortago.Controller
}

func (c Controller) Init cortago.Result {
	// return c.Render()
}