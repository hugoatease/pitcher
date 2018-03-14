package pitcher

// Artist structure
type Artist struct {
	GID  string `db:"gid" json:"mbid"`
	Name string `db:"name"json:"name"`
}

// ReleaseDate structure
type ReleaseDate struct {
	Year  string `db:"date_year" json:"year"`
	Month string `db:"date_month" json:"month"`
	Day   string `db:"date_day" json:"day"`
}

// Album structure
type Album struct {
	ID          int          `db:"id" json:"-"`
	GID         string       `db:"gid" json:"mbid"`
	Name        string       `db:"name" json:"name"`
	ReleaseDate *ReleaseDate `json:"release_date"`
}

// Track structure
type Track struct {
	GID            string  `db:"gid" json:"mbid"`
	RecordingID    string  `db:"recording_id" json:"recording_mbid"`
	Name           string  `db:"name" json:"name"`
	MediumPosition int     `db:"medium_position" json:"medium_position"`
	Position       int     `db:"position" json:"position"`
	Length         int     `db:"length" json:"length"`
	Artist         *Artist `json:"artist"`
	Album          *Album  `json:"album"`
}
