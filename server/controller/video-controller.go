pacakge controller

import (
	"github.com/Fadeleke57/clipit/server/entity"
	"github.com/Fadeleke57/clipit/server/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller {
		service: service,
	}
}

func (c *controller ) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

