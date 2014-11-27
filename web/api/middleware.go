package api

import (
	"github.com/gorilla/context"
	"github.com/samalba/dockerclient"
	"net/http"
)

type ApiMiddleware struct {
	dockerClient *dockerclient.DockerClient
}

func NewApiMiddleware(dockerAddr string) *ApiMiddleware {
	docker, _ := dockerclient.NewDockerClient(dockerAddr, nil)
	if err != nil {
		return nil
	}
	return &ApiMiddleware{docker}
}

func (l *ApiMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context.Set(r, "dockerclient", l.dockerClient)
	next(w, r)
}
