# Prueba tecnica Falabella
Desarrollo de APIs 
Bender es fanático de las cervezas y quiere tener un registro de todas las cervezas que prueba y como calcular el precio que necesita para comprar una caja de algún tipo especifico de cervezas. Para esto necesita una API REST con esta información que posteriormente compartirá con sus amigos.

## Comenzando 🚀

_Estas instrucciones te permitirán obtener una copia del proyecto en funcionamiento en tu máquina local para propósitos de desarrollo y pruebas._

Mira **Deployment** para conocer como desplegar el proyecto.


### Pre-requisitos 📋

Instalacion de Docker

### Instalación 🔧

Primeramenente debo de clonar el proyecto y ejecutar el siguiente comando para crear mi imagen de docker en base al archivo docker-compose utilizando el -d para que este levantado el servicio como un demon
```
docker-compose up -d
```

Si deseo visualizar con la siguiente instruccion

```
docker-compose ps
```

para realizar las validaciones puedo utilizar postman siguiendo las indicaciones de el archivo orignal de la prueba
https://bitbucket.org/lgaetecl/microservices-test/src/master/openapi.yaml  donde puedo validar como debo ejecutar y que respuestas esperar


## Ejecutando las pruebas ⚙️

Se desarrollaron exactamente los mismos EndPoint solicitados que son


GET /beers --Consulta de todas las cervezas


GET /beers/{BeerId} --aqui debo colocar el ID que quiero consultar

POST /beers - debo colocar en el body en formato json los valores segun el archivo de especificacion de swanger version 3.0 

GET /beers/{BeerId}/boxprice --aqui debo colocar 2 variables tipo query que son la moneda que deseo pagar, y la cantidad que deseo validar por default es 6


### 89% de cobertura de código 🔩

Se realizo las validaciones en con GO test y se genero un archivo llamado cover.out  con la siguiente insutruccion 
se puede tener el total de cobertura de proyecto en este caso estos fueron los valores devueltos por cada end-point

github.com/baemestrada-source/falabella/handlers/addBeers.go:10:                AddBeers                100.0%


github.com/baemestrada-source/falabella/handlers/boxBeerPriceById.go:13:        BoxBeerPriceById        82.4% 


github.com/baemestrada-source/falabella/handlers/searchBeerById.go:12:          SearchBeerById          100.0%

github.com/baemestrada-source/falabella/handlers/searchBeers.go:9:              SearchBeers             100.0%

Total:                                                                                                  89.8% 


```
go tool cover -func cover.out
```

### api de consulta de tasa de cambio ⌨️

se utilizo la api se extrajo de https://currencylayer.com/ para la consulta de tasa de cambio
se logro en base a la moneda de USD que es el unico source autorizado de forma gratuita de esta lista devuelta fue posible calcular la moneda y valores correspodientes segun el requerimiento



```
https://api.currencylayer.com/live?access_key=YOUR_ACCESS_KEY
```

## Rest en Heroku  📦

Siempre el proyecto para que si no se desea montar localmente pueden consultar a las siguientes rutas

## Construido con 🛠️

Se utilizo lenguaje GO y algunas librerias como 

github.com/gorilla/mux v1.8.0  -- sirve para facilitar el ruteo de mis end-point 


github.com/stretchr/testify v1.7.0 -- sirve para utilizar assert en el desarrollo de testeos y evitar hacer IF por cada validacion


el resto es propio de GO


## Docker y Dockerfile 🖇️

Se dejo en el repositorio los archivos tanto el DockerFile como el docker-compose.yml el unico servicio fue la web
ya que se manejo la data en memoria teniendo en cuenta que no era requisito un servicio de base de datos y para mayor rapides

## Autores ✒️

Byron Arturo Estrada Moreira

## Licencia 📄
OpenSource