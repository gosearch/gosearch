package main

import (
	"flag"

	"github.com/gosearch/gosearch/http"
	"github.com/gosearch/gosearch/service"
)

func main() {
	port := flag.Int("p", 9093, "The `port` gosearch will listen on.")
	help := flag.Bool("h", false, "Print gosearch's help.")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	server := &http.Server{Index: &service.DefaultIndexService{}}
	server.Listen(*port)
}
