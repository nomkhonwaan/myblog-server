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

// // APIServer is a wrapper of `http.Server`
// type APIServer struct {
// 	*http.ServeMux
// 	ShutdownTimeout time.Duration
// }

// // NewInsecureAPIServer returns a new APIServer without SSL & TLS configure
// func NewInsecureAPIServer() (*APIServer, error) {
// 	return &APIServer{
// 		ServeMux:        http.NewServeMux(),
// 		ShutdownTimeout: time.Second * 10,
// 	}, nil
// }

// // Handle registers the handler for the given pattern.
// // If a handler already exists for pattern, Handle panics.
// func (s *APIServer) Handle(pattern string, handler http.Handler) {
// 	s.ServeMux.Handle(pattern, handler)
// }

// // ListenAndServe listens and serves on given address
// func (s *APIServer) ListenAndServe(addr string, stopCh <-chan struct{}) error {
// 	server := http.Server{
// 		Handler: s.ServeMux,
// 	}

// 	l, err := net.Listen("tcp", addr)
// 	if err != nil {
// 		return err
// 	}

// 	go func() {
// 		<-stopCh

// 		logrus.Infof("server is shutting down...")
// 		ctx, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeout)
// 		defer cancel()

// 		server.Shutdown(ctx)
// 	}()

// 	go func() {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				logrus.Errorf("recover: %v", r)
// 			}
// 		}()

// 		err := server.Serve(l)

// 		select {
// 		case <-stopCh:
// 			logrus.Infof("a server has been stopped")
// 		default:
// 			logrus.Fatal(err)
// 		}
// 	}()

// 	logrus.Infof("a server is listening on address: %s", l.Addr().String())
// 	return nil
// }
