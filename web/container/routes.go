package container

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/samalba/dockerclient"
)

func RegisterRoutes(r *Router) {
	r.handleFunc("/containers", list).Methods("GET")
	r.HandleFunc("/containers/{id}", inspect).Methods("GET")
}

func kill(w http.ResponseWriter, r *http.Request) {
	docker, ok := context.GetOk(r, "docker")

}

func inspect(w http.ResponseWriter, r *http.Request) {
	docker, ok := context.GetOk(r, "docker")
	containers, err := docker.ListContainers(false)
	if err != nil {
		log.Fatal(err)
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	docker, ok := context.GetOk(r, "docker")
	containers, err := docker.ListContainers(true)
	if err != nil {
		log.Fatal(err)
	}
}

func restart(w http.ResponseWriter, r *http.Request) {
	docker, ok := context.GetOk(r, "docker")

	vars := mux.Vars(r)
	id := vars["id"]

	err = docker.RestartContainer(id)
	if err != nil {
		log.Fatal(err)
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	docker, ok := context.GetOk(r, "docker")

	vars := mux.Vars(r)
	id := vars["id"]

	err = docker.StartContainer(id)
	if err != nil {
		log.Fatal(err)
	}
}

func stop(w http.ResponseWriter, r *http.Request) {
	docker, ok := context.GetOk(r, "docker")

	vars := mux.Vars(r)
	id := vars["id"]

	err = docker.StopContainer(id)
	if err != nil {
		log.Fatal(err)
	}
}
