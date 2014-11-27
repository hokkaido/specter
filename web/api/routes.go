package api

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/hokkaido/specter/web/api/container"
)

func RegisterRoutes(r *mux.Router, dockerAddr string) {
	apiRouter := mux.NewRouter()
	container.RegisterRoutes(apiRouter)
	r.Handle("/api/", negroni.New(
		NewApiMiddleware(dockerAddr),
		negroni.Wrap(apiRouter),
	))
}
