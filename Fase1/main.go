package main

import (
	"Tda/estructuras"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

var listaEmpleados = &estructuras.ListaEmpleados{Inicio: nil, Longitud: 0}

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
	fmt.Println("3. Cargar usuarios")
	fmt.Println("4. Actualizar Cola")
	fmt.Println("5. Reportes Estructuras")
	fmt.Println("6. Salir")
	fmt.Println("Elige una opcion: ")
	opcion := 0
	fmt.Scanln(&opcion)
	switch opcion {
	case 1:
		cargaEmpleados()
	case 6:
		fmt.Println("Saliendo del modo administrador")
		main()

	}
}

func cargaEmpleados() {
	fmt.Println("----------Cargar Empleados----------")
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
	listaEmpleados.MostrarLista()
	menuAdmin()

}

func cargarClientes() {
	fmt.Println("----------Cargar Clientes----------")

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
		if id == "X" {
			idradom := generateRandomID(5)
			id = idradom
		}
		fmt.Println(id)
		nombre := record[1]
		fmt.Println(nombre)

	}
}

func generateRandomID(length int) string {
	rand.Seed(time.Now().UnixNano())

	// Caracteres válidos para el ID
	charset := "0123456789"

	id := make([]byte, length)
	for i := range id {
		id[i] = charset[rand.Intn(len(charset))]
	}
	return string(id)
}
