## Manual Tecnico
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
![image](https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/8227229b-5ec7-4ba7-8450-c0655c2ba574)
<br>


## Manual Usuario
