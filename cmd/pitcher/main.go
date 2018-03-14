package main

import (
	"log"
	"os"

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
	}

	app.Action = func(c *cli.Context) error {
		config := pitcher.Config{
			Bind:   c.String("bind"),
			DbHost: c.String("dbhost"),
			DbPort: c.String("dbport"),
			DbName: c.String("dbname"),
			DbUser: c.String("dbuser"),
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
