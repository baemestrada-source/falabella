package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/baemestrada-source/falabella/models"
)

func SearchBeers(w http.ResponseWriter, req *http.Request){
	//Simplemente devuelvo todo lo que tenga en el array
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.BeerArray)
  
}