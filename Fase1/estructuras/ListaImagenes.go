package estructuras

import (
	"strconv"
)


type ListaImagenes struct {
	Primero *Nodo_Imagen
	Ultimo *Nodo_Imagen
}

func (l *ListaImagenes) Insertar(imagen string, capa string){
	nuevoNodo := &Nodo_Imagen{imagen: imagen, capa: capa, anterios: nil, siguiente: nil}
	
	if l.Primero == nil {
		l.Primero = nuevoNodo
		l.Ultimo = nuevoNodo
	} else {
		nuevoNodo.anterios = l.Ultimo
		l.Ultimo.siguiente = nuevoNodo
		l.Ultimo = nuevoNodo
	}
}

func (l *ListaImagenes) MostrarLista() {
	aux := l.Primero

	for aux != nil {
		println(aux.imagen, aux.capa)
		aux = aux.siguiente
	}
}

func (l *ListaImagenes) GenererarGraphvizImagenes() string{
	texto := "digraph ListaEmpleados {\n"
	contador := 0
	aux := l.Primero
	if aux != nil {
		for aux != nil {
			texto += "nodo"+ strconv.Itoa(contador) + "[label = \"" + aux.imagen + ", " + aux.capa + "\"];\n"
			contador++
			aux = aux.siguiente
		}
		for i := 0; i < contador-1; i++ {
			texto += "nodo"+strconv.Itoa(i) + "->" +"nodo"+ strconv.Itoa(i+1) + ";\n"
			texto += "nodo"+strconv.Itoa(i+1) + "->" +"nodo"+ strconv.Itoa(i) + ";\n"
		}
		
	}else{
		texto += "ListaEmpleadosVacia[label = \"Lista de Empleados Vacia\"];\n"
	}
	texto += "}"
	return texto
}