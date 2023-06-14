package estructuras

import (
	"fmt"
	"strconv"
)

//Estructura de la lista simple de empleados
type ListaEmpleados struct {
	Inicio *Nodo_Empleado
	Longitud int
}

func (l *ListaEmpleados) estaVacia() bool {
	return l.Longitud == 0
}


func (l *ListaEmpleados) Insertar(id string, nombre string, cargo string, password string) {
	if l.estaVacia() {
		l.Inicio = &Nodo_Empleado{id: id, nombre: nombre, cargo: cargo, password: password, siguiente: nil}
		l.Longitud++
	}else{
		aux := l.Inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = &Nodo_Empleado{id: id, nombre: nombre, cargo: cargo, password: password,siguiente: nil}
		l.Longitud++
	}
}

func(l *ListaEmpleados) MostrarLista(){
	aux := l.Inicio
	fmt.Println("Lista de Empleados")
	for aux != nil {
		fmt.Println("ID: ", aux.id, " Nombre: ", aux.nombre, " Cargo: ", aux.cargo, " Password: ", aux.password)
		aux = aux.siguiente
	}
}

func(l *ListaEmpleados) GenererarGraphvizEmpleado() string{
	texto := "digraph ListaEmpleados {\n"
	contador := 0
	aux := l.Inicio
	if aux != nil {
		for aux != nil {
			texto += strconv.Itoa(contador) + "[label = \"" + aux.nombre + ", " + aux.cargo + "\"];\n"
			contador++
			aux = aux.siguiente
		}
		for i := 0; i < contador-1; i++ {
			texto += strconv.Itoa(i) + "->" + strconv.Itoa(i+1) + ";\n"
		}
		
	}else{
		texto += "ListaEmpleadosVacia[label = \"Lista de Empleados Vacia\"];\n"
	}
	texto += "}"
	return texto
}