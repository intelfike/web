package main

import (
	"flag"

	"github.com/intelfike/mulpage"
	"github.com/intelfike/web/contents"
)

var port = flag.String("http", ":8888", "HTTP port number.")

func main() {
	flag.Parse()

	mulpage.All(&contents.App{}, "web", *port)
}
