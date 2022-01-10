package handlers

import ( 
	"testing"
	"bytes"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/baemestrada-source/falabella/models"
)

func TestAddBeer(t *testing.T) {

	//Json de entrada para prueba de ingreso
	var json = []byte(`{
		"Id": 3,
		"Name": "Wurst",
		"Brewery": "Eisbock",
		"Country": "Alemania",
		"Price": 3.5,
		"Currency": "EUR"
		}`)
	
	req, err := http.NewRequest("POST", "/beers", bytes.NewBuffer(json))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(AddBeers)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	status := response.Code
	assert.Equal(t, status, http.StatusCreated)

}

func TestAddBeerIdExistente(t *testing.T) {
	//Limpio cualquier valor que traiga de array de otro test
	models.BeerArray = nil
	//En test creo el ID 1 para que cuando intente crearlo por la API debe de regresar 409
	models.BeerArray = append(models.BeerArray, models.BeerItem{Id: 1, Name:"Golden", Brewery:"kross",Country:"Chile", Price:10.5, Currency:"USD" })

	//Json de entrada para prueba de ingreso con el mismo ID
	var json = []byte(`{
		"Id": 1,
		"Name": "Golden1",
		"Brewery": "Eisbock",
		"Country": "Alemania",
		"Price": 3.5,
		"Currency": "EUR"
		}`)
	
	req, err := http.NewRequest("POST", "/beers", bytes.NewBuffer(json))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(AddBeers)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	status := response.Code
	assert.Equal(t, status, 409) //Compara que el valor de retorno sea el 409

}

func TestAddBeerBodyInvalido(t *testing.T) {
	//Json invalido
	var json = []byte(``)
	
	req, err := http.NewRequest("POST", "/beers", bytes.NewBuffer(json))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(AddBeers)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	status := response.Code
	//Compara que el valor de retorno sea el 400 ya que el json es invalido
	assert.Equal(t, status, http.StatusBadRequest) 
}