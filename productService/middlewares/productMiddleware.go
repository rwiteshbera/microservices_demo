package middlewares

import (
	"github.com/go-playground/validator/v10"
	"github.com/rwiteshbera/microservices_demo/productService/data"
	"net/http"
)

func MiddlewareProductValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		p := validator.New()
		err := p.Struct(&data.Product{})
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(res, req)
	})
}
