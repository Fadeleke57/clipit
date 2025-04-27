package service

import "https://github.com/Fadeleke57/clipit/server/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []eneity.Video
}

type videoService struct {
	videos [] entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}