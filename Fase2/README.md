
## *Manual Técnico*
El programa esta realizado con el lenguaje de Go en su version "go1.20.1" que se utiliza para backennd, con la herramienta de "graphviz" en su version "8.0.5". Y se implemento lo que es frontend en React.

### Estructura Utilizadas 
* Lista simple enlazada
* Arbol AVL
* BlackChain
* Grafo
* Matriz Dispersa
* Cola
* Tabla Hash
<br>
En el proyecto se usarion varios modulos los cuales contendra diferentes metodos, se usaron para realizar una instrucción especifica.
Tiene dos capetas iniciales que son "Backend" y "frontend".
<br>
En la carpeta de Backend hay diferentes modulos y acciones que realiza. Por ejemplo el de iniciar sesión, la carga masiva, etc. Contendra una carpeta llamada "estructuras" la cual contiene los nodos y constructores de las estructuras andes descritas. Tiene una carpeta llamada "Rerportes" donde estaran las imagenes correspondiente a cada una de las estructuras realizadas en el proyecto. Tambien esta la carpeta llamada "Vistas" la cual contiene el css y hmtl de las imagenes ingresadas. <br>
Tambien en la carpeta "frontend" que es de React la cual se crea por medio de un codigo en la terminal. Ademas este contiene una carpeta que se llama "src"  en donde ser realizaran todas las acciones que y las vistas estan en las carpetas "components" y "css" los cuales nos permitiran ver e interactuar con el usuario y administrador. 
A continuacion se mostrara una imagen de los archivos usados en el proyecto.

<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/fbe3ea30-86ac-4f2b-8b6b-fea197dcd69c" width="200"><br>

### Servidor
Aca estara el metodo main el cual es el que hace todas las acciones que se soliciten en el backend. Se mostrara una imagen de todas las importaciones que se hicieron segun las estructuras usadas.<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/6a019b57-ea4e-4311-be13-48a5822c8a01" width="400"><br>
En este caso se uso Go como servidor y realizar las diferentes acciones que se solicitara segun lo requiere el usuario. A continuacion se le presentara un fracmento de codigo el cual cosite en un metodo "POST" el cual se requiere la validacion de los usuarios tanto administrador como empleados. Y se retorna un "status" con un valor de 400 el cual consiste que se inicio correctamente, de lo contrario enviara un 404 el cual indicara que que la contreseña o usuario son incorrectos.
<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/db9d4577-b853-4675-a2e0-13f5637158a5" width="450"><br>

### Lectura de csv
Se realiza una lectura de los archivos con formato csv los cuales contendran los datos de clientes, empleados y las diferentes imagenes que se daran. 
<br>
Por medio de la importacion de "encoding/csv" se pudo realizar la lectura del archivo con extension de "csv" el cual fue establecido para el proyecto. Para ellos se verificara por medio de if si hay algun error a la hora de leer el archivo, si lo hay nos lo dira y si no seguira leyendo el archivo. Por medio de otro if verificara si hay un archivo con ese nombre y si no lo hay nos mostrara error. Despues se procede a leer por linea cada unos de los datos que contiene el archivo por medi de un ciclo for, el cual se iran almacenando en diferentes variables cada dato. Despues se hara la carga a las diferentes estructuras dependiendo de que csv se lea. Se lo mostrara un ejemplo con la lectura de las imagenes.
<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/8227229b-5ec7-4ba7-8450-c0655c2ba574" width="450"><br>
<br>

### Matriz
En la matriz tenemos diferentes clases para las inserciones de casa uno de los datos para graficar lo que contiene la matriz. Se mostrara como leer el inicial el cual contiene la lectura del primer csv que se almanezara y dara las configuraciones y ademas se cargaran todas las capas a la matriz que contiene la imagen establecida o enviada. <br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/56c113f8-7c81-418f-8bc8-db6127ae6576" width="450"><br>
<br>
Se mostrara una pequeña parte de la insercion a la matriz dependiendo del caso que se establezca por si el archivo .csv viene desordenado. <br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/0dc4fead-e607-4c4b-bd8c-9c07b49d130e" width="450"><br>
<br>



## *Manual Usuario*
EDD Creative es un programa que se usara para generar imagenes pixeles. Con el objetivo de que se genere la imagen en un documetno html y con su respectivo css.
<br>
A la hora de correr el programa se mostrara un menu para iniciar Sesion o salir de la aplicacion. <br>
Nota: este menu solo permitira iniciar en modo administrador, mas adelante se permitira para los empleados que esten en el sistema.<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/6635d4f0-6f72-4282-8cd4-f3cd1a900892" width="350"><br>
Para poder acceder al modo administrador se usa el usuario "admin_202001228" y la contraseña es "admin" y el cual generara el menu administrado y las acciones que se pueden hacer.<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/6afd7d63-1cfe-4c1c-ba02-4660c17a321d" width="350"><br>
A la  hora de elegir cualquiera de las primeras 4 opciones le pedira la ruta del archivo para cargarlo en el sistema y poder hacer las operaciones necesarias y almacenarlos en la estructura. <br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/8fc46aa7-39f0-49ff-b755-c3ae160bf601" width="350"><br>
Y si se quiere ver el reporte de todas las estructuras que se encuentran en el sistema se elige el numero "5" el cual generara todos los reportes de cada una de las estructuras.<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/0810b536-e940-432f-a695-08ca9a2c7b62" width="350"><br>
Al salir del menu administrador se le desplegara otro menu, donde podra elegir si quiere seguir como administrado o entrar como empleado al sistema. <br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/f1a6447e-eb75-492f-b028-0a029ba41da1" width="350"><br>
Si elegimos la opccion de entrar como un usuario, nos pedira  nuestras credenciales, las cueales tenemos que colocar. Si no son correctas repetira el proceso y si son correctas entrará al sistema. <br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/9fde338d-3dde-444c-a88b-5f699c1e9468" width="350"><br>
Como se puede ver en la imagen anterior solo tiene 3 opciones.  Podra observar las imagenes disponibles en el sistema o puede proceder a atender al cliente que se encuentra en la cola.
