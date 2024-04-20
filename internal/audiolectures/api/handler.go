package api

import (
	"context"

	"github.com/azdaev/fitra/api/cmd/main/internal/audiolectures/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type service interface {
	GetAllAudioLectures(ctx context.Context) ([]model.AudioLecture, error)
	GetAudioLecture(ctx context.Context, id int) (model.AudioLecture, error)
}

type Handler struct {
	service service
	log     *logrus.Logger
}

func New(service service, log *logrus.Logger) *gin.Engine {
	h := &Handler{
		service: service,
		log:     log,
	}

	r := gin.New()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api/v1/audiolectures/")

	api.GET("", h.GetAllAudioLectures)
	api.GET("/:id", h.GetAudioLecture)

	return r
}
