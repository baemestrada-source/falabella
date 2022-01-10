package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/baemestrada-source/falabella/models"
	"github.com/gorilla/mux"
	"strconv"
	"io/ioutil"  
)

  //boxBeerPriceById Obtiene el precio de una caja de cerveza por su Id
  func BoxBeerPriceById(w http.ResponseWriter, req *http.Request){
    params := mux.Vars(req)

    query := req.URL.Query()
    currency := query.Get("currency")

    //Cantidad de cervezas a comprar (valor por defecto 6), la cantidad viene en string la paso a float32
    quantity, err := strconv.ParseFloat(query.Get("quantity"),32)
    if err != nil {
      quantity = 6
    }

    //paso a entero para comparar 
    beerID, err := strconv.Atoi(params["beerID"])
    if err == nil {
        for _, item := range models.BeerArray {
            if item.Id == beerID {
              //consulto la API con la variable que ya tenia definida
              response, err := http.Get("http://api.currencylayer.com/live?access_key="+models.Access_key+"&format=1")      
              if err != nil {
                  http.Error(w, "Ocurrio un error al leer la API de tipo de cambio", http.StatusBadRequest)
                  return
              }
              
              responseData, err := ioutil.ReadAll(response.Body)
              if err != nil {
                http.Error(w, "Ocurrio un error al leer la API de tipo de cambio", http.StatusBadRequest)
                return
              }

              //log.Println(string(responseData))
              
              var responseObject models.ResponseTC
              json.Unmarshal(responseData, &responseObject)

              if responseObject.Success {
                for k, v := range responseObject.Quotes {
                  // Busca el tipo de cambio con respecto al USD de la moneda con la que se pagara
                    if k[3:6] == currency { 
                      models.TpcIn = v
                    }
                    // Busca el tipo de cambio con respecto al USD a la moneda que esta en la data
                    if k[3:6] == item.Currency {  
                      models.TpcDb = v
                    }
                }
                //valor homologado utilizando los valores de la API
                Valor_H :=  ( item.Price / models.TpcDb ) * models.TpcIn
                //log.Println(TpcIn,TpcDb, Valor_H)


                //Calculo el precio total ya con la moneda que corresponde
                ValPTot:=Valor_H  * float32(quantity)  

                // Lleno mi estructura de respuesta
                BeerBox:=models.BeerBox{PriceTotal:ValPTot} 

                //respondo la API con la estructura
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusOK)
                json.NewEncoder(w).Encode(&BeerBox)
              }
          
              return
            }
          }
    } 
    http.Error(w,"El Id de la cerveza no existe", http.StatusNotFound)
  }

