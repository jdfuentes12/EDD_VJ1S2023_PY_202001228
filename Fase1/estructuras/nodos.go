package estructuras

//Nodo de los empleados
type Nodo_Empleado struct {
	id        string
	nombre    string
	cargo     string
	password  string
	siguiente *Nodo_Empleado
}

//Nodo de los clientes
type Nodo_Cliente struct {
	id        string
	nombre    string
	siguiente *Nodo_Cliente
}

//Nodo de las imagenes
type Nodo_Imagen struct {
	imagen    string
	capa      string
	anterios  *Nodo_Imagen
	siguiente *Nodo_Imagen
}

//Codigo de los nodos para la Cola
type Nodo_Cola struct {
	id        string
	nombre    string
	siguiente *Nodo_Cola
}

//Codigo de los nodos para la Pila
type Nodo_Pila struct {
	id        string
	imagen    string
	siguiente *Nodo_Pila
}

type NodoMatriz struct {
	Siguiente *NodoMatriz
	Anterior  *NodoMatriz
	Abajo     *NodoMatriz
	Arriba    *NodoMatriz
	PosX      int
	PosY      int
	Color     string
}
