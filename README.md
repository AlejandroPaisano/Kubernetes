# Kubernetes

# Introduccion
En esta actividad se plantea el usar el servicio de kubernetes para poder simular el crear un servidor y poder administrar multiples contenedores desde una misma ubicacion, ademas de poder manejar  como estos interactuan entre si. Para poder hacer esto, haremos uso del servicio de digital ocean, el lenguaje de go para el codigo dentro del contenedor y dcocker para crear los contenedores.

# Desarrollo
Primero diseñaremos un ejemplo de codigo que guardaremos para poder contener dentro de docker. En este caso, usaremos go y las bibliotecas de fmt y net/http para poder crear una pagina web de ejemplo en la que introduciremos un saludo para el usuario.
![image](https://github.com/AlejandroPaisano/Kubernetes/assets/91223611/592508e0-7e5f-4249-af7a-1feac29e70e8)

En este caso, nuestro codigo corre en el puerto 3001. Hecho esto, ahora debemos crear un archivo dockerfile donde generaremos nuestra imagen para luego guardar dentro de un contenedor y poder hacer un push de nuestra imagen y guardarla dentro de nuestra cuenta de docker para poder usarla en el futuro.

![image](https://github.com/AlejandroPaisano/Kubernetes/assets/91223611/1b3d6815-7f67-48f9-a36d-792f242b65cc)

Hecho esto, crearmos un cluster para kubernetes en DigitalOcean. Con esto hecho, copiaremos el archivo config para alimentarselo al servicio de kubecl para guardar las direcciones en las que guardaremos nuestras imagenes para montarlas y controlarlas desde kubernetes.

![image](https://github.com/AlejandroPaisano/Kubernetes/assets/91223611/8f8da924-618e-4c64-9609-856acb6242ae)

Por otro lado, para que nuestras aplicaciones de kuberentees funcione. Debemos generar un archivo  de deployment tipo yml, para poder guardar las funciones bnsicas que debe realizar el kubernete, antes de iniciar. Hecho el deployment, debemos de realizar un archivo services, donde le indicaremos al kubernete los servicios que debe levantar, y el como los debe levantar en caso de ser necesario

![image](https://github.com/AlejandroPaisano/Kubernetes/assets/91223611/f4cb81e4-8714-44f2-8f48-1011289a7190)
![image](https://github.com/AlejandroPaisano/Kubernetes/assets/91223611/c1d7ec00-ad0a-439f-be76-21dbf8a224a8)

con esto hecho, deberemos de montar ambos archivos con su respectivo comando: kubectl apply -f "nombre del archivo".yml, aplicando primero el archivo deployment y luego el services. Despues vamos a usar el comando kubectl get services para poder verificar si nuestra imagen se monto de forma correcta.

![image](https://github.com/AlejandroPaisano/Kubernetes/assets/91223611/d0a19f4e-0153-498d-84bc-386e6b4aa900)

Si por alguna razon  en lugar de tener una direccion ip en "external IP" tenemos la leyenda "pending" significa que nuestro kubernete aun se esta montando y ddeberemos esperar un tiempo hasta que aparezca con su direccion ip. Una vez que la direccion ip aparece, podemos escribir esa direccion ip en nuestro buscador y el contenedor hara su trabajo.

![image](https://github.com/AlejandroPaisano/Kubernetes/assets/91223611/821c7340-c157-434b-b37d-cca7f4023a81)


# Conclusion
Los kubernetes son bastante curiosos de trabajar dado que estos representan una forma mas real de poder crear aplicaciones de microservicios en la web, pues estos ya estan montados en servidores en la nube. Ademas de eso, los kubernetes nos facilitan el creear este codigo dado que solo toman las imagenes que tengamos en nuestro dockerhub, por lo que podemos comprobar que estops funcionan de antemano y no lestarnos en tener que asegurarnos que este codigo funcione de forma perfecta en un servidor.

Quiza su mayor desventaja, es que como opera en servidores, todos los servicios de kubernetes cobnran de una forma u otra.
