package handlers

import (
	"github.com/rwiteshbera/microservices_demo/productService/data"
	"github.com/rwiteshbera/microservices_demo/productService/helpers"
	"log"
	"net/http"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger: logger}
}

func (product *Products) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		product.getProducts(res, req)
		return
	}
	// Catch
	res.WriteHeader(http.StatusMethodNotAllowed)
}

func (product *Products) getProducts(res http.ResponseWriter, req *http.Request) {
	productList := data.GetProducts()

	err := helpers.WriteJSON(res, http.StatusOK, productList)
	if err != nil {
		log.Panic(err.Error())
	}
}
