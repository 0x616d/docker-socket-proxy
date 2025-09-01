package proxy

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strings"
	"time"
)

type Config struct {
	LogRequests bool
	DockerSocket string
}

type Proxy struct {
	c Config
	r map[string][]*regexp.Regexp
	x *httputil.ReverseProxy
}

func New(c Config) *Proxy {
	p := new(Proxy)

	p.c = c

	p.x = &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = "http"
			r.URL.Host = "docker"
		},
		Transport: &http.Transport{
			MaxIdleConns:        500,
			MaxIdleConnsPerHost: 500,
			MaxConnsPerHost:     500,

			IdleConnTimeout:       5 * time.Second,
			ResponseHeaderTimeout: 3 * time.Second,
			ExpectContinueTimeout: 3 * time.Second,

			Dial: func(_, _ string) (net.Conn, error) {
				return net.Dial("unix", p.c.DockerSocket)
			},
		},
	}

	p.r = routes()

	return p
}

func (p *Proxy) logReq(action string, r *http.Request) {
	if !p.c.LogRequests {
		return
	}
	if r.URL.Path == "/_ping" && strings.HasPrefix(r.RemoteAddr, "127.0.0.1") {
		return // don't log healthchecks requests
	}
	log.Printf(`%s - %s "%s %s"`, action, r.RemoteAddr, r.Method, r.URL.Path)
}

func (p *Proxy) route(method, path string) bool {
	for _, r := range p.r[method] {
		if r.MatchString(path) {
			return true
		}
	}
	return false
}

func (p *Proxy) handler(w http.ResponseWriter, r *http.Request) {
	if !p.route(r.Method, r.URL.Path) {
		p.logReq("DENY", r)
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	p.logReq("PASS", r)
	p.x.ServeHTTP(w, r)
}

func (p *Proxy) ListenAndServe() error {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", p.handler)
	return http.ListenAndServe(":2375", mux)
}
