package pitcher

import (
	"fmt"

	sqltrace "github.com/DataDog/dd-trace-go/contrib/database/sql"
	sqlxtrace "github.com/DataDog/dd-trace-go/contrib/jmoiron/sqlx"
	"github.com/jmoiron/sqlx"
	pq "github.com/lib/pq"
)

const trackQuery = `SELECT track.gid, rec.gid as recording_id, track.name,
       track.length, track.position, medium.position AS medium_position,
			 album.gid "album.gid", album.name "album.name",
			 artist.gid "artist.gid", artist.name "artist.name", album.id "album.id",
			 release_date.date_year "album.releasedate.date_year",
			 release_date.date_month "album.releasedate.date_month",
			 release_date.date_day "album.releasedate.date_day"
       FROM track JOIN recording AS rec ON (rec.id = track.recording)
			 JOIN artist AS artist ON artist.id = track.artist_credit
       INNER JOIN medium ON medium.id = track.medium
       INNER JOIN release as album ON album.id = medium.release
			 LEFT JOIN LATERAL (SELECT date_year, date_month, date_day FROM release_country WHERE release=album.id) release_date ON true
       WHERE track.gid = :gid`

type trackQueryParams struct {
	GID string `db:"gid"`
}

// CreateDB returns database connection
func CreateDB(config Config) (db *sqlx.DB, err error) {
	connString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable search_path=musicbrainz",
		config.DbHost, config.DbPort, config.DbName, config.DbUser, config.DbPassword)
	sqltrace.Register("postgres", &pq.Driver{}, sqltrace.WithTracer(config.Tracer), sqltrace.WithServiceName("pitcher.db"))
	return sqlxtrace.Open("postgres", connString)
}

// GetTrackData returns Track matching MusicBrainz ID
func GetTrackData(db *sqlx.DB, trackID string) (*Track, error) {
	params := trackQueryParams{
		GID: trackID,
	}

	track := Track{}

	query, err := db.PrepareNamed(trackQuery)
	if err != nil {
		return nil, err
	}

	err = query.Get(&track, params)
	if err != nil {
		return nil, err
	}

	return &track, nil
}
