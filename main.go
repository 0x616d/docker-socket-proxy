package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var logRequests bool
var doHealthcheck bool
var backendAddr string
var frontedAddr string

func init() {
	rules["allow-system-events"].enabled = true
	rules["allow-system-info"].enabled = true
	rules["allow-system-version"].enabled = true
	rules["allow-system-ping"].enabled = true
	rules["allow-system-ping-head"].enabled = true
}

func main() {
	flag.BoolVar(&logRequests, "log-requests", false, "log requests")
	flag.BoolVar(&doHealthcheck, "healthcheck", false, "perform a healthcheck")

	flag.StringVar(&backendAddr, "backend-addr", "/var/run/docker.sock", "")
	flag.StringVar(&frontedAddr, "fronted-addr", ":2375", "")

	for name, rule := range rules {
		flag.BoolVar(&rule.enabled, name, rule.enabled, rule.description)
	}

	flag.Parse()

	if doHealthcheck {
		healthcheck()
	}

	var dialfn func(n, a string) (net.Conn, error)
	var directorfn func(r *http.Request)

	if strings.HasPrefix(backendAddr, "/") {
		dialfn = func(_, _ string) (net.Conn, error) {
			return net.Dial("unix", backendAddr)
		}
		directorfn = func(r *http.Request) {
			r.URL.Host = "localhost"
			r.URL.Scheme = "http"
		}
	} else {
		dialfn = func(_, _ string) (net.Conn, error) {
			return net.Dial("tcp", backendAddr)
		}
		backendURL, err := url.Parse("http://" + backendAddr)
		if err != nil {
			log.Fatalf("bad backend address: %s", err)
		}
		directorfn = func(r *http.Request) {
			r.URL.Host = backendURL.Host
			r.URL.Scheme = backendURL.Scheme
		}
	}

	proxy := &httputil.ReverseProxy{
		Director: directorfn,
		Transport: &http.Transport{
			IdleConnTimeout:       30 * time.Second,
			ResponseHeaderTimeout: 1 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			Dial:                  dialfn,
		},
	}

	var listener net.Listener
	var err error

	if strings.HasPrefix(frontedAddr, "/") {
		listener, err = net.Listen("unix", frontedAddr)
		defer os.Remove(frontedAddr)
	} else {
		listener, err = net.Listen("tcp", frontedAddr)
	}

	if err != nil {
		log.Fatalln(err)
	}

	rt := newRouter()

	for _, rule := range rules {
		if !rule.enabled {
			continue
		}
		for method, paths := range rule.routes {
			for _, path := range paths {
				rt.insert(method, path)
			}
		}
	}

	s := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !rt.search(r.Method, removeVersionPrefix(r.URL.Path)) {
				logRequest("DENY", r)
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			logRequest("PASS", r)
			proxy.ServeHTTP(w, r)
		}),
	}

	go s.Serve(listener)

	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGTERM, syscall.SIGINT)

	<-term

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(15*time.Second))
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		s.Close()
	}
}

func healthcheck() {
	var addr string
	var client *http.Client

	if strings.HasPrefix(frontedAddr, "/") {
		client = &http.Client{
			Transport: &http.Transport{
				Dial: func(_, _ string) (net.Conn, error) {
					return net.Dial("unix", frontedAddr)
				},
			},
		}

		addr = "localhost"
	} else {
		client = &http.Client{}

		if strings.HasPrefix(frontedAddr, ":") {
			addr = "localhost" + frontedAddr
		} else {
			addr = frontedAddr
		}
	}

	r, err := client.Head("http://" + addr + "/_ping")
	if err != nil || r.StatusCode != http.StatusOK {
		os.Exit(1)
	}
	os.Exit(0)
}

func logRequest(action string, r *http.Request) {
	if !logRequests {
		return
	}
	if r.URL.Path == "/_ping" && strings.HasPrefix(r.RemoteAddr, "127") {
		return
	}
	log.Printf(`%s - %s %s %s\n`, action, r.RemoteAddr, r.Method, r.URL.Path)
}
