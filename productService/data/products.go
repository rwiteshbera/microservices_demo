package data

import "time"

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	CreatedOn   string  `json:"createdOn"`
	UpdatedOn   string  `json:"updatedOn"`
	DeletedOn   string  `json:"deletedOn"`
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Milky Coffee",
		Price:       299,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},

	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Strong coffee without milk",
		Price:       199,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
