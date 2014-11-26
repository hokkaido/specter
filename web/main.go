package main

import (
	"flag"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/specter/web/container"
	"net/http"
)

var (
	listenAddr string
	dockerAddr string
	logger     = logrus.New()
)

func init() {
	flag.StringVar(&listenAddr, "listen", ":8080", "listen address")
	flag.StringVar(&dockerAddr, "docker", "unix:///tmp/docker.sock", "docker address")
}

func main() {

	mainRouter := mux.newRouter()

	apiRouter := mainRouter.PathPrefix("/api").Subrouter()
	container.RegisterRoutes(apiRouter)

	http.Handle("/", mainRouter)

	if err := http.ListenAndServe(listenAddr, mainRouter); err != nil {
		logger.Fatal(err)
	}

}
