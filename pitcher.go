package pitcher

import (
	"log"
	"net/http"

	muxtrace "github.com/DataDog/dd-trace-go/contrib/gorilla/mux"
	"github.com/DataDog/dd-trace-go/tracer"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Bind       string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	Tracer     *tracer.Tracer
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
	r := muxtrace.NewRouter(muxtrace.WithTracer(app.Config.Tracer),
		muxtrace.WithServiceName("pitcher.web"))
	r.HandleFunc("/tracks/{trackID}", app.TrackHandler)
	r.HandleFunc("/releases/{releaseID}/image", app.ReleaseImageHandler)
	log.Print("serving on ", app.Config.Bind)
	log.Fatal(http.ListenAndServe(app.Config.Bind, r))
}
