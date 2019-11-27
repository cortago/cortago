package main

import (
	"cortago/examples/controller"

	"cortago/router"
)

func main() {
	router.Port = 8000
	main := router.New()
	mainController := controller.MainController{}
	mainController.Name = "Main Controller"
	main.Handle("/index", mainController.Index)

	second := router.New()
	secondController := controller.SecondController{}
	secondController.Name = "Second Controller"
	second.Handle("/index", secondController.Index)
	second.Handle("/list", secondController.List)

	y := router.New()
	y.FileServer("/assets")
	y.AppendRouter("/main", main)
	y.AppendRouter("/second", second)

	y.StartApp()
}
