package main

import (
	"log"
	"net/http"

	"github.com/intelfike/mulpage"
	"github.com/intelfike/web/contents"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	pack := mulpage.NewPackage(&contents.App{})

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mulpage.Handler(w, r, pack, "web")
	})
	log.Fatal(http.Serve(autocert.NewListener("isear.intelfike.net"), mux))
}
