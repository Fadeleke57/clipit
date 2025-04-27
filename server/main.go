package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Fadeleke57/clipit/server/controller"
	"github.com/Fadeleke57/clipit/server/service"
	"github.com/Fadeleke57/clipit/server/entity"
)

var(
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)
func main() {
	server := gin.Default()

	server.GET("/posts", func(ctx *gin.Context){
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/posts", func(ctx *gin.Context){
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to CLIPIT!!",
		})
	})

	server.Run(":8000")
}