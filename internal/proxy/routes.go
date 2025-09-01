package proxy

import (
	"os"
	"regexp"
)

func routes() map[string][]*regexp.Regexp {
	r := make(map[string][]*regexp.Regexp)

	for e, routes := range api {
		if os.Getenv(e) != "1" {
			continue
		}

		for method, paths := range routes {
			for _, path := range paths {
				r[method] = append(r[method], regexp.MustCompile(`^`+versionMatcher+path+`$`))
			}
		}
	}

	return r
}
