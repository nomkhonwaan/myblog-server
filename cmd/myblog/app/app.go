package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nomkhonwaan/myblog-server/pkg/tag"

	"github.com/facebookgo/inject"
	"github.com/nomkhonwaan/myblog-server/pkg/auth"
	"github.com/nomkhonwaan/myblog-server/pkg/graphql"
	"github.com/nomkhonwaan/myblog-server/pkg/graphql/resolver"
	"github.com/nomkhonwaan/myblog-server/pkg/post"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gopkg.in/mgo.v2"
)

const (
	defaultDatabaseName = "nomkhonwaan_com"
	usage               = `
        __  ___      ____  __           
       /  |/  /_  __/ __ )/ /___  ____ _
      / /|_/ / / / / __  / / __ \/ __ ` + "`" + `/
     / /  / / /_/ / /_/ / / /_/ / /_/ / 
    /_/  /_/\__, /_____/_/\____/\__, /  
           /____/              /____/
`
)

var (
	version, revision string
)

func init() {
	cli.VersionPrinter = func(ctx *cli.Context) {
		fmt.Println(ctx.App.Name, ctx.App.Version, revision)
	}
}

// New returns a new application
func New() *cli.App {
	app := cli.NewApp()
	app.Usage = usage
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "logging-level",
			Usage:  "a minimum level of the log that will print out",
			EnvVar: "LOGGING_LEVEL",
			Value:  logrus.InfoLevel.String(),
		},
		cli.StringFlag{
			Name:   "auth0-jwks-uri",
			Usage:  "an Auth0 JWKS URI",
			EnvVar: "AUTH0_JWKS_URI",
		},
		cli.StringSliceFlag{
			Name:   "auth0-audience",
			Usage:  "an Auth0 audience",
			EnvVar: "AUTH0_AUDIENCE",
		},
		cli.StringFlag{
			Name:   "auth0-issuer",
			EnvVar: "AUTH0_ISSUER",
			Value:  "an Auth0 issuer",
		},
		cli.StringFlag{
			Name:   "listen-address",
			Usage:  "a listening address of the API server",
			EnvVar: "LISTEN_ADDRESS",
			Value:  "0.0.0.0:8080",
		},
		cli.StringFlag{
			Name:   "mongodb-connection-uri",
			Usage:  "a MongoDB connection URI",
			EnvVar: "MONGODB_CONNECTION_URI",
			Value:  "mongodb://localhost/nomkhonwaan_com",
		},
	}
	app.Before = before
	app.Action = action

	return app
}

func before(ctx *cli.Context) error {
	setLogLevelAndOutputTextFormat(ctx.String("logging-level"))

	return nil
}

func action(ctx *cli.Context) error {
	session, err := makeANewConnectionToMongodb(ctx.String("mongodb-connection-uri"))
	if err != nil {
		return err
	}

	authMiddleware := auth.NewMiddleware(
		ctx.String("auth0-jwks-uri"),
		ctx.StringSlice("auth0-audience"),
		ctx.String("auth0-issuer"),
	)

	graphqlHandler := graphql.Handler{}

	var g inject.Graph

	err = g.Provide(
		&inject.Object{Name: "pkg/graphql.Handler", Value: &graphqlHandler},
		&inject.Object{Name: "pkg/graphql/resolver.Resolver", Value: &resolver.Resolver{}},
		&inject.Object{Name: "pkg/post.Repositorier", Value: post.NewRepository(session.Clone().DB(defaultDatabaseName))},
		&inject.Object{Name: "pkg/tag.Repositorier", Value: tag.NewRepository(session.Clone().DB(defaultDatabaseName))},
	)
	if err != nil {
		return err
	}
	if err := g.Populate(); err != nil {
		return err
	}

	server := InsecureServer{
		ServeMux:        http.NewServeMux(),
		ShutdownTimeout: time.Second * 15,
	}

	server.Handle("/graphql", authMiddleware(graphqlHandler))

	stopCh := handleSignals()
	server.ListenAndServe(ctx.String("listen-address"), stopCh)

	<-stopCh

	return nil
}

func handleSignals() <-chan struct{} {
	stopCh := make(chan struct{})
	sigsCh := make(chan os.Signal, 2)
	signal.Notify(sigsCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigsCh
		close(stopCh)
		<-sigsCh
		os.Exit(1)
	}()

	return stopCh
}

func setLogLevelAndOutputTextFormat(loggingLevel string) {
	level, err := logrus.ParseLevel(loggingLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
}

func makeANewConnectionToMongodb(connectionURI string) (*mgo.Session, error) {
	dialInfo, err := mgo.ParseURL(connectionURI)
	if err != nil {
		return nil, err
	}

	return mgo.DialWithInfo(dialInfo)
}
