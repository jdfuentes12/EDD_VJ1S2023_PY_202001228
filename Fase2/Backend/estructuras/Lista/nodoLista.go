package Lista

type Nodo_Empleado struct {
	id        string
	nombre    string
	cargo     string
	password  string
	siguiente *Nodo_Empleado
}
