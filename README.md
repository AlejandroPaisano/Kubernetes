# Kubernetes

# Introduccion
En esta actividad se plantea el usar el servicio de kubernetes para poder simular el crear un servidor y poder administrar multiples contenedores desde una misma ubicacion, ademas de poder manejar  como estos interactuan entre si. Para poder hacer esto, haremos uso del servicio de digital ocean, el lenguaje de go para el codigo dentro del contenedor y dcocker para crear los contenedores.

# Desarrollo
Primero diseñaremos un ejemplo de codigo que guardaremos para poder contener dentro de docker. En este caso, usaremos go y las bibliotecas de fmt y net/http para poder crear una pagina web de ejemplo en la que introduciremos un saludo para el usuario.
![image](https://github.com/AlejandroPaisano/Kubernetes/assets/91223611/592508e0-7e5f-4249-af7a-1feac29e70e8)

En este caso, nuestro codigo corre en el puerto 3001. Hecho esto, ahora debemos crear un archivo dockerfile donde generaremos nuestra imagen para luego guardar dentro de un contenedor y poder hacer un push de nuestra imagen y guardarla dentro de nuestra cuenta de docker para poder usarla en el futuro.
![image](https://github.com/AlejandroPaisano/Kubernetes/assets/91223611/1b3d6815-7f67-48f9-a36d-792f242b65cc)

Hecho esto, crearmos un cluster para kubernetes en DigitalOcean. Con esto hecho, copiaremos el archivo config para alimentarselo al servicio de kubecl para guardar las direcciones en las que guardaremos nuestras imagenes para montarlas y controlarlas desde kubernetes.

Por otro lado, apra que nuestras apliccaiones de kuberentees funcione. Debemos generar un archivo  de deployment tipo yml, para poder guardar las funciones bnsicas que debe realizar el kubernete, antes de iniciar. Hecho el deployment, debemos de realizar un archivo services, donde le indicaremos al jubernete los servicios que debe levantar, y el como los debe levantar en caso de ser necesario
