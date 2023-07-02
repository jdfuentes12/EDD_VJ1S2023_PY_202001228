package Peticiones

type Login struct {
	Username string
	Password string
}

type Archivo struct {
	Nombre string
}

type Pedido struct {
	Id_Cliente    int    `json:"id_cliente"`
	Nombre_Imagen string `json:"imagen"`
}

type ArbolPeticion struct {
	Pedidos []Pedido
}

type RespuestaImagen struct {
	ImagenBase64 string
	Nombre       string
}

type PeticionFiltro struct {
	Tipo          int
	Nombre_Imagen string
}

type BloquePeticion struct {
	Timestamp string
	Biller    string
	Customer  string
	Payment   string
}
