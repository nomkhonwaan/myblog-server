package app

import "net/http"

// Handler is an HTTP handler interface that allowed any services register its handlers functions
type Handler interface {
	Init(*http.ServeMux) error
}
