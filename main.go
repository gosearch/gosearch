package main

import (
	"github.com/gosearch/gosearch/http"
	"github.com/gosearch/gosearch/service"
)

func main() {
	server := &http.Server{Index: &service.DefaultIndexService{}}
	server.Listen(9093)
}
