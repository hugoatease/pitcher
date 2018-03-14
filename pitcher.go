package pitcher

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Bind       string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
}

type App struct {
	DB     *sqlx.DB
	Config Config
}

func NewApp(config Config) (*App, error) {
	db, err := CreateDB(config)
	if err != nil {
		return nil, err
	}

	app := &App{
		DB:     db,
		Config: config,
	}

	return app, nil
}

func (app *App) Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/tracks/{trackID}", app.TrackHandler)
	log.Print("serving on ", app.Config.Bind)
	log.Fatal(http.ListenAndServe(app.Config.Bind, r))
}
