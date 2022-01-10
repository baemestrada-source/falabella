package handlers

import (
//	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/baemestrada-source/falabella/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)
//Valida el retorno de cuando si encuentra un registro codigo 200
func TestGetBeerByIdExits(t *testing.T) {
    //Limpio cualquier valor que traiga de array de otro test
	models.BeerArray = nil
    
	models.BeerArray = append(models.BeerArray, models.BeerItem{Id: 1, Name:"Golden", Brewery:"kross",Country:"Chile", Price:10.5, Currency:"USD" })

    r, _ := http.NewRequest("GET", "/beers/beerID", nil)
    w := httptest.NewRecorder()

    vars := map[string]string{
        "beerID": "1", //el ID lo creo en el momento del test por lo cual si existira
    }

    r = mux.SetURLVars(r, vars)

    SearchBeerById(w, r)

    assert.Equal(t, http.StatusOK, w.Code)

	//t.Error(w.Body.String())
    //assert.Equal(t, []byte("beerID"), w.Body.Bytes())
}
////Valida el retorno de cuando si encuentra un registro codigo 404
func TestGetBeerByIdNotExits(t *testing.T) {
	t.Parallel()

    r, _ := http.NewRequest("GET", "/beers/beerID", nil)
    w := httptest.NewRecorder()

    vars := map[string]string{
        "beerID": "8",//un numero que en el test no existe es 8
    }

    r = mux.SetURLVars(r, vars)

    SearchBeerById(w, r)

    assert.Equal(t, http.StatusNotFound, w.Code)
}