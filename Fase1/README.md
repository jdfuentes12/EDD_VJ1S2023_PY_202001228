## Manual Técnico
El programa esta realizado con el lenguaje de Go en su version "go1.20.1", con la herramienta de "graphviz" en su version "8.0.5".

### Estructura Utilizadas 
* Lista simple enlazada
* Lista doblemente enlazada
* Lista Circular
* Matriz dispersa
* Pila
* Cola

### Modulos y carpetas usadas
En este modo se cargaran diferentes archivos de entrada de formato "csv" con el objetivo de almacenarlos en las estructuras que se realizan en Go.
<br>
En el proyecto se usarion varios modulos los cuales contendra diferentes metodos, se usaron para realizar una instrucción especifica. Por ejemplo el de iniciar sesión, la carga masiva, etc. Contendra una carpeta llamada "estructuras" la cual contiene los nodos y constructores de las estructuras andes descritas. Tiene una carpeta llamada "graficas" donde estaran las imagenes correspondiente a cada una de las estructuras realizadas en el proyecto.
A continuacion se mostrara una imagen de los archivos usados en el proyecto.
<br>
![image](https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/26ca72d1-7ffb-4fa7-90c2-dd26f1f74b49)
<br>
### Modulo main
Este es el modulo el cual contiene parte de las acciones necesarias para realizar ingresar a la diferentes estructuras e importaciones necesarias que se utilizaron.
<br>
![image](https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/e61e8643-5632-42df-968e-6886044fe723)
<br>
### Lectura de csv
Se realiza una lectura de los archivos con formato csv los cuales contendran los datos de clientes, empleados y las diferentes imagenes que se daran. 
<br>
Por medio de la importacion de "encoding/csv" se pudo realizar la lectura del archivo con extension de "csv" el cual fue establecido para el proyecto. Para ellos se verificara por medio de if si hay algun error a la hora de leer el archivo, si lo hay nos lo dira y si no seguira leyendo el archivo. Por medio de otro if verificara si hay un archivo con ese nombre y si no lo hay nos mostrara error. Despues se procede a leer por linea cada unos de los datos que contiene el archivo por medi de un ciclo for, el cual se iran almacenando en diferentes variables cada dato. Despues se hara la carga a las diferentes estructuras dependiendo de que csv se lea. Se lo mostrara un ejemplo con la lectura de las imagenes.
<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/8227229b-5ec7-4ba7-8450-c0655c2ba574" width="450"><br>
<br>

### Agregar Imagenes
Se realizara un ejemplo de agregar un dato en este caso una imagen a la estructura de una lista que en este caso es una doblemente enlazada, la cual contendra el  nombre de la imagen y la cantidad de capas que tiene. 
<br>
Se le mostrara el nodo de la imagen y como es una doblemente enlazada contiene 2 apuntadores, que es haciea delante y hacia atras.
<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/ff2a0dc6-a622-4ace-9c9d-4c8b3942dabc" width="450"><br>
<br>
Despues de tener el nodo, procedemos a insertar algun valor a nuestra lista doblemente enlazada. Se verifica si no hay algun valor antesmente ingresado y si no lo hay se coloca como el primer dato ahora si ya hay un dato en nuestra lista doblemente enlazada si procede a colocar el nuevo valor como segundo, pero siempre tratando de apuntarse entre ellos y como ultimo, el valor ingresado tiene que ir apuntando al primer valor, y se sigue haciendo este mismo proceso para ingresar cualquier dato.
<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/6b74d27f-7399-4266-8560-2d75d334e5e2" width="750"><br>

### Rerporte en Ghaphviz
Para verificar si tenemos datos en nustres estructuras usamos los reportes en graphviz, que despues se convierten en un formato jpg para poder  visualizarlos de mejor manera. A continuacion se lo mostrara el codigo para graficas los diferentes reportes de las estructuras. Realizamos un contador y una variable con el nombre de "texto" que contendra todo lo almacenado para generar el archivo .dot al igual se hace un parseo a string para colocar el contador y poder escribirlo y al macenarlo en la variable "texto". A continuacion se le musestra el codigo.
<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/636e8362-f08b-4a6c-aeab-03e0a18a64bd" width="650"><br>
Y se le mostrara el codigo que se usa pra la conversion de "dot" a "jpg".
<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/2c33d12c-6e89-4f31-8d49-92c767b45b22" width="650"><br>
Imagen del reporte
<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/88aae8ca-9d29-48c8-aac7-24e5dad2be73" width="100"><br>


## Manual Usuario
