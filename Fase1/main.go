package main

import (
	"Tda/estructuras"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var listaEmpleados = &estructuras.ListaEmpleados{Inicio: nil, Longitud: 0}
var listaClientes = &estructuras.ListaClientes{Primero: nil, Ultimo: nil, Longitud: 0}
var ListaImagenes = &estructuras.ListaImagenes{Primero: nil, Ultimo: nil}
var colaClientes = &estructuras.ListaColaClientes{Inicio: nil, Final: nil, Longitud: 0}

func main() {
	salir := false

	for !salir {
		fmt.Println("----------Bienvenido----------")
		fmt.Println("1. Inisiar sesion")
		fmt.Println("2. Salir")
		fmt.Println("Elige una opcion: ")
		opcion := 0
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Println()
			sesion()
		case 2:
			fmt.Println("Saliendo de la aplicacion")
			salir = true
		}
	}
}

func sesion() {
	fmt.Println("----------Iniciaar Sesión----------")
	fmt.Println("Ingrese su usuario: ")
	usuario := ""
	fmt.Scanln(&usuario)
	fmt.Println("Ingrese su contraseña: ")
	password := ""
	fmt.Scanln(&password)

	if usuario == "admin_202001228" && password == "admin" {
		fmt.Println("")
		menuAdmin()
	} else {
		fmt.Println("Usuario o contraseña incorrectos")
		main()
	}
}

func menuAdmin() {
	fmt.Println("----------Menú Administrador----------")
	fmt.Println("1. Cargar Empleados")
	fmt.Println("2. Cargar Imagen")
	fmt.Println("3. Cargar Clientes")
	fmt.Println("4. Actualizar Cola")
	fmt.Println("5. Reportes Estructuras")
	fmt.Println("6. Atender cliente en la Cola")
	fmt.Println("7. Salir")
	fmt.Println("Elige una opcion: ")
	opcion := 0
	fmt.Scanln(&opcion)
	switch opcion {
	case 1:
		cargaEmpleados()
	case 2:
		cargarImagen()
	case 3:
		cargarClientes()
	case 4:
		actualizarCola()
	case 5:
		Reportes()
	case 6:
		menuColaClientes()
	case 7:
		fmt.Println("Saliendo del modo administrador")
		main()

	}
}

func cargaEmpleados() {
	f, err := os.Open("archivos\\empleados.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	if _, err := r.Read(); err != nil {
		log.Fatal(err)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		id := record[0]
		nombre := record[1]
		cargo := record[2]
		password := record[3]
		listaEmpleados.Insertar(id, nombre, cargo, password)
	}
	fmt.Println("Lista de Empleados Ingresados")
	//listaEmpleados.MostrarLista()
	fmt.Println()
	menuAdmin()
}

func cargarClientes() {
	fmt.Println("----------Cargar Clientes----------")

	f, err := os.Open("archivos\\clientes_registrados.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	if _, err := r.Read(); err != nil {
		log.Fatal(err)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		id := record[0]
		nombre := record[1]
		listaClientes.Insertar(id, nombre)
	}
	fmt.Println("Lista de Clientes Ingresados\n")
	//listaClientes.MostrarLista()
	menuAdmin()
}

func cargarImagen() {
	fmt.Println("----------Cargar Imagen----------")
	f, err := os.Open("archivos\\imagenes.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	if _, err := r.Read(); err != nil {
		log.Fatal(err)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		imagen := record[0]
		capa := record[1]
		ListaImagenes.Insertar(imagen, capa)
	}

	fmt.Println("Caga de Imagenes Exitosa\n")
	menuAdmin()
}

func actualizarCola() {
	f, err := os.Open("archivos\\clientes_cola.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	if _, err := r.Read(); err != nil {
		log.Fatal(err)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		id := record[0]
		nombre := record[1]
		colaClientes.Encolar(id, nombre)
	}
	fmt.Println("Clientes Ingresados a la cola\n")
	menuColaClientes()
}

func menuColaClientes() {
	id := colaClientes.RetornarID()
	nombre := colaClientes.RetornarNombre()
	fmt.Println("----------Cola de Clientes----------")
	fmt.Println("1. Atender al cliente con id: ", id, " y nombre: ", nombre)
	fmt.Println("2. Salir")
	fmt.Println("Elige una opcion: ")
	opcion := 0
	fmt.Scanln(&opcion)
	switch opcion {
	case 1:
		fmt.Println("Cliente Atendido\n")
		if id == "X" {
			idNuevo := generateRandomID(4)
			listaClientes.Insertar(idNuevo, nombre)
		}
		colaClientes.Desencolar()
		menuColaClientes()
	case 2:
		fmt.Println("Saliendo de la cola de clientes\n")
		menuAdmin()
	}
}

func Reportes() {
	texto := listaEmpleados.GenererarGraphvizEmpleado()
	guardarGraphviz("ListaEmpleados", texto)
	texto1 := listaClientes.GenererarGraphvizCliente()
	guardarGraphviz("ListaClientes", texto1)
	texto2 := ListaImagenes.GenererarGraphvizImagenes()
	guardarGraphviz("ListaImagenes", texto2)
	texto3 := colaClientes.GenererarGraphvizCola()
	guardarGraphviz("ColaClientes", texto3)
	menuAdmin()
}

func guardarGraphviz(nombre string, texto string) {
	archivo, err := os.Create("graficas\\" + nombre + ".dot")
	if err != nil {
		fmt.Println("Error al crear el archivo.\n", err)
		return
	}
	defer archivo.Close()
	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}
	ejecutarDot("graficas\\"+nombre+".jpg", "graficas\\"+nombre+".dot")
	fmt.Println("Archivo guardado correctamente.\n")
}

func ejecutarDot(nombre string, archivo_dot string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tjpg", archivo_dot).Output()
	mode := 0777
	_ = ioutil.WriteFile(nombre, cmd, os.FileMode(mode))
}

func generateRandomID(length int) string {
	rand.Seed(time.Now().UnixNano())

	// Caracteres válidos para el ID
	charset := "0123456789"

	id := make([]byte, length)
	for i := range id {
		id[i] = charset[rand.Intn(len(charset))]
	}

	valor := listaClientes.BuscarClienteId(string(id))
	if valor == true {
		generateRandomID(4)
	}
	return string(id)
}
