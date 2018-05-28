package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type InsecureServer struct {
	*http.ServeMux
	ShutdownTimeout time.Duration
}

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
			logrus.Infof("server has been stopped")
		default:
			logrus.Fatalf("insecure_server: %v", err)
		}
	}()

	logrus.Infof("server is listening on address: %s", l.Addr().String())
	return nil
}
