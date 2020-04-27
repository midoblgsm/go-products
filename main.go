package main

import (
	"log"

	"github.com/midoblgsm/go-products/web_server"
)

func main() {
	//ratingEndpoint := os.Getenv("RATING_ENDPOINT")
	ratingEndpoint := "http://0.0.0.0:9998/v1/ratings"

	server := web_server.NewProductsApiServer(9999, ratingEndpoint)

	log.Fatal(server.Start())
}
