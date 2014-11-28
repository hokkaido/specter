package main

import (
	"flag"
	"github.com/Sirupsen/logrus"
	"github.com/hokkaido/specter/web"
)

var (
	listenAddr string
	dockerAddr string
	logger     = logrus.New()
)

func init() {
	flag.StringVar(&listenAddr, "listen", ":8080", "listen address")
	flag.StringVar(&dockerAddr, "docker", "unix:///var/run/docker.sock", "docker address")
}

func main() {

	if err := web.ListenAndServe(listenAddr, dockerAddr); err != nil {
		logger.Fatal(err)
	}

}
