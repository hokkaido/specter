package web

import (
	"github.com/gorilla/context"
	"github.com/hokkaido/specter/web/api"
	"net/http"
)

func ListenAndServe(listenAddr string, dockerAddr string) error {
	mainMux := http.NewServeMux()

	mainMux.Handle("/", http.FileServer(http.Dir("static")))
	api.RegisterRoutes(mainMux, dockerAddr)

	return http.ListenAndServe(listenAddr, context.ClearHandler(mainMux))
}
