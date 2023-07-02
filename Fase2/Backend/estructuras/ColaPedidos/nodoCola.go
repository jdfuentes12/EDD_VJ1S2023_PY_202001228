package ColaPedidos

type NodoCola struct {
	Pedido    *PedidoCola
	Siguiente *NodoCola
}

type PedidoCola struct {
	IdCliente    int
	NombreImagen string
}
