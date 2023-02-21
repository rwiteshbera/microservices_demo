package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rwiteshbera/microservices_demo/productService/data"
	"github.com/rwiteshbera/microservices_demo/productService/helpers"
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

func (product *Products) GetProduct(res http.ResponseWriter, req *http.Request) {
	query := mux.Vars(req)
	id := query["id"]
	product.logger.Println("Handle Find Product Request")

	idINT, err := strconv.Atoi(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	productData, err := findProduct(idINT)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	err = helpers.WriteJSON(res, http.StatusOK, productData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Find Product
func findProduct(id int) (*data.Product, error) {
	for _, p := range data.GetProducts() {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, errors.New("product not found")
}
