package Lista

// Estructura de la lista simple de empleados
type ListaEmpleados struct {
	Inicio   *Nodo_Empleado
	Longitud int
}

func (l *ListaEmpleados) estaVacia() bool {
	return l.Longitud == 0
}

func (l *ListaEmpleados) Insertar(id string, nombre string, cargo string, password string) {
	if l.estaVacia() {
		l.Inicio = &Nodo_Empleado{id: id, nombre: nombre, cargo: cargo, password: password, siguiente: nil}
		l.Longitud++
	} else {
		aux := l.Inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = &Nodo_Empleado{id: id, nombre: nombre, cargo: cargo, password: password, siguiente: nil}
		l.Longitud++
	}
}

func (l *ListaEmpleados) Buscar(username string, password string) bool {
	aux := l.Inicio
	for aux != nil {
		if aux.id == username && aux.password == password {
			return true
		}
		aux = aux.siguiente
	}
	return false
}
