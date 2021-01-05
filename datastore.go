package pitcher

import (
	"context"
	"fmt"

	pb "github.com/hugoatease/pitcher/protobuf"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const trackQuery = `SELECT track.gid, rec.gid as recording_id, track.name,
       track.length, track.position, medium.position AS medium_position,
			 album.gid "album.gid", album.name "album.name",
			 artist.gid "artist.gid", artist.name "artist.name",
			 album.id "album.id", release.gid "album.gid",
			 COALESCE(release_date.year, -1) "album.release_date.year",
			 COALESCE(release_date.month, -1) "album.release_date.month",
			 COALESCE(release_date.day, -1) "album.release_date.day"
       FROM track JOIN recording AS rec ON (rec.id = track.recording)
			 JOIN artist_credit_name AS artist_credit_name ON artist_credit_name.artist_credit = track.artist_credit
			 JOIN artist AS artist ON artist.id = artist_credit_name.artist
       JOIN medium ON medium.id = track.medium
       JOIN release as release ON release.id = medium.release
			 JOIN release_group AS album ON album.id = release.release_group
			 LEFT JOIN LATERAL (SELECT date_year AS year, date_month AS month, date_day AS day FROM release_country WHERE release=release.id) release_date ON true
       WHERE track.gid = :gid
			 ORDER BY artist_credit_name.position LIMIT 1`

type trackQueryParams struct {
	GID string `db:"gid"`
}

const tracksQuery = `SELECT track.gid, rec.gid as recording_id, track.name,
       track.length, track.position, medium.position AS medium_position,
			 album.gid "album.gid", album.name "album.name",
			 artist.gid "artist.gid", artist.name "artist.name",
			 album.id "album.id", release.gid "album.gid",
			 COALESCE(release_date.year, -1) "album.release_date.year",
			 COALESCE(release_date.month, -1) "album.release_date.month",
			 COALESCE(release_date.day, -1) "album.release_date.day"
       FROM track JOIN recording AS rec ON (rec.id = track.recording)
			 JOIN artist_credit_name AS artist_credit_name ON artist_credit_name.artist_credit = track.artist_credit
			 JOIN artist AS artist ON artist.id = artist_credit_name.artist
       JOIN medium ON medium.id = track.medium
       JOIN release as release ON release.id = medium.release
			 JOIN release_group AS album ON album.id = release.release_group
			 LEFT JOIN LATERAL (SELECT date_year AS year, date_month AS month, date_day AS day FROM release_country WHERE release=release.id) release_date ON true
       WHERE track.gid IN (?)
			 ORDER BY artist_credit_name.position`

const coverQuery = `SELECT listing.id AS id, release.gid AS release_mbid,
        listing.is_front AS is_front, listing.is_back AS is_back,
				listing.mime_type AS mime_type, image_type.suffix AS suffix
				FROM musicbrainz.release as release
				JOIN cover_art_archive.cover_art as coverart on (coverart.release=release.id)
				JOIN cover_art_archive.index_listing as listing ON (coverart.id=listing.id)
				JOIN cover_art_archive.image_type as image_type ON (image_type.mime_type=listing.mime_type)
				WHERE release.gid = :gid AND is_front=true ORDER BY ordering LIMIT 1`

type coverQueryParams struct {
	GID string `db:"gid"`
}

// CreateDB returns database connection
func CreateDB(config Config) (db *sqlx.DB, err error) {
	connString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable search_path=musicbrainz",
		config.DbHost, config.DbPort, config.DbName, config.DbUser, config.DbPassword)
	return sqlx.Open("postgres", connString)
}

// GetTrackData returns Track matching MusicBrainz ID
func GetTrackData(ctx context.Context, db *sqlx.DB, trackID string) (*pb.Track, error) {
	params := map[string]interface{}{
		"gid": trackID,
	}

	track := pb.Track{}

	query, args, err := sqlx.Named(trackQuery, params)
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)

	err = db.Get(&track, query, args...)
	if err != nil {
		return nil, err
	}

	return &track, nil
}

// GetTracksData returns Tracks matching a list of MusicBrainz IDs
func GetTracksData(ctx context.Context, db *sqlx.DB, trackIDs []string) ([]*pb.Track, error) {
	tracks := []*pb.Track{}

	query, args, err := sqlx.In(tracksQuery, trackIDs)
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)

	err = db.Select(&tracks, query, args...)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

// GetReleaseImageData returns image data for releaseID
func GetReleaseImageData(db *sqlx.DB, releaseID string) (*CoverArtListing, error) {
	params := coverQueryParams{
		GID: releaseID,
	}

	listing := CoverArtListing{}

	query, err := db.Preparex(coverQuery)
	if err != nil {
		return nil, err
	}

	err = query.Get(&listing, params)
	if err != nil {
		return nil, err
	}

	return &listing, nil
}
