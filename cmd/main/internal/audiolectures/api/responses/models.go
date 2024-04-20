package responses

type AudioLecture struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	AuthorName   string `json:"author_name"`
	AuthorID     int    `json:"author_id"`
	PlaylistName string `json:"playlist_name"`
	PlaylistID   int    `json:"playlist_id"`
	AudioURL     string `json:"audio_url"`
	LengthSec    int    `json:"length_sec"`
}
