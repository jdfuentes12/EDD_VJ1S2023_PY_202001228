package estructuras

import (
	"fmt"
	"strconv"
)

type ListaClientes struct {
	Primero  *Nodo_Cliente
	Ultimo   *Nodo_Cliente
	Longitud int
}

func (l *ListaClientes) Insertar(id string, nombre string) {
	nuevoCliente := &Nodo_Cliente{id: id, nombre: nombre, siguiente: nil}

	if l.Primero == nil {
		l.Primero = nuevoCliente
		l.Ultimo = nuevoCliente
		nuevoCliente.siguiente = nuevoCliente
		l.Longitud++
	} else {
		nuevoCliente.siguiente = l.Primero
		l.Ultimo.siguiente = nuevoCliente
		l.Ultimo = nuevoCliente
		l.Longitud++
	}
}

func (l *ListaClientes) BuscarClienteId(id string) bool {
	aux := l.Primero
	for {
		if aux.id == id {
			return true
		}
		aux = aux.siguiente
		if aux == l.Primero {
			break
		}
	}
	return false
}

func (l *ListaClientes) MostrarLista() {
	if l.Primero == nil {
		fmt.Println("Lista Vacía")
	}
	aux := l.Primero
	for {
		fmt.Println(aux.id, aux.nombre)
		aux = aux.siguiente
		if aux == l.Primero {
			break
		}
	}
}

func (l *ListaClientes) GenererarGraphvizCliente() string {
	texto := "digraph ListaClientes {\n"
	contador := 0
	texto += "node [shape=record, height = 0.5];\n"
	aux := l.Primero
	if aux != nil {
		for aux != nil {
			texto += strconv.Itoa(contador) + "[label = \"" + aux.id + ", " + aux.nombre + "\"];\n"
			contador++
			aux = aux.siguiente
			if aux == l.Primero {
				break
			}
		}
		for i := 0; i < contador-1; i++ {
			texto += strconv.Itoa(i) + "->" + strconv.Itoa(i+1) + ";\n"
		}
		texto += strconv.Itoa(contador-1) + "->" + strconv.Itoa(0) + ";\n"
	} else {
		texto += "ListaClientesVacia[label = \"Lista de Clientes Vacia\"];\n"
	}
	texto += "}"
	return texto
}
