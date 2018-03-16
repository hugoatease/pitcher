package pitcher

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TrackHandler returns data about a MusicBrainz track
func (app *App) TrackHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	trackID := vars["trackID"]

	track, err := GetTrackData(app.DB, trackID)
	if err != nil {
		log.Print("Error fetching track data", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(track)
	if err != nil {
		log.Print("Error marshalling")
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Write(result)
}

// ReleaseImageHandler returns data about a MusicBrainz release image
func (app *App) ReleaseImageHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	releaseID := vars["releaseID"]

	log.Print(releaseID)

	listing, err := GetReleaseImageData(app.DB, releaseID)
	if err != nil {
		log.Print("Error fetching data", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(listing)
	if err != nil {
		log.Print("Error marshalling")
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Write(result)
}
