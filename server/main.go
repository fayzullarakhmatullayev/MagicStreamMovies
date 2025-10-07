package main

import (
	"fmt"

	"github.com/fayzullarakhmatullayev/MagicStreamMovies/server/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, GO!")
	})

	router.GET("/movies", controllers.GetMovies())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
