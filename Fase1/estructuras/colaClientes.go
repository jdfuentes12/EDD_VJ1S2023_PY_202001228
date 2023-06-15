package estructuras

import (
	"fmt"
	"strconv"
)

type ListaColaClientes struct {
	Inicio   *Nodo_Cola
	Final    *Nodo_Cola
	Longitud int
}

func (l *ListaColaClientes) Encolar(id string, nombre string) {
	nuevoCliente := &Nodo_Cola{id: id, nombre: nombre, siguiente: nil}

	if l.Inicio == nil {
		l.Inicio = nuevoCliente
		l.Final = nuevoCliente
		l.Longitud++
	} else {
		l.Final.siguiente = nuevoCliente
		l.Final = nuevoCliente
		l.Longitud++
	}
}

func (l *ListaColaClientes) Desencolar() {
	if l.Inicio == nil {
		fmt.Println("Cola Vacia")
	} else {
		l.Inicio = l.Inicio.siguiente
		l.Longitud--
	}
}

func (l *ListaColaClientes) RetornarID() string {
	aux := l.Inicio
	return aux.id
}

func (l *ListaColaClientes) RetornarNombre() string {
	aux := l.Inicio
	return aux.nombre
}

func (l *ListaColaClientes) MostrarCola() {
	if l.Inicio == nil {
		fmt.Println("Cola Vacia")
	}
	aux := l.Inicio
	for aux != nil {
		fmt.Println(aux.id, aux.nombre)
		aux = aux.siguiente
	}
}

func (l *ListaColaClientes) GenererarGraphvizCola() string {
	texto := "digraph ListaClientes {\n"
	contador := 0
	texto += "node [shape=record, height = 0.5];\n"
	texto += "rankdir = LR;\n"

	aux := l.Inicio
	if aux != nil {
		for aux != nil {
			texto += strconv.Itoa(contador) + "[label = \"" + aux.id + ", " + aux.nombre + "\"];\n"
			contador++
			aux = aux.siguiente
			if aux == l.Inicio {
				break
			}
		}
		for i := 0; i < contador-1; i++ {
			texto += strconv.Itoa(i) + "->" + strconv.Itoa(i+1) + ";\n"
		}
	} else {
		texto += "ListaClientesVacia[label = \"Lista de Clientes Vacia\"];\n"
	}
	texto += "}"
	return texto
}
