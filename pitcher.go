package pitcher

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type App struct {
	DB *sqlx.DB
}

func NewApp() (*App, error) {
	db, err := CreateDB()
	if err != nil {
		return nil, err
	}

	app := &App{
		DB: db,
	}

	return app, nil
}

func (app *App) Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/track/{trackID}", app.TrackHandler)
	http.ListenAndServe(":5000", r)
}
