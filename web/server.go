package web

import (
	"github.com/gorilla/mux"
	"github.com/hokkaido/specter/web/api"
	"net/http"
)

func ListenAndServe(listenAddr string, dockerAddr string) error {
	mainRouter := mux.NewRouter()

	api.RegisterRoutes(mainRouter, dockerAddr)

	http.Handle("/", mainRouter)

	return http.ListenAndServe(listenAddr, mainRouter)
}
