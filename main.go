package main

import (
	"imagescdn/handlers"

	"github.com/gin-gonic/gin"
)

// http://localhost:8080/v1/images?url=dGVzdGluZy5qcGVn
// http://localhost:8080/v1/images?url=c2FtcGxlX21lZGlhX2ZpbGUucG5n

func main() {
	router := gin.Default()
	router.GET("v1/images", handlers.GetImageData)
	router.Run("localhost:8080")
}
