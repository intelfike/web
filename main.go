package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/intelfike/mulpage"
	"github.com/intelfike/web/contents"
	"golang.org/x/crypto/acme/autocert"
)

var https = flag.Bool("https", false, "Using HTTPS protocol. (default HTTP)")

var port = flag.String("port", ":80", "HTTP port number")

func main() {
	flag.Parse()
	if *https {
		pack := mulpage.NewPackage(&contents.App{})

		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			mulpage.Handler(w, r, pack, "web")
		})
		log.Fatal(http.Serve(autocert.NewListener("isear.intelfike.net"), mux))
		return
	}
	mulpage.All(&contents.App{}, "web", *port)
}
