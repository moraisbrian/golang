package http

import (
	"Teste/server2/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WebServer struct {
	Products *models.Products
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, w.Products)
}

func (w WebServer) create(c echo.Context) error {
	product := models.NewProduct()

	if err := c.Bind(product); err != nil {
		return err
	}

	w.Products.Add(*product)
	return c.JSON(http.StatusCreated, product)
}

func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/products", w.getAll)
	e.POST("/products", w.create)
	e.Logger.Fatal(e.Start(":3000"))
}
