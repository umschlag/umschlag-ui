//
// Use this server to run this project in standalone mode to avoid having to
// re-compile your Go code and re-launch Umschlag every time you make a code
// change. This server proxies all traffic to a running Umschlag instance. This
// can be a local instance, or a remote instance, so you can develop against
// real world data.
//
// go run server.go --scheme=http --host=localhost:8080
//

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

var (
	host   = flag.String("host", "localhost:8000", "instance url")
	scheme = flag.String("scheme", "http", "instance scheme")
)

func main() {
	flag.Parse()

	// Serve the static page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index, err := ioutil.ReadFile("index.html")

		if err != nil {
			log.Println(err)
		}

		w.Write([]byte(index))
	})

	// Serve static content
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(
				http.Dir("assets/"),
			),
		),
	)

	// Proxy requests to API
	http.Handle(
		"/api/",
		&httputil.ReverseProxy{
			Director: func(req *http.Request) {
				req.URL.Scheme = *scheme
				req.URL.Host = *host
				req.Host = *host

				req.Header.Set("X-Forwarded-For", *host)
				req.Header.Set("X-Forwarded-Proto", *scheme)
			},
		},
	)

	log.Println("Listening on port 9000...")
	http.ListenAndServe(":9000", nil)
}
