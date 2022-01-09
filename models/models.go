package models


//ResponseTC estructura para leer y decodificar el json de la API de tipo de cambio
type ResponseTC struct {
	Success   bool   `json:"success"` //si devuelve lo correcto es true
	Quotes map[string]float32 `json:"quotes"` //devuelve un array en forma de mapa de mis claves de monedas y valores
  }
  
  // BeerBox defines model for BeerBox.
  type BeerBox struct {
	  PriceTotal float32 `json:"Price Total,omitempty"`
  }
  
  // BeerItem defines model for BeerItem.
  type BeerItem struct {
	  Id       int     `json:"Id"`
	  Name     string  `json:"Name"`
	Brewery  string  `json:"Brewery"`
	  Country  string  `json:"Country"`
	  Price    float32 `json:"Price"`
	Currency string  `json:"Currency"`
  }
  
  //BeerArray estructura donde se llena la informacion en memoria
  var BeerArray []BeerItem
  
  //Acces Key que me devolvio la API para tipo de cambio
  const Access_key = "a88c3af250d29d5d2c77e71066c27a92" //clave para la API de tipo de cambio
  
  //TpcIn aqui se guardara el tipo de cambio USD versus la moneda que se desea 
  var TpcIn float32 
  
  //TpcDb aqui se guardara el tipo de cambio USD versus la moneda que tiene el precio en mi data
  var TpcDb float32 
  