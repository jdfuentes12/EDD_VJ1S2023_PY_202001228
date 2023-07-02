package ArbolAVL

import (
	"Fase2/estructuras/ColaPedidos"
	"Fase2/estructuras/GenerarArchivos"
	"Fase2/estructuras/Peticiones"
	"math"
	"strconv"
)

type Arbol struct {
	Raiz *NodoArbol
}

func (a *Arbol) altura(raiz *NodoArbol) int {
	if raiz == nil {
		return 0
	}
	return raiz.Altura
}

func (a *Arbol) equilibrio(raiz *NodoArbol) int {
	if raiz == nil {
		return 0
	}
	return a.altura(raiz.Derecha) - a.altura(raiz.Izquierda)
}

func (a *Arbol) InsertarElemento(id_cliente int, nombre_imagen string) {
	nuevoNodo := &NodoArbol{Valor: &Peticiones.Pedido{Id_Cliente: id_cliente, Nombre_Imagen: nombre_imagen}}
	a.Raiz = a.insertarNodo(a.Raiz, nuevoNodo)
}

func (a *Arbol) rotacionI(raiz *NodoArbol) *NodoArbol {
	raiz_derecho := raiz.Derecha
	hijo_izquierdo := raiz_derecho.Izquierda
	raiz_derecho.Izquierda = raiz
	raiz.Derecha = hijo_izquierdo
	numeroMax := math.Max(float64(a.altura(raiz.Izquierda)), float64(a.altura(raiz.Derecha))) + 1
	raiz.Altura = 1 + int(numeroMax)
	numeroMax = math.Max(float64(a.altura(raiz_derecho.Izquierda)), float64(a.altura(raiz_derecho.Derecha)))
	raiz_derecho.Altura = 1 + int(numeroMax)
	raiz.Factor_Equilibrio = a.equilibrio(raiz)
	raiz_derecho.Factor_Equilibrio = a.equilibrio(raiz_derecho)
	return raiz_derecho
}

func (a *Arbol) rotacionD(raiz *NodoArbol) *NodoArbol {
	raiz_izquierdo := raiz.Izquierda
	hijo_derecho := raiz_izquierdo.Derecha
	raiz_izquierdo.Derecha = raiz
	raiz.Izquierda = hijo_derecho
	numeroMax := math.Max(float64(a.altura(raiz.Izquierda)), float64(a.altura(raiz.Derecha)))
	raiz.Altura = 1 + int(numeroMax)
	numeroMax = math.Max(float64(a.altura(raiz_izquierdo.Izquierda)), float64(a.altura(raiz_izquierdo.Derecha)))
	raiz_izquierdo.Altura = 1 + int(numeroMax)
	raiz.Factor_Equilibrio = a.equilibrio(raiz)
	raiz_izquierdo.Factor_Equilibrio = a.equilibrio(raiz_izquierdo)
	return raiz_izquierdo
}

func (a *Arbol) insertarNodo(raiz *NodoArbol, nuevoNodo *NodoArbol) *NodoArbol {
	if raiz == nil {
		raiz = nuevoNodo
	} else {
		if raiz.Valor.Id_Cliente > nuevoNodo.Valor.Id_Cliente {
			raiz.Izquierda = a.insertarNodo(raiz.Izquierda, nuevoNodo)
		} else {
			raiz.Derecha = a.insertarNodo(raiz.Derecha, nuevoNodo)
		}
	}
	numeroMax := math.Max(float64(a.altura(raiz.Izquierda)), float64(a.altura(raiz.Derecha)))
	raiz.Altura = 1 + int(numeroMax)
	balanceo := a.equilibrio(raiz)
	raiz.Factor_Equilibrio = balanceo
	/*Rotacion simple a la izquierda*/
	if balanceo > 1 && nuevoNodo.Valor.Id_Cliente > raiz.Derecha.Valor.Id_Cliente {
		return a.rotacionI(raiz)
	}
	if balanceo < -1 && nuevoNodo.Valor.Id_Cliente < raiz.Izquierda.Valor.Id_Cliente {
		return a.rotacionD(raiz)
	}
	if balanceo > 1 && nuevoNodo.Valor.Id_Cliente < raiz.Derecha.Valor.Id_Cliente {
		raiz.Derecha = a.rotacionD(raiz.Derecha)
		return a.rotacionI(raiz)
	}
	if balanceo < -1 && nuevoNodo.Valor.Id_Cliente > raiz.Izquierda.Valor.Id_Cliente {
		raiz.Izquierda = a.rotacionI(raiz.Izquierda)
		return a.rotacionD(raiz)
	}
	return raiz
}

func (a *Arbol) Graficar() {
	cadena := ""
	nombre_archivo := "./Reportes/arbolAVL.dot"
	nombre_imagen := "./Reportes/arbolAVL.jpg"
	if a.Raiz != nil {
		cadena += "digraph arbol{ \n"
		cadena += a.retornarValoresArbol(a.Raiz, 0)
		cadena += "}"
	}
	GenerarArchivos.CrearArchivo(nombre_archivo)
	GenerarArchivos.EscribirArchivo(cadena, nombre_archivo)
	GenerarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}

func (a *Arbol) retornarValoresArbol(raiz *NodoArbol, indice int) string {
	cadena := ""
	numero := indice + 1
	if raiz != nil {
		cadena += "\""
		cadena += strconv.Itoa(raiz.Valor.Id_Cliente)
		cadena += "\" ;\n"
		if raiz.Izquierda != nil && raiz.Derecha != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += strconv.Itoa(raiz.Valor.Id_Cliente)
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Izquierda, numero)
			cadena += "\""
			cadena += strconv.Itoa(raiz.Valor.Id_Cliente)
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Derecha, numero)
			cadena += "{rank=same" + "\"" + strconv.Itoa(raiz.Izquierda.Valor.Id_Cliente) + "\"" + " -> " + "\"" + strconv.Itoa(raiz.Derecha.Valor.Id_Cliente) + "\"" + " [style=invis]}; \n"
		} else if raiz.Izquierda != nil && raiz.Derecha == nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];\n"
			cadena += "\""
			cadena += strconv.Itoa(raiz.Valor.Id_Cliente)
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Izquierda, numero)
			cadena += "\""
			cadena += strconv.Itoa(raiz.Valor.Id_Cliente)
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "{rank=same" + "\"" + strconv.Itoa(raiz.Izquierda.Valor.Id_Cliente) + "\"" + " -> " + "x" + strconv.Itoa(numero) + " [style=invis]}; \n"
		} else if raiz.Izquierda == nil && raiz.Derecha != nil {
			cadena += " x" + strconv.Itoa(numero) + " [label=\"\",width=.1,style=invis];"
			cadena += "\""
			cadena += strconv.Itoa(raiz.Valor.Id_Cliente)
			cadena += "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis]"
			cadena += "; \n \""
			cadena += strconv.Itoa(raiz.Valor.Id_Cliente)
			cadena += "\" -> "
			cadena += a.retornarValoresArbol(raiz.Derecha, numero)
			cadena += "{rank=same" + " x" + strconv.Itoa(numero) + " -> \"" + strconv.Itoa(raiz.Derecha.Valor.Id_Cliente) + "\"" + " [style=invis]}; \n"
		}
	}
	return cadena
}

func (a *Arbol) RecorridoInOrden(raiz *NodoArbol, colaActual *ColaPedidos.ListaCola) {
	if raiz != nil {
		if raiz.Izquierda != nil {
			a.RecorridoInOrden(raiz.Izquierda, colaActual)
		}
		colaActual.Encolar(raiz.Valor.Id_Cliente, raiz.Valor.Nombre_Imagen)
		if raiz.Derecha != nil {
			a.RecorridoInOrden(raiz.Derecha, colaActual)
		}
	}
}
