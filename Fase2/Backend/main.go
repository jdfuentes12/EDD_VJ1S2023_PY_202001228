package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"Fase2/estructuras/ArbolAVL"
	"Fase2/estructuras/Lista"
	"Fase2/estructuras/Peticiones"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var ListaEmpleado = &Lista.ListaEmpleados{Inicio: nil, Longitud: 0}
var ArbolPedidos = &ArbolAVL.Arbol{Raiz: nil}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Post("/login", func(c *fiber.Ctx) error {
		fmt.Println("Login")
		var usuario Peticiones.Login
		c.BodyParser(&usuario)
		fmt.Println(usuario)
		if usuario.Username == "admin" && usuario.Password == "admin" {
			return c.JSON(&fiber.Map{
				"status": 400,
			})
		} else {
			validar := ListaEmpleado.Buscar(usuario.Username, usuario.Password)
			if validar {
				return c.JSON(&fiber.Map{
					"status": 200,
				})
			} else {
				return c.JSON(&fiber.Map{
					"status": 404,
				})
			}
		}
	})

	app.Post("/cargarempleados", func(c *fiber.Ctx) error {
		var nombreArchivo Peticiones.Archivo
		c.BodyParser(&nombreArchivo)
		fmt.Println(nombreArchivo)
		cargaEmpleados("archivos/" + nombreArchivo.Nombre)
		return c.JSON(&fiber.Map{
			"datos": ListaEmpleado,
		})
	})

	app.Post("/cargarpedidos", func(c *fiber.Ctx) error {
		var pedidosNuevos Peticiones.ArbolPeticion
		c.BodyParser(&pedidosNuevos)
		fmt.Println(pedidosNuevos)
		for i := 0; i < len(pedidosNuevos.Pedidos); i++ {
			fmt.Println(pedidosNuevos.Pedidos[i].Id_Cliente, pedidosNuevos.Pedidos[i].Nombre_Imagen)
			ArbolPedidos.InsertarElemento(pedidosNuevos.Pedidos[i].Id_Cliente, pedidosNuevos.Pedidos[i].Nombre_Imagen)
		}
		ArbolPedidos.Graficar()
		return c.JSON(&fiber.Map{
			"status": 200,
		})
	})

	app.Get("/reporteavl", func(c *fiber.Ctx) error {
		var imagen Peticiones.RespuestaImagen = Peticiones.RespuestaImagen{Nombre: "Reportes/arbolAVL.jpg"}

		imageByte, err := ioutil.ReadFile(imagen.Nombre)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status": 400,
			})
		}
		//codificar a base64
		imagen.ImagenBase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageByte)
		return c.JSON(&fiber.Map{
			"status": 200,
			"imagen": imagen,
		})
	})

	app.Listen(":3001")
}

func cargaEmpleados(ruta string) {

	f, err := os.Open(ruta)
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

		ListaEmpleado.Insertar(record[0], record[1], record[2], record[3])
	}
	fmt.Println("Lista de Empleados Ingresados")
}
