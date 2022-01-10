package models

// BeerItem defines model for BeerItem.
  type BeerItem struct {
	Id       int     `json:"Id"`
	Name     string  `json:"Name"`
	Brewery  string  `json:"Brewery"`
	Country  string  `json:"Country"`
	Price    float32 `json:"Price"`
	Currency string  `json:"Currency"`
}
