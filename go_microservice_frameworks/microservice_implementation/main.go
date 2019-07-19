package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main()  {
	engine := gin.Default()

	// Keep alive polling endpoint
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Hello World response message JSON formatted endpoint
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	// API call endpoints

	// Get all cars
	engine.GET("/api/cars", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetAllCars())
	})

	// Create a car
	engine.POST("/api/cars", func(c *gin.Context) {
		var car Car
		if c.BindJSON(&car) == nil {
			carID, created := CreateCar(car)
			if created {
				c.Header("Location", "/api/cars" + carID)
				c.Status(http.StatusCreated)
			} else {
				c.Status(http.StatusConflict)
			}
		}
	})

	// Serve static files and templates
	engine.LoadHTMLGlob("./templates/*.html")
	engine.StaticFile("/favicon.ico", "./favicon.ico")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Cloud implemented with Go using the Gin framework"})
	})

	// Get car by carID
	engine.GET("/api/cars/:carid", func(c *gin.Context) {
		carID := c.Params.ByName("carid")
		car, found := GetCar(carID)
		if found {
			c.JSON(http.StatusOK, car)
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	// Update car
	engine.PUT("/api/cars/:carid", func(c *gin.Context) {
		carID := c.Params.ByName("carid")
		var car Car

		if c.BindJSON(&car) == nil {
			exists := UpdateCar(carID, car)
			if exists {
				c.Status(http.StatusOK)
			} else {
				c.Status(http.StatusNotFound)
			}
		}
	})

	// Delete car
	engine.DELETE("/api/cars/:carid", func(c *gin.Context) {
		carID := c.Params.ByName("carid")
		deleted := DeleteCar(carID)
		if deleted {
			c.Status(http.StatusOK)
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	// Run the server on PORT
	engine.Run(port())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}