package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/baemestrada-source/falabella/models"
	"github.com/gorilla/mux"
	"strconv"
)

//searchBeerById Busca una cerveza por su Id
func SearchBeerById(w http.ResponseWriter, req *http.Request){
    params := mux.Vars(req)

    // la variable que obtengo de el path
    beerID, err := strconv.Atoi(params["beerID"])

    // si no hay errores inicio
    if err == nil {
        for _, item := range models.BeerArray {
            if item.Id == beerID {
              //el ID fue encontrado lo retorno en JSON
              w.Header().Set("Content-Type", "application/json")
              w.WriteHeader(http.StatusOK)
              json.NewEncoder(w).Encode(item)
              return
            }
          }
    } 
    http.Error(w,"El Id de la cerveza no existe", http.StatusNotFound)
  }