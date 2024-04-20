package repo

type AudioLecture struct {
	ID           int    `db:"id"`
	Title        string `db:"title"`
	AuthorName   string `db:"author-name"`
	AuthorID     int    `db:"author-id"`
	PlaylistName string `db:"playlist-name"`
	PlaylistID   int    `db:"playlist-id"`
	AudioURL     string `db:"audio-url"`
	LengthSec    int    `db:"length-sec"`
}
