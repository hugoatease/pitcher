package pitcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"

	pb "github.com/hugoatease/pitcher/protobuf"
)

type Config struct {
	Bind       string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	SolrURL    string
}

type PitcherServer struct {
	pb.UnimplementedPitcherServer
	Config
	DB *sqlx.DB
}

func NewServer(config Config) (*PitcherServer, error) {
	db, err := CreateDB(config)
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)

	if err != nil {
		return nil, err
	}

	app := &PitcherServer{
		DB:     db,
		Config: config,
	}

	return app, nil
}

func (s *PitcherServer) MatchTrack(ctx context.Context, request *pb.MatchingRequest) (*pb.MatchingResponse, error) {
	var query string
	baseQuery := fmt.Sprintf("track_name:(%s) AND artist_name:(%s)", request.TrackName, request.ArtistName)
	query = baseQuery
	if request.ReleaseName != "" {
		query = query + fmt.Sprintf(" AND release_name:(%s)", request.ReleaseName)
	}

	httpQuery := url.Values{}
	httpQuery.Set("defType", "edismax")
	httpQuery.Set("bq", "release_group_secondary_type:\"-1\"^10 release_group_type:1^10 status:1^10")
	httpQuery.Set("fl", "gid")
	httpQuery.Set("rows", "1")
	httpQuery.Set("q", query)

	solrReqURL, err := url.Parse(s.Config.SolrURL)
	if err != nil {
		return nil, err
	}

	solrReqURL.Path = "/solr/pitcher.track-info/select"
	solrReqURL.RawQuery = httpQuery.Encode()

	solrRes, err := http.Get(solrReqURL.String())
	if err != nil {
		return nil, err
	}

	defer solrRes.Body.Close()
	var parsedSolrRes SolrResponse
	err = json.NewDecoder(solrRes.Body).Decode(&parsedSolrRes)

	if err != nil {
		return nil, err
	}

	if len(parsedSolrRes.Response.Docs) == 0 {
		return nil, errors.New("No track found")
	}

	track, err := GetTrackData(ctx, s.DB, parsedSolrRes.Response.Docs[0].GID)
	if err != nil {
		return nil, err
	}

	response := pb.MatchingResponse{
		Track: track,
	}

	return &response, nil
}

func (s *PitcherServer) GetTrack(ctx context.Context, request *pb.TrackRequest) (*pb.TrackResponse, error) {
	track, err := GetTrackData(ctx, s.DB, request.Gid)
	if err != nil {
		return nil, err
	}

	response := pb.TrackResponse{
		Track: track,
	}

	return &response, nil
}

func (s *PitcherServer) GetTracks(ctx context.Context, request *pb.TracksRequest) (*pb.TracksResponse, error) {
	tracks, err := GetTracksData(ctx, s.DB, request.Gids)
	if err != nil {
		return nil, err
	}

	response := pb.TracksResponse{
		Tracks: tracks,
	}

	return &response, nil
}
