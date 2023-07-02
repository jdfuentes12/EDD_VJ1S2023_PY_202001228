
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
Le damos la bievenida al sistema por medio de una pagina web en donde se podra iniciar sesion tanto como empleado y administrador.<br>
### Administrador
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/bfbd1118-4c1b-4ba4-8f9b-31549063b41a" width="450"><br>

<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/5474fc0b-d25a-4707-9451-04a24b0cb8b8" width="450"><br>

<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/0dd4f628-1a61-43f4-8a69-2d1976535ebb" width="450"><br>

<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/b81908d9-f0ae-45ff-bb10-7135b0335a86" width="450"><br>

<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/c98c2d72-e540-4a17-9914-f4889dd43bc2" width="450"><br>

### Empleado
Menu Empleado<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/0318e842-cfa6-4ae8-9d82-c0847d431790" width="450"><br>
Aplicar filtros<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/4c02e943-2312-4d3b-9a5f-56239cb4e617" width="450"><br>

Elegir filtro<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/df72ade3-2a55-403a-b468-64cafdf2cff7" width="450"><br>

Imagen con filtro<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/de463526-2312-474f-b93b-6f22ab68b211" width="450"><br>

imagen generada<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/09267092-dbf9-468a-a872-3886c0d984d1" width="450"><br>

Generar Factura<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/7b162a6b-7064-440f-89c9-dfa6f9474b93" width="450"><br>

<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/f8f0cd31-eab2-494c-b1cd-0d8732dfe2ac" width="450"><br>

Ver facturas<br>
<img align='center' src="https://github.com/jdfuentes12/EDD_VJ1S2023_PY_202001228/assets/88565998/9d8335ea-367a-4239-9cfe-af68d73a9b22" width="450"><br>







