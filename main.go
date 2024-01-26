package main

import (
	"github.com/gin-gonic/gin"
	"github.com/greenfructose/golangtutorial/items"
)


func main() {
	router := gin.Default()
	router.GET("/albums", items.getAlbums)

	router.Run("localhost:8080")
}
