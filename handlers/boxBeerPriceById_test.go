package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/baemestrada-source/falabella/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

const curr = "GTQ"
const quan = "4"

// segun los datos de prueba el valor de el dato de prueba son 10.5 EUR si convierto segun la cantidad y tipo de cambio
// me da 366.5824 con este valor validare si estoy bien con el test
const Esperado = 366.5824   

func TestGetRequest(t *testing.T) {
	//Limpio cualquier valor que traiga de array de otro test
	models.BeerArray = nil

	//Ingreso un dato de para el test moneda EUR valor 10.5 
	models.BeerArray = append(models.BeerArray, models.BeerItem{Id: 1, Name:"Golden", Brewery:"kross",Country:"Chile", Price:10.5, Currency:"EUR" })
    r, _ := http.NewRequest("GET", "/beers/beerID/boxprice?currency="+curr+"&quantity="+quan, nil)
    w := httptest.NewRecorder()


    //Hack to try to fake gorilla/mux vars
    vars := map[string]string{
        "beerID": "1",
	}

    r = mux.SetURLVars(r, vars)

    BoxBeerPriceById(w, r)

	//primero valido que realmente tenga el retorno de que si existe el id 
    assert.Equal(t, http.StatusOK, w.Code)

	var beerBox models.BeerBox
	json.Unmarshal([]byte(w.Body.String()), &beerBox)

	//valido que los valores sean iguales segun el resultado de mi test
    assert.Equal(t, float32(Esperado), float32(beerBox.PriceTotal))
}