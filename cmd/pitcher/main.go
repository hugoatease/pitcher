package main

import (
	"log"
	"os"

	"github.com/DataDog/dd-trace-go/tracer"
	"github.com/hugoatease/pitcher"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Pitcher"
	app.Usage = "Musicbrainz database explorer"
	app.Version = "0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "bind",
			Value:  ":5000",
			Usage:  "host and port to bind to",
			EnvVar: "PITCHER_BIND",
		},
		cli.StringFlag{
			Name:   "dbhost",
			Value:  "localhost",
			Usage:  "musicbrainz's postgresql database hostname",
			EnvVar: "PITCHER_DBHOST",
		},
		cli.StringFlag{
			Name:   "dbport",
			Value:  "5432",
			Usage:  "musicbrainz's postgresql database port",
			EnvVar: "PITCHER_DBPORT",
		},
		cli.StringFlag{
			Name:   "dbname",
			Value:  "musicbrainz",
			Usage:  "musicbrainz's postgresql database name",
			EnvVar: "PITCHER_DBNAME",
		},
		cli.StringFlag{
			Name:   "dbuser",
			Value:  "musicbrainz",
			Usage:  "musicbrainz's postgresql database name",
			EnvVar: "PITCHER_DBUSER",
		},
		cli.StringFlag{
			Name:   "dbpassword",
			Value:  "musicbrainz",
			Usage:  "musicbrainz's postgresql database password",
			EnvVar: "PITCHER_DBPASSWORD",
		},
		cli.BoolFlag{
			Name:   "tracing",
			EnvVar: "PITCHER_TRACING",
		},
		cli.StringFlag{
			Name:   "ddhost",
			Value:  "localhost",
			Usage:  "hostname of the DataDog tracing agent",
			EnvVar: "PITCHER_DATADOG_HOST",
		},
	}

	app.Action = func(c *cli.Context) error {
		config := pitcher.Config{
			Bind:       c.String("bind"),
			DbHost:     c.String("dbhost"),
			DbPort:     c.String("dbport"),
			DbName:     c.String("dbname"),
			DbUser:     c.String("dbuser"),
			DbPassword: c.String("dbpassword"),
		}

		config.Tracer = tracer.DefaultTracer

		if c.Bool("tracing") {
			tracerTransport := tracer.NewTransport(c.String("ddhost"), "8126")
			config.Tracer = tracer.NewTracerTransport(tracerTransport)
		} else {
			config.Tracer = tracer.NewTracer()
			config.Tracer.SetEnabled(false)
		}

		server, err := pitcher.NewApp(config)
		if err != nil {
			log.Fatal(err)
			return err
		}

		server.Serve()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
