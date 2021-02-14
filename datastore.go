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
			 artist.gid "artist.gid", artist.name "artist.name", album.id "album.id",
			 COALESCE(release_date.year, -1) "album.release_date.year",
			 COALESCE(release_date.month, -1) "album.release_date.month",
			 COALESCE(release_date.day, -1) "album.release_date.day"
       FROM track JOIN recording AS rec ON (rec.id = track.recording)
			 JOIN artist_credit_name AS artist_credit_name ON artist_credit_name.artist_credit = track.artist_credit
			 JOIN artist AS artist ON artist.id = artist_credit_name.artist
       JOIN medium ON medium.id = track.medium
       JOIN release as release ON release.id = medium.release
			 JOIN release_group AS album ON album.id = release.release_group
			 LEFT JOIN LATERAL (SELECT date_year AS year, date_month AS month, date_day AS day FROM release_country WHERE release=release.id LIMIT 1) release_date ON true
       WHERE track.gid = :gid AND artist_credit_name.position = 0`

const tracksQuery = `SELECT track.gid, rec.gid as recording_id, track.name,
       track.length, track.position, medium.position AS medium_position,
			 album.gid "album.gid", album.name "album.name",
			 artist.gid "artist.gid", artist.name "artist.name", album.id "album.id",
			 COALESCE(release_date.year, -1) "album.release_date.year",
			 COALESCE(release_date.month, -1) "album.release_date.month",
			 COALESCE(release_date.day, -1) "album.release_date.day"
       FROM track JOIN recording AS rec ON (rec.id = track.recording)
			 JOIN artist_credit_name AS artist_credit_name ON artist_credit_name.artist_credit = track.artist_credit
			 JOIN artist AS artist ON artist.id = artist_credit_name.artist
       JOIN medium ON medium.id = track.medium
       JOIN release as release ON release.id = medium.release
			 JOIN release_group AS album ON album.id = release.release_group
			 LEFT JOIN LATERAL (SELECT date_year AS year, date_month AS month, date_day AS day FROM release_country WHERE release=release.id LIMIT 1) release_date ON true
       WHERE track.gid IN (?) AND artist_credit_name.position = 0`

const preferredCoverReleaseQuery = `SELECT DISTINCT ON (release.release_group)
          release.gid AS mbid
        FROM cover_art_archive.index_listing
        JOIN musicbrainz.release
          ON musicbrainz.release.id = cover_art_archive.index_listing.release
        JOIN musicbrainz.release_group
          ON release_group.id = release.release_group
        LEFT JOIN (
          SELECT release, date_year, date_month, date_day
          FROM musicbrainz.release_country
          UNION ALL
          SELECT release, date_year, date_month, date_day
          FROM musicbrainz.release_unknown_country
        ) release_event ON (release_event.release = release.id)
        FULL OUTER JOIN cover_art_archive.release_group_cover_art
        ON release_group_cover_art.release = musicbrainz.release.id
        WHERE release_group.gid = :gid
        AND is_front = true
        ORDER BY release.release_group, release_group_cover_art.release,
          release_event.date_year, release_event.date_month,
					release_event.date_day`

const coverFileInfoQuery = `SELECT index_listing.id, image_type.suffix
              FROM cover_art_archive.index_listing
              JOIN musicbrainz.release
                ON cover_art_archive.index_listing.release = musicbrainz.release.id
              JOIN cover_art_archive.image_type
                ON cover_art_archive.index_listing.mime_type = cover_art_archive.image_type.mime_type
             WHERE musicbrainz.release.gid = :gid
               AND is_front = true
					ORDER BY ordering ASC LIMIT 1`

const artistQuery = `SELECT id, gid, name
							FROM artist
							WHERE artist.gid = :gid`

const artistsQuery = `SELECT id, gid, name
							FROM artist
							WHERE artist.gid IN (?)`

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

// GetCoverFileInfoByReleaseGroup returns image data for releaseGroupID
func GetCoverFileInfoByReleaseGroup(ctx context.Context, db *sqlx.DB, releaseGroupID string) (*CoverFileInfo, error) {
	params := map[string]interface{}{
		"gid": releaseGroupID,
	}

	preferredCoverRelease := PreferredCoverRelease{}

	query, args, err := sqlx.Named(preferredCoverReleaseQuery, params)
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)

	err = db.Get(&preferredCoverRelease, query, args...)
	if err != nil {
		return nil, err
	}

	infoQueryParams := map[string]interface{}{
		"gid": preferredCoverRelease.MbID,
	}

	coverFileInfo := CoverFileInfo{
		ReleaseMbID: preferredCoverRelease.MbID,
	}

	query, args, err = sqlx.Named(coverFileInfoQuery, infoQueryParams)
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)

	err = db.Get(&coverFileInfo, query, args...)
	if err != nil {
		return nil, err
	}
	return &coverFileInfo, nil
}

// GetArtistData returns Artist matching MusicBrainz ID
func GetArtistData(ctx context.Context, db *sqlx.DB, artistID string) (*pb.Artist, error) {
	params := map[string]interface{}{
		"gid": artistID,
	}

	artist := pb.Artist{}

	query, args, err := sqlx.Named(artistQuery, params)
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)

	err = db.Get(&artist, query, args...)
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

// GetArtistsData returns Artists matching a list of MusicBrainz IDs
func GetArtistsData(ctx context.Context, db *sqlx.DB, artistIDs []string) ([]*pb.Artist, error) {
	artists := []*pb.Artist{}

	query, args, err := sqlx.In(artistsQuery, artistIDs)
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)

	err = db.Select(&artists, query, args...)
	if err != nil {
		return nil, err
	}

	return artists, nil
}
