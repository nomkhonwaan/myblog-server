package app

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// APIServer is a wrapper of `http.Server`
type APIServer struct {
	*http.Server
	ShutdownTimeout time.Duration
}

// NewInsecureAPIServer returns a new APIServer without SSL & TLS configure
func NewInsecureAPIServer(handlers ...Handler) (*APIServer, error) {
	mux := http.NewServeMux()

	for _, h := range handlers {
		if err := h.Init(mux); err != nil {
			return nil, err
		}
	}

	return &APIServer{
		Server: &http.Server{
			Handler: mux,
		},
		ShutdownTimeout: time.Second * 10,
	}, nil
}

// ListenAndServe listens and serves on given address
func (s *APIServer) ListenAndServe(addr string, stopCh <-chan struct{}) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		<-stopCh

		logrus.Infof("server is shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeout)
		defer cancel()

		s.Server.Shutdown(ctx)
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				logrus.Errorf("recover: %v", r)
			}
		}()

		err := s.Server.Serve(l)

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
