package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/0x616d/docker-socket-proxy/internal/proxy"
)

var healthcheck *bool = flag.Bool("H", false, "perform a healthcheck")
var logRequests *bool = flag.Bool("V", false, "log requests")
var dockerSocketPath *string = flag.String("s", "/var/run/docker.sock", "path to the docker.sock")

func main() {
	flag.Parse()

	if *healthcheck {
		doHealthcheck()
	}

	c := proxy.Config{
		LogRequests: *logRequests,
		DockerSocket: *dockerSocketPath,
	}

	p := proxy.New(c)

	if err := p.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func doHealthcheck() {
	r, err := http.Head("http://127.0.0.1:2375/_ping")
	if err != nil || r.StatusCode != http.StatusOK {
		os.Exit(1)
	}
	os.Exit(0)
}
