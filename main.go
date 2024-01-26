package main

import (
	"github.com/greenfructose/golangtutorial/album"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.GET("/albums", album.getAlbums)

	router.Run("localhost:8080")
}
