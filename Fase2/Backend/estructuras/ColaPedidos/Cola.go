package ColaPedidos

type ListaCola struct {
	Inicio   *NodoCola
	Longitud int
}

func (c *ListaCola) Encolar(id_cliente int, imagen string) {
	nuevoNodo := &PedidoCola{IdCliente: id_cliente, NombreImagen: imagen}
	if c.Longitud == 0 {
		nuevoNodo := &NodoCola{nuevoNodo, nil}
		c.Inicio = nuevoNodo
		c.Longitud++
	} else {
		nuevoNodo := &NodoCola{nuevoNodo, nil}
		aux := c.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = nuevoNodo
		c.Longitud++
	}
}

func (c *ListaCola) Desencolar() *PedidoCola {
	if c.Longitud == 0 {
		return nil
	} else {
		aux := c.Inicio
		c.Inicio = aux.Siguiente
		c.Longitud--
		return aux.Pedido
	}
}
