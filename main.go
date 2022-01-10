package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/baemestrada-source/falabella/models"
  "github.com/baemestrada-source/falabella/routers"
  "github.com/rs/cors"
  "os"
)

type LocationResponse struct {
	Country string `json:"country"`
}

func main() {
  router := mux.NewRouter()
  
  // Informacion de ejemplo de entrada
  models.BeerArray = append(models.BeerArray, models.BeerItem{Id: 1, Name:"Golden", Brewery:"kross",Country:"Chile", Price:10.5, Currency:"USD" })
  models.BeerArray = append(models.BeerArray, models.BeerItem{Id: 2, Name:"Gallo",  Brewery:"Cerveceria CA",Country:"Guate", Price:5,    Currency:"GTQ" })

  // endpoints rutas de mi API
  routers.SetBeersRoutes(router)

  //log 
  log.Println("Servidor corriendo en el puerto 4000")

  PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "4000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}