package Facturas

import (
	"Fase2/estructuras/GenerarArchivos"
	"Fase2/estructuras/TablaHash"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type BlockChain struct {
	Inicio          *NodoBloque
	Bloques_Creados int
}

func (b *BlockChain) InsertarBloque(fecha string, biller string, customer string, payment string) {
	cadenaFuncion := strconv.Itoa(b.Bloques_Creados) + fecha + biller + customer + payment
	hash := SHA256(cadenaFuncion)
	if b.Bloques_Creados == 0 {
		datosBloque := map[string]string{
			"index":        strconv.Itoa(b.Bloques_Creados),
			"timestamp":    fecha,
			"biller":       biller,
			"customer":     customer,
			"payment":      payment,
			"previoushash": "0000",
			"hash":         hash,
		}
		nuevoBloque := &NodoBloque{Bloque: datosBloque}
		b.Inicio = nuevoBloque
	} else {
		aux := b.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		datosBloque := map[string]string{
			"index":        strconv.Itoa(b.Bloques_Creados),
			"timestamp":    fecha,
			"biller":       biller,
			"customer":     customer,
			"payment":      payment,
			"previoushash": aux.Bloque["hash"],
			"hash":         hash,
		}
		nuevoBloque := &NodoBloque{Bloque: datosBloque, Anterior: aux}
		aux.Siguiente = nuevoBloque
	}
	b.Bloques_Creados++
}

func SHA256(cadena string) string {
	hexaString := ""
	h := sha256.New()
	h.Write([]byte(cadena))
	hash := h.Sum(nil)
	hexaString = hex.EncodeToString(hash)
	return hexaString
}

func (b *BlockChain) InsertarTabla(tabla *TablaHash.TablaHash, idEmpleado string) {
	aux := b.Inicio
	for aux != nil {
		if aux.Bloque["biller"] == idEmpleado {
			tabla.Insertar(aux.Bloque["customer"], aux.Bloque["hash"])
		}
		aux = aux.Siguiente
	}
}

func (b *BlockChain) ReporteBlockChain() {
	nombre_archivo := "./Reportes/factura.dot"
	nombre_imagen := "./Reportes/factura.jpg"
	contenido := "digraph G{\n"
	contador := 0
	aux := b.Inicio

	if aux != nil {
		for aux != nil {
			contenido += "nodo" + strconv.Itoa(contador) + "[label=\"" + aux.Bloque["index"] + "\"];\n"
			contador++
			aux = aux.Siguiente
		}
		for i := 0; i < contador-1; i++ {
			contenido += "nodo" + strconv.Itoa(i) + "-> nodo" + strconv.Itoa(i+1) + ";\n"
		}
	} else {
		contenido += "nodo0[label=\"Bloque Vacio\"];\n"
	}
	contenido += "}"
	GenerarArchivos.CrearArchivo(nombre_archivo)
	GenerarArchivos.EscribirArchivo(contenido, nombre_archivo)
	GenerarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}
