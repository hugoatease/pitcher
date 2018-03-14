package pitcher

// Artist structure
type Artist struct {
	GID  string `db:"gid"`
	Name string `db:"name"`
}

// ReleaseDate structure
type ReleaseDate struct {
	Year  string `db:"date_year"`
	Month string `db:"date_month"`
	Day   string `db:"date_day"`
}

// Album structure
type Album struct {
	ID          string `db:"id"`
	GID         string `db:"gid"`
	Name        string `db:"name"`
	ReleaseDate *ReleaseDate
}

// Track structure
type Track struct {
	GID            string `db:"gid"`
	RecordingID    string `db:"recording_id"`
	Name           string `db:"name"`
	MediumPosition int    `db:"medium_position"`
	Position       int    `db:"position"`
	Length         int    `db:"length"`
	Artist         *Artist
	Album          *Album
}
