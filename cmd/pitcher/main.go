package main

import (
	"log"
	"net"
	"os"

	"github.com/hugoatease/pitcher"
	pb "github.com/hugoatease/pitcher/protobuf"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
		cli.StringFlag{
			Name:   "solr",
			Value:  "http://localhost:8983/",
			Usage:  "solr server url",
			EnvVar: "SOLR_URL",
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
			SolrURL:    c.String("solr"),
		}

		lis, err := net.Listen("tcp", config.Bind)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		server, err := pitcher.NewServer(config)
		if err != nil {
			log.Fatal(err)
			return err
		}

		pb.RegisterPitcherServer(grpcServer, server)
		reflection.Register(grpcServer)
		grpcServer.Serve(lis)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
