package app

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// InsecureServer provides an HTTP server without SSL, TLS configure
type InsecureServer struct {
	*http.ServeMux
	ShutdownTimeout time.Duration
}

// ListenAndServe listens and serves an HTTP server in the background on given address
func (s *InsecureServer) ListenAndServe(addr string, stopCh <-chan struct{}) error {
	server := http.Server{
		Handler: s.ServeMux,
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		<-stopCh

		logrus.Infof("server is shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeout)
		defer cancel()

		server.Shutdown(ctx)
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				logrus.Errorf("recover: %v", r)
			}
		}()

		err := server.Serve(l)

		select {
		case <-stopCh:
			logrus.Infof("a server has been stopped")
		default:
			logrus.Fatal(err)
		}
	}()

	logrus.Infof("a server is listening on address: %s", l.Addr().String())
	return nil
}
