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
	"Fase2/estructuras/ColaPedidos"
	"Fase2/estructuras/Facturas"
	"Fase2/estructuras/Lista"
	"Fase2/estructuras/Matriz"
	"Fase2/estructuras/Peticiones"
	"Fase2/estructuras/TablaHash"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var ListaEmpleado *Lista.ListaEmpleados
var ArbolPedidos *ArbolAVL.Arbol
var MatrizOriginal *Matriz.Matriz
var MatrizFiltro *Matriz.Matriz
var Cola *ColaPedidos.ListaCola
var FacturasRealizadas *Facturas.BlockChain
var VerFacturasRealizadas *TablaHash.TablaHash
var FiltrosColocados string
var EmpleadoLogeado string

func main() {
	//Inicializacion de estructuras
	ListaEmpleado = &Lista.ListaEmpleados{Inicio: nil, Longitud: 0}
	ArbolPedidos = &ArbolAVL.Arbol{Raiz: nil}
	MatrizOriginal = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
	Cola = &ColaPedidos.ListaCola{Inicio: nil, Longitud: 0}
	FacturasRealizadas = &Facturas.BlockChain{Bloques_Creados: 0}
	VerFacturasRealizadas = &TablaHash.TablaHash{Capacidad: 5, Utilizacion: 0}
	FiltrosColocados = ""
	EmpleadoLogeado = ""

	app := fiber.New()
	app.Use(cors.New())

	app.Post("/login", func(c *fiber.Ctx) error {
		fmt.Println("Login")
		var usuario Peticiones.Login
		c.BodyParser(&usuario)
		fmt.Println(usuario)
		if usuario.Username == "ADMIN_202001228" && usuario.Password == "admin" {
			return c.JSON(&fiber.Map{
				"status": 400,
			})
		} else {
			if ListaEmpleado.Inicio != nil {
				if ListaEmpleado.Buscar(usuario.Username, usuario.Password) {
					VerFacturasRealizadas = &TablaHash.TablaHash{Capacidad: 5, Utilizacion: 0}
					VerFacturasRealizadas.NewTablaHash()
					EmpleadoLogeado = usuario.Username
					return c.JSON(&fiber.Map{
						"status": 200,
					})
				}
			}
		}
		return c.JSON(&fiber.Map{
			"status": 404,
		})
	})

	app.Post("/cargarempleados", func(c *fiber.Ctx) error {
		var nombreArchivo Peticiones.Archivo
		c.BodyParser(&nombreArchivo)
		fmt.Println(nombreArchivo)
		cargaEmpleados("archivos/" + nombreArchivo.Nombre)
		return c.JSON(&fiber.Map{
			"status": 201,
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
		ArbolPedidos.RecorridoInOrden(ArbolPedidos.Raiz, Cola)
		return c.JSON(&fiber.Map{
			"status": 200,
			"cola":   Cola,
		})
	})

	app.Get("/reporteavl", func(c *fiber.Ctx) error {
		var imagen Peticiones.RespuestaImagen = Peticiones.RespuestaImagen{Nombre: "Reportes/arbolAVL.jpg"}

		imageByte, err := ioutil.ReadFile(imagen.Nombre)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status": 404,
			})
		}
		//codificar a base64
		imagen.ImagenBase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageByte)
		return c.JSON(&fiber.Map{
			"status": 200,
			"imagen": imagen,
		})
	})

	app.Post("/aplicarfiltro", func(c *fiber.Ctx) error {
		var tipo Peticiones.PeticionFiltro
		MatrizFiltro = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
		c.BodyParser(&tipo)
		fmt.Println(tipo)
		tipo.Nombre_Imagen = Cola.Inicio.Pedido.NombreImagen
		if tipo.Tipo == 1 {
			MatrizFiltro.LeerInicial("csv/"+tipo.Nombre_Imagen+"/inicial.csv", tipo.Nombre_Imagen)
			MatrizFiltro.Negativo(tipo.Nombre_Imagen + "Negativo")
			return c.JSON(&fiber.Map{
				"status": 200,
			})
		} else if tipo.Tipo == 2 {
			MatrizFiltro.LeerInicial("csv/"+tipo.Nombre_Imagen+"/inicial.csv", tipo.Nombre_Imagen)
			MatrizFiltro.EscalaGrises(tipo.Nombre_Imagen + "Gris")
			return c.JSON(&fiber.Map{
				"status": 200,
			})
		} else if tipo.Tipo == 3 {
			MatrizFiltro.LeerInicial("csv/"+tipo.Nombre_Imagen+"/inicial.csv", tipo.Nombre_Imagen)
			MatrizFiltro.RotacionX()
			MatrizFiltro.GenerarImagen(tipo.Nombre_Imagen + "RX")
			return c.JSON(&fiber.Map{
				"status": 200,
			})
		} else if tipo.Tipo == 4 {
			MatrizFiltro.LeerInicial("csv/"+tipo.Nombre_Imagen+"/inicial.csv", tipo.Nombre_Imagen)
			MatrizFiltro.RotacionY()
			MatrizFiltro.GenerarImagen(tipo.Nombre_Imagen + "RY")
			return c.JSON(&fiber.Map{
				"status": 200,
			})
		} else if tipo.Tipo == 5 {
			MatrizFiltro.LeerInicial("csv/"+tipo.Nombre_Imagen+"/inicial.csv", tipo.Nombre_Imagen)
			MatrizFiltro.RotacionDoble()
			MatrizFiltro.GenerarImagen(tipo.Nombre_Imagen + "RD")
			return c.JSON(&fiber.Map{
				"status": 200,
			})
		}
		return c.JSON(&fiber.Map{
			"status": 404,
		})
	})

	app.Get("/obtenerPedido", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"datos": Cola.Inicio.Pedido,
		})

	})

	app.Post("/generarFactura", func(c *fiber.Ctx) error {
		var nuevoBloque Peticiones.BloquePeticion
		c.BodyParser(&nuevoBloque)
		FacturasRealizadas.InsertarBloque(nuevoBloque.Timestamp, nuevoBloque.Biller, nuevoBloque.Customer, nuevoBloque.Payment)
		VerFacturasRealizadas.NewTablaHash()
		FacturasRealizadas.InsertarTabla(VerFacturasRealizadas, EmpleadoLogeado)
		MatrizOriginal = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
		MatrizFiltro = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
		Cola.Desencolar()
		return c.JSON(&fiber.Map{
			"status": 200,
		})
	})

	app.Get("/facturaempleado", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"status":  200,
			"factura": VerFacturasRealizadas.Tabla,
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
