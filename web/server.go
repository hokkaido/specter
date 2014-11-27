package web

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/hokkaido/specter/web/api"
	"net/http"
)

func ListenAndServe(listenAddr string, dockerAddr string) error {
	mainRouter := mux.NewRouter()

	mainRouter.Handle("/", http.FileServer(http.Dir("static")))
	api.RegisterRoutes(mainRouter, dockerAddr)

	return http.ListenAndServe(listenAddr, context.ClearHandler(mainRouter))
}
