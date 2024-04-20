package api

import (
	"net/http"
	"strconv"

	"github.com/azdaev/template/internal/audiolectures/api/responses"
	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAudioLecture(ctx *gin.Context) {
	idRaw := ctx.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		h.log.Error(errors.Wrap(err, "strconv.Atoi failed to parse audiolecture id"))
		ctx.JSON(http.StatusBadRequest, "invalid audiolecture id")
		return
	}

	audioLectureDB, err := h.service.GetAudioLecture(ctx, id)
	if err != nil {
		h.log.Error(errors.Wrap(err, "service.GetAudioLecture"))
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := responses.AudioLecture{
		ID:           audioLectureDB.ID,
		Title:        audioLectureDB.Title,
		AuthorName:   audioLectureDB.AuthorName,
		AuthorID:     audioLectureDB.AuthorID,
		PlaylistName: audioLectureDB.PlaylistName,
		PlaylistID:   audioLectureDB.PlaylistID,
		AudioURL:     audioLectureDB.AudioURL,
		LengthSec:    audioLectureDB.LengthSec,
	}

	ctx.JSON(http.StatusOK, rsp)
}
