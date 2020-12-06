package pitcher

import (
	"database/sql"
	"encoding/json"
)

// NullInt64 that can be marshalled with null value
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON that returns null-values on NULL sql columns
func (r NullInt64) MarshalJSON() ([]byte, error) {
	if r.Valid {
		return json.Marshal(r.Int64)
	}
	return json.Marshal(nil)
}

// Artist structure
type Artist struct {
	GID  string `db:"gid" json:"mbid"`
	Name string `db:"name"json:"name"`
}

// ReleaseDate structure
type ReleaseDate struct {
	Year  NullInt64 `db:"date_year" json:"year"`
	Month NullInt64 `db:"date_month" json:"month"`
	Day   NullInt64 `db:"date_day" json:"day"`
}

// Album structure
type Album struct {
	ID          int          `db:"id" json:"-"`
	GID         string       `db:"gid" json:"mbid"`
	ReleaseGID  string       `db:"release_gid" json:"release_mbid"`
	Name        string       `db:"name" json:"name"`
	ReleaseDate *ReleaseDate `json:"release_date"`
}

// Track structure
type Track struct {
	ID             int64   `db:"id" json:"id"`
	GID            string  `db:"gid" json:"mbid"`
	RecordingID    string  `db:"recording_id" json:"recording_mbid"`
	Name           string  `db:"name" json:"name"`
	MediumPosition int     `db:"medium_position" json:"medium_position"`
	Position       int     `db:"position" json:"position"`
	Length         int     `db:"length" json:"length"`
	Artist         *Artist `json:"artist"`
	Album          *Album  `json:"album"`
}

// CoverArtListing structure
type CoverArtListing struct {
	ID          int    `db:"id" json:"id"`
	ReleaseMbid string `db:"release_mbid" json:"release_mbid"`
	IsFront     bool   `db:"is_front" json:"is_front"`
	IsBack      bool   `db:"is_back" json:"is_back"`
	MimeType    string `db:"mime_type" json:"mime_type"`
	Suffix      string `db:"suffix" json:"suffix"`
}

type SolrDoc struct {
	GID string `json:"gid"`
}

type SolrResponseBody struct {
	Docs []SolrDoc `json:"docs"`
}

type SolrResponse struct {
	Response SolrResponseBody `json:"response"`
}
