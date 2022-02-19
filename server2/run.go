package server2

import (
	"Teste/server2/http"
	"Teste/server2/models"
)

func Run() {
	server := http.NewWebServer()
	server.Products = models.Seed()
	server.Serve()
}
