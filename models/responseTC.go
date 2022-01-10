package models

//ResponseTC estructura para leer y decodificar el json de la API de tipo de cambio
type ResponseTC struct {
	  Success   bool   `json:"success"` //si devuelve lo correcto es true
	  Quotes map[string]float32 `json:"quotes"` //devuelve un array en forma de mapa de mis claves de monedas y valores
}