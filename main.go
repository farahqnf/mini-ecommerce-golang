package main

import (
	"net/http"

	"github.com/tugasmeilyanto/go-trial-class/config"
	"github.com/tugasmeilyanto/go-trial-class/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("Hello World")

	// cli.MainMenu()

	s := gin.Default()
	config.DBConnect()
	s.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello World")
	})

	// List semua produk
	s.GET("/products", controller.HandlerGetProduct)

	// Create order
	s.POST("/orders", controller.HandlerPostOrder)

	// List order
	s.GET("/orders", controller.HandlerGetOrder)

	// cli.LoginMenu()
	s.Run(":8000")
}
