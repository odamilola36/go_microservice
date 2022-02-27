package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"microservices/api/middlewares"
	"net/http"
)

type Route struct {
	Method       string
	Path         string
	HandlerFunc  http.HandlerFunc
	AuthRequired bool
}

func Install(router *mux.Router, routes []*Route) {
	for _, route := range routes {
		if route.AuthRequired {
			router.
				Handle(route.Path, middlewares.LogRequests(middlewares.CheckAuth(route.HandlerFunc))).
				Methods(route.Method).Name(route.Path)
		} else {
			router.
				HandleFunc(route.Path, middlewares.LogRequests(route.HandlerFunc)).
				Methods(route.Method)
		}
	}
}

func WithCors(router *mux.Router) http.Handler {
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	return handlers.CORS(headers, origins, methods)(router)
}
