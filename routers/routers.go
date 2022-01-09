package routers

import (
	"github.com/gorilla/mux"
	"github.com/baemestrada-source/falabella/handlers"

)

func SetBeersRoutes(r *mux.Router) {
	r.HandleFunc("/beers", handlers.SearchBeers).Methods("GET")
	r.HandleFunc("/beers", handlers.AddBeers).Methods("POST")
	r.HandleFunc("/beers/{beerID}", handlers.SearchBeerById).Methods("GET")
	r.HandleFunc("/beers/{beerID}/boxprice", handlers.BoxBeerPriceById).Methods("GET")
}