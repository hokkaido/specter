package web

import (
	"github.com/gorilla/mux"
	"github.com/hokkaido/specter/web/container"
	"net/http"
)

func ListenAndServe(addr string) error {
	mainRouter := mux.NewRouter()

	apiRouter := mainRouter.PathPrefix("/api").Subrouter()
	container.RegisterRoutes(apiRouter)

	http.Handle("/", mainRouter)

	return http.ListenAndServe(addr, mainRouter)
}