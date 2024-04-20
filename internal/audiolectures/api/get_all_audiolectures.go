package api

import (
	"net/http"

	"github.com/azdaev/template/internal/audiolectures/api/responses"
	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllAudioLectures(ctx *gin.Context) {
	audioLectures, err := h.service.GetAllAudioLectures(ctx)

	if err != nil {
		h.log.Error(errors.Wrap(err, "failed to get all audio lectures"))
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := make([]responses.AudioLecture, 0, len(audioLectures))

	for _, lecture := range audioLectures {
		rsp = append(rsp, responses.AudioLecture{
			ID:           lecture.ID,
			Title:        lecture.Title,
			AuthorName:   lecture.AuthorName,
			AuthorID:     lecture.AuthorID,
			PlaylistName: lecture.PlaylistName,
			PlaylistID:   lecture.PlaylistID,
			AudioURL:     lecture.AudioURL,
			LengthSec:    lecture.LengthSec,
		})
	}

	ctx.JSON(http.StatusOK, rsp)
}
