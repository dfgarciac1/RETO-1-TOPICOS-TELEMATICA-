# RETO-1-TOPICOS-TELEMATICA-


## README 
Este proyecto fue realizado para la materia de tópicos de telematica para este mismo se tiene el uso de los sockets y después de realizar la conexión el uso de http por medio de la url `ws:///localhost:PORT/html` 
el cual recibe por parámetro la url del sitio al cual se desea obtener su HTML  y los header  que se dan en la respuesta por medio de TCP

  ## MANERA DE COMO SE ENVÍA LA URL 
 - [ ] **ws:///localhost:PORT** (CONEXIÓN)
 - [ ] **ws:///localhost:PORT/html** (PETICIONES)



![](https://i.imgur.com/ywUOhOf.png)

## REQUISITOS

 1. Primeramente tener instalado Go en caso que no se tenga dentro de los archivos esta el   programa compilado el cual se podrá ejecutar de la siguiente manera 

**EN CASO DE ESTAR EN LINUX**


![](https://i.imgur.com/680oP3V.png)


**EN CASO DE ESTAR EN WINDOWS O MAC**
![](https://i.imgur.com/u3F9uiq.png)

## FUNCIONAMIENTO

 - Cuando se ejecute el programa y sea correcto el proceso este deberá enviar la siguiente respuesta 

![](https://i.imgur.com/WgVwtuN.png)

 - Además de esto sera capaz de parsear los elementos del HTML más comunes y guardarlos como un txt  con su tag como principio en el txt.
## RECOMENDACIONES 
 - [ ] Para mayor facilidad cuando se este probando se recomienda el uso de herramientas como **POSTMAN** debido a que facilita visualmente las peticiones y la conexión al socket
