package middlewares

import (
	"log"
	"net/http"
)

func MiddlewareProductValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Println("Product validator executed")
		next.ServeHTTP(res, req)
	})
}
