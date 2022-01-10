package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/baemestrada-source/falabella/models"
)

//addBeers Ingresa una nueva cerveza
 func AddBeers(w http.ResponseWriter, req *http.Request){
    var beer models.BeerItem
    err := json.NewDecoder(req.Body).Decode(&beer)

    //Unicamente entra al if si no hay ningun error con el body
    if err == nil {
      for _, item := range models.BeerArray {
        if item.Id == beer.Id {
          http.Error(w,"El ID de la cerveza ya existe", 409)
          return
        }
      }
      //aqui si no encontro ningun error regreso el item y respondo que todo esta bien status 201
      models.BeerArray = append(models.BeerArray, beer)

      //retorno el mensaje al servidor para indicar que la cerveza fue creada exitosamente
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusCreated)
      json.NewEncoder(w).Encode("Cerveza creada")

      } else { 
      //error 400 ya que el body tiene error 
      http.Error(w,"Request invalida", http.StatusBadRequest)
    }
  }