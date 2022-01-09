package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/baemestrada-source/falabella/models"
  "github.com/baemestrada-source/falabella/routers"
)

func main() {
  router := mux.NewRouter()
  
  // Informacion de ejemplo de entrada
  models.BeerArray = append(models.BeerArray, models.BeerItem{Id: 1, Name:"Golden", Brewery:"kross",Country:"Chile", Price:10.5, Currency:"USD" })
  models.BeerArray = append(models.BeerArray, models.BeerItem{Id: 2, Name:"Gallo",  Brewery:"Cerveceria CA",Country:"Guate", Price:5,    Currency:"GTQ" })

  // endpoints rutas de mi API
  routers.SetBeersRoutes(router)

  //log 
  log.Println("Servidor corriendo en el puerto 4000")

  //escucha del servidor
  log.Fatal(http.ListenAndServe(":4000", router))
}