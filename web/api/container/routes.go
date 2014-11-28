package container

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/samalba/dockerclient"
	"html"
	"log"
	"net/http"
)

func getDockerClient(r *http.Request) dockerclient.Client {
	if rv := context.Get(r, "dockerclient"); rv != nil {
		fmt.Printf("omg")
		return rv.(dockerclient.Client)
	}
	fmt.Printf("so nil")
	return nil
}

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/containers", list).Methods("GET")
	r.HandleFunc("/api/containers/{id}", inspect).Methods("GET")
	r.HandleFunc("/api/hello", hello).Methods("GET")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func kill(w http.ResponseWriter, r *http.Request) {
	return
}

func inspect(w http.ResponseWriter, r *http.Request) {
	docker := getDockerClient(r)
	vars := mux.Vars(r)
	id := vars["id"]
	container, err := docker.InspectContainer(id)
	if err != nil {
		log.Fatal(err)
	}

	js, err := json.Marshal(container)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func list(w http.ResponseWriter, r *http.Request) {
	docker := getDockerClient(r)
	containers, err := docker.ListContainers(true, true, "")
	if err != nil {
		log.Fatal(err)
	}

	js, err := json.Marshal(containers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func restart(w http.ResponseWriter, r *http.Request) {
	docker := getDockerClient(r)

	vars := mux.Vars(r)
	id := vars["id"]

	err := docker.RestartContainer(id, 100)
	if err != nil {
		log.Fatal(err)
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	docker := getDockerClient(r)

	vars := mux.Vars(r)
	id := vars["id"]

	err := docker.StartContainer(id, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func stop(w http.ResponseWriter, r *http.Request) {
	docker := getDockerClient(r)

	vars := mux.Vars(r)
	id := vars["id"]

	err := docker.StopContainer(id, 100)
	if err != nil {
		log.Fatal(err)
	}
}
