package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-example-jiw/internal/config"
	"github.com/go-example-jiw/internal/pointer"
)

func main() {
	fmt.Println("License: MIT LICENSE")
	fmt.Println("(c) 2022")
	fmt.Println()

	r := gin.Default()

	d := pointer.PointerExample{}
	d.Load()
	log.Println(d)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": "go-example-echo",
			"copyright":   "(c) 2022",
		})
	})

	err := r.Run(config.PublishAddress)
	if err != nil {
		log.Fatal(err)
	}
}
