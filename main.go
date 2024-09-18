package main

import (
	"log"

	"national-weather-service-api/dependancies"
	_ "national-weather-service-api/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			National Weather Service API
// @version		1.0
// @description	This is a simple API designed to call the official National Weather Service API
func main() {
	d := dependancies.Initilize()
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	forcastController := d.GetForcastController()
	f := e.Group("/forcast")
	f.GET("/get/:lat/:long", forcastController.GetForcast)

	log.Fatal(e.Start(":9001"))
}
