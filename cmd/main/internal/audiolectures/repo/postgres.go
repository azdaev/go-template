package repo

import (
	"context"
	"database/sql"

	"github.com/azdaev/fitra/api/cmd/main/internal/audiolectures/model"
	"github.com/pkg/errors"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAllAudioLectures(ctx context.Context) ([]model.AudioLecture, error) {
	query := "SELECT id, title, author_name, author_id, playlist_name, playlist_id, audio_url, length_sec FROM audio_lecture;"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "query context")
	}
	defer rows.Close()

	audioLectures := make([]model.AudioLecture, 0)
	for rows.Next() {
		audioLectureDB := AudioLecture{}
		if err = rows.Scan(
			&audioLectureDB.ID,
			&audioLectureDB.Title,
			&audioLectureDB.AuthorName,
			&audioLectureDB.AuthorID,
			&audioLectureDB.PlaylistName,
			&audioLectureDB.PlaylistID,
			&audioLectureDB.AudioURL,
			&audioLectureDB.LengthSec,
		); err != nil {
			return nil, errors.Wrap(err, "scan")
		}

		audioLectures = append(audioLectures, model.AudioLecture{
			ID:           audioLectureDB.ID,
			Title:        audioLectureDB.Title,
			AuthorName:   audioLectureDB.AuthorName,
			AuthorID:     audioLectureDB.AuthorID,
			PlaylistName: audioLectureDB.PlaylistName,
			PlaylistID:   audioLectureDB.PlaylistID,
			AudioURL:     audioLectureDB.AudioURL,
			LengthSec:    audioLectureDB.LengthSec,
		})
	}

	return audioLectures, nil
}

func (r *Repository) GetAudioLecture(ctx context.Context, id int) (model.AudioLecture, error) {
	query := "SELECT id, title, author_name, author_id, playlist_name, playlist_id, audio_url, length_sec FROM audio_lecture WHERE id = $1;"

	row := r.db.QueryRowContext(ctx, query, id)
	audioLectureDB := AudioLecture{}
	err := row.Scan(
		&audioLectureDB.ID,
		&audioLectureDB.Title,
		&audioLectureDB.AuthorName,
		&audioLectureDB.AuthorID,
		&audioLectureDB.PlaylistName,
		&audioLectureDB.PlaylistID,
		&audioLectureDB.AudioURL,
		&audioLectureDB.LengthSec,
	)
	if err != nil {
		return model.AudioLecture{}, errors.Wrap(err, "scan")
	}

	return model.AudioLecture{
		ID:           audioLectureDB.ID,
		Title:        audioLectureDB.Title,
		AuthorName:   audioLectureDB.AuthorName,
		AuthorID:     audioLectureDB.AuthorID,
		PlaylistName: audioLectureDB.PlaylistName,
		PlaylistID:   audioLectureDB.PlaylistID,
		AudioURL:     audioLectureDB.AudioURL,
		LengthSec:    audioLectureDB.LengthSec,
	}, nil
}
