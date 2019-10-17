package main

import (
	"cortago/examples/controller"

	"github.com/cortago/cortago/router"
)

func main() {
	router.Port = 8000
	main := router.New()
	mainController := controller.MainController{}
	main.Handle("/index", mainController.Index)

	second := router.New()
	secondController := controller.SecondController{}
	second.Handle("/index", secondController.Index)
	second.Handle("/list", secondController.List)

	y := router.New()
	y.AppendRouter("/main", main)
	y.AppendRouter("/second", second)

	y.StartApp()
}
