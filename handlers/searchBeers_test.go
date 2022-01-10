package handlers

import ( 
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

//Valida el GET donde trae todo que retorne un codigo de OK
func TestGetBeerAll(t *testing.T) {
	req, err := http.NewRequest("GET", "/beers", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(SearchBeers)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	status := response.Code
	assert.Equal(t, status, http.StatusOK)
}