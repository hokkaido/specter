package api

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/hokkaido/specter/web/api/container"
	"net/http"
)

func RegisterRoutes(r *http.ServeMux, dockerAddr string) {
	apiRouter := mux.NewRouter()
	container.RegisterRoutes(apiRouter)

	apiWithMiddleware := negroni.New()
	apiMiddleware := NewApiMiddleware(dockerAddr)
	apiWithMiddleware.Use(apiMiddleware)
	apiWithMiddleware.UseHandler(apiRouter)
	r.Handle("/api/", apiRouter)
	/*r.Handle("/api/", negroni.New(
		NewApiMiddleware(dockerAddr),
		negroni.Wrap(apiRouter),
	))*/
}
