package service

import (
	"context"

	"github.com/azdaev/fitra/api/cmd/main/internal/audiolectures/model"
	"github.com/pkg/errors"
)

type audioLecturesRepo interface {
	GetAllAudioLectures(ctx context.Context) ([]model.AudioLecture, error)
	GetAudioLecture(ctx context.Context, id int) (model.AudioLecture, error)
}

func New(audioLecturesRepo audioLecturesRepo) *Service {
	return &Service{
		audioLecturesRepo: audioLecturesRepo,
	}
}

type Service struct {
	audioLecturesRepo audioLecturesRepo
}

func (s *Service) GetAllAudioLectures(ctx context.Context) ([]model.AudioLecture, error) {
	audioLectures, err := s.audioLecturesRepo.GetAllAudioLectures(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "audioLecturesRepo.GetAllAudioLectures")
	}

	return audioLectures, nil
}

func (s *Service) GetAudioLecture(ctx context.Context, id int) (model.AudioLecture, error) {
	audioLecture, err := s.audioLecturesRepo.GetAudioLecture(ctx, id)
	if err != nil {
		return model.AudioLecture{}, errors.Wrap(err, "audioLecturesRepo.GetAudioLecture")
	}

	return audioLecture, nil
}
