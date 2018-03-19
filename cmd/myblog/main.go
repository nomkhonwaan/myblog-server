package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/facebookgo/inject"
	"github.com/nomkhonwaan/myblog-server/cmd/myblog/app"
	"github.com/nomkhonwaan/myblog-server/pkg/auth"
	"github.com/nomkhonwaan/myblog-server/pkg/graphql"
	"github.com/nomkhonwaan/myblog-server/pkg/graphql/resolver"
	"github.com/nomkhonwaan/myblog-server/pkg/post"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

const (
	defaultDatabaseName = "nomkhonwaan_com"
)

var version, revision string

var (
	domain               = flag.String("auth0-domain", "", "an Auth0 domain")
	audience             = flag.String("auth0-api-audience", "", "an Auth0 API audience")
	clientSecret         = flag.String("auth0-api-client-secret", "", "an Auth0 API client secret")
	loggingLevel         = flag.String("logging-level", logrus.InfoLevel.String(), "a minimum level of the log that will print out")
	listenAddress        = flag.String("listen-address", "0.0.0.0:8080", "a listening address of the API server")
	mongodbConnectionURI = flag.String("mongodb-connection-uri", "mongodb://localhost/nomkhonwaan_com", "a MongoDB connection URI")
	printVersion         = flag.Bool("version", false, "display the version information")
)

func main() {
	flag.Parse()

	if *printVersion {
		fmt.Printf("myblog: version %s, revision %s\n", version, revision)
		return
	}

	// Setup Logrus with given logging-level (or "info" if parse error) and force display log full timestamp
	setLogrusLevelAndOutputTextFormat(*loggingLevel)

	// Make a new connection to MongoDB server
	session, err := makeANewConnectionToMongoDB(*mongodbConnectionURI)
	handleErrors(err)

	// Initialize all handlers
	authHandler := auth.NewHandler(*domain, []string{*audience}, *clientSecret)
	graphqlHandler := graphql.Handler{}

	// Setup the dependency injection using facebookgo/inject
	var g inject.Graph
	handleErrors(
		g.Provide(
			&inject.Object{Name: "pkg/graphql.Handler", Value: &graphqlHandler},
			&inject.Object{Name: "pkg/graphql/resolver.Resolver", Value: &resolver.Resolver{}},
			&inject.Object{Name: "pkg/post.Repositorier", Value: post.NewRepository(session.Clone().DB(defaultDatabaseName))},
		),
		g.Populate(),
	)

	// Create a new API server
	server, err := app.NewInsecureAPIServer(
		authHandler,
		graphqlHandler,
	)
	handleErrors(err)

	// Start the API server in background and wait for the stop signal
	stopCh := handleSignals()
	handleErrors(server.ListenAndServe(*listenAddress, stopCh))

	<-stopCh
}

func handleErrors(errs ...error) {
	for _, err := range errs {
		if err != nil {
			logrus.Fatal(err)
		}
	}
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

func setLogrusLevelAndOutputTextFormat(loggingLevel string) {
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

func makeANewConnectionToMongoDB(connectionURI string) (*mgo.Session, error) {
	dialInfo, err := mgo.ParseURL(connectionURI)
	if err != nil {
		return nil, err
	}

	return mgo.DialWithInfo(dialInfo)
}
