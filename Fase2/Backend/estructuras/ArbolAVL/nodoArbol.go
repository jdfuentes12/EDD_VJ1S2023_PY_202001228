package ArbolAVL

import "Fase2/estructuras/Peticiones"

type NodoArbol struct {
	Izquierda         *NodoArbol
	Derecha           *NodoArbol
	Valor             *Peticiones.Pedido
	Altura            int
	Factor_Equilibrio int
}
