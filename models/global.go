package models

  //BeerArray estructura donde se llena la informacion en memoria
  var BeerArray []BeerItem
  
  //Acces Key que me devolvio la API para tipo de cambio
  const Access_key = "a88c3af250d29d5d2c77e71066c27a92" 
  
  //TpcIn aqui se guardara el tipo de cambio USD versus la moneda que se desea pagar
  var TpcIn float32 
  
  //TpcDb aqui se guardara el tipo de cambio USD versus la moneda que tiene el precio en mi data o array
  var TpcDb float32