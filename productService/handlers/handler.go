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

func NewProductHandler(logger *log.Logger) *Products {
	return &Products{logger: logger}
}

func (product *Products) GetProducts(res http.ResponseWriter, req *http.Request) {
	productList := data.GetProducts()

	err := helpers.WriteJSON(res, http.StatusOK, productList)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (product *Products) AddProduct(res http.ResponseWriter, req *http.Request) {
	product.logger.Println("Handle POST Product")

	reqData := &data.Product{}
	err := helpers.ReadJSON(res, req, reqData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	data.AddProduct(reqData)
}
