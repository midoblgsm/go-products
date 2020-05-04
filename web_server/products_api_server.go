package web_server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/midoblgsm/go-products/resources"
	"github.com/midoblgsm/go-products/utils"

	"github.com/gorilla/mux"
)

type ProductsApiServer struct {
	port           int
	ratingEndpoint string
}

func NewProductsApiServer(port int, ratingEndpoint string) ProductsApiServer {
	return ProductsApiServer{
		port:           port,
		ratingEndpoint: ratingEndpoint,
	}
}

func (p *ProductsApiServer) InitializeHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/v1/products", p.ListProducts()).Methods("GET")
	return router
}

func (p *ProductsApiServer) Start() error {
	router := p.InitializeHandler()
	http.Handle("/", router)

	log.Printf("Starting Products API server on port %d ....", p.port)
	log.Println("CTL-C to exit/stop Products API server")

	return http.ListenAndServe(fmt.Sprintf(":%d", p.port), nil)

}

func (p *ProductsApiServer) ListProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		p.EnableCORS(w, req)
		var products []resources.Product

		p1 := resources.CreateProduct("Potato", "Potatoes are frequently served whole or mashed as a cooked vegetable and are also ground into potato flour, used in baking and as a thickener for sauces.")
		p1.GetRating(p.ratingEndpoint)
		p2 := resources.CreateProduct("Banana", " If you eat a banana every day for breakfast, your roommate might nickname you 'the monkey.' A banana is a tropical fruit that's quite popular all over the world. It grows in bunches on a banana tree.")
		p2.GetRating(p.ratingEndpoint)
		p3 := resources.CreateProduct("Pepper", "Pepper or black pepper is the dried unripe fruit grown in the plant called piper nigrum. It's pungent smell, peppery/hot taste and health friendly properties make pepper a favorite spice all over the world and it is commonly used in all cuisines.")
		p3.GetRating(p.ratingEndpoint)


		products = append(products, p1)
		products = append(products, p2)
		products = append(products, p3)

		log.Printf("%#v", products)
		utils.WriteResponse(w, http.StatusOK, products)
	}
}

func (p *ProductsApiServer) EnableCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")

	return
}
