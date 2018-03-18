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

	setLogrusLevelAndOutputTextFormat(*loggingLevel)

	session, err := makeANewConnectionToMongoDB(*mongodbConnectionURI)
	handleErrors(err)

	graphqlHandler := graphql.Handler{}

	var g inject.Graph
	handleErrors(
		g.Provide(
			&inject.Object{Name: "pkg/graphql.Handler", Value: &graphqlHandler},
			&inject.Object{Name: "pkg/graphql/resolver.Resolver", Value: &resolver.Resolver{}},
			&inject.Object{Name: "pkg/post.Repositorier", Value: post.NewRepository(session.Clone().DB(defaultDatabaseName))},
		),
		g.Populate(),
	)

	server, err := app.NewInsecureAPIServer(
		graphqlHandler,
	)
	handleErrors(err)

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
