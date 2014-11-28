package api

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/samalba/dockerclient"
	"net/http"
)

type ApiMiddleware struct {
	dockerClient dockerclient.Client
}

func NewApiMiddleware(dockerAddr string) *ApiMiddleware {
	docker, err := dockerclient.NewDockerClient(dockerAddr, nil)
	if err != nil {
		return nil
	}
	return &ApiMiddleware{dockerClient: docker}
}

func (l *ApiMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "dockerclient", l.dockerClient)
	if next != nil {
		next(w, r)
	}
}
