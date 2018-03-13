package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"gopkg.in/mgo.v2"

	"github.com/nomkhonwaan/myblog-server/pkg/graphql"
	"github.com/nomkhonwaan/myblog-server/pkg/post"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var version, revision string

func init() {
	cli.VersionPrinter = func(cc *cli.Context) {
		fmt.Println(cc.App.Name, cc.App.Version, revision)
	}
}

func main() {
	a := cli.NewApp()
	a.Name = "myblog-server"
	a.Version = version
	a.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "l,logging-level",
			Usage: "set application logging level",
			Value: logrus.InfoLevel.String(),
		},
		cli.StringFlag{
			Name:   "p,port",
			Usage:  "a listening port",
			EnvVar: "PORT",
			Value:  "8080",
		},
		cli.StringFlag{
			Name:   "mongodb-url",
			Usage:  "a MongoDB connection URL",
			EnvVar: "MONGODB_URL",
		},
	}
	a.Before = before
	a.Action = action

	if err := a.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func before(cc *cli.Context) error {
	lvl, err := logrus.ParseLevel(cc.String("logging-level"))
	if err != nil {
		return err
	}
	logrus.SetLevel(lvl)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})

	return nil
}

func action(cc *cli.Context) error {
	ctx := context.Background()

	dialInfo, err := mgo.ParseURL(cc.String("mongodb-url"))
	if err != nil {
		return err
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return err
	}

	db := session.DB(dialInfo.Database)
	repositories := make(map[string]interface{})

	repositories["post"], err = post.NewRepository(db)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, "repositories", repositories)

	mux := http.NewServeMux()
	mux.Handle("/graphql", addContext(ctx, graphql.New()))
	mux.HandleFunc("/schema.graphql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/graphql")
		w.Write([]byte(graphql.Schema()))
	})

	addr := "0.0.0.0:" + cc.String("port")
	go func() {
		http.ListenAndServe(addr, mux)
	}()
	logrus.Infof("a server is listening on address: %v", addr)

	select {}
}

func addContext(ctx context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
