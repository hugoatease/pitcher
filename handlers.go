package pitcher

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TrackHandler returns data abouy a MusicBrainz track
func (app *App) TrackHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	trackID := vars["trackID"]

	tracks, err := GetTrackData(app.DB, trackID)
	if err != nil {
		log.Print("Error fetching track data", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(tracks)
	if err != nil {
		log.Print("Error marshalling")
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Write(result)
}
