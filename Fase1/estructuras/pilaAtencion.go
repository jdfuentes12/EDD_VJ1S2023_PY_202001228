package estructuras

import (
	"strconv"
)

type PilaVitacora struct {
	primero  *Nodo_Pila
	longitud int
}

func (p *PilaVitacora) Push(id string, imagen string) {
	nuevo := &Nodo_Pila{id: id, imagen: imagen, siguiente: nil}
	if p.primero == nil {
		p.primero = nuevo
	} else {
		nuevo.siguiente = p.primero
		p.primero = nuevo
	}
	p.longitud++
}

func (p *PilaVitacora) MostrarPila() {
	aux := p.primero
	for aux != nil {
		println(aux.id, aux.imagen)
		aux = aux.siguiente
	}
}

func (p *PilaVitacora) GenererarGraphvizPila() string {
	texto := "digraph PilaVitacora {\n"
	contador := 0
	aux := p.primero
	if aux != nil {
		for aux != nil {
			texto += "nodo" + strconv.Itoa(contador) + "[label = \"" + aux.id + ", " + aux.imagen + "\"];\n"
			contador++
			aux = aux.siguiente
		}
		texto += "empty[label = \"Pila\"];\n"
		texto += "empty->nodo0;\n"
		for i := 0; i < contador-1; i++ {
			texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(i+1) + ";\n"
		}
	} else {
		texto += "empty[label = \"PilaVacia\"];\n"
	}
	texto += "}"

	return texto
}
