package server2

import (
	"estudos/server2/http"
	"estudos/server2/models"
)

func Run() {
	server := http.NewWebServer()
	server.Products = models.Seed()
	server.Serve()
}
