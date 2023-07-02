import React,{useState,useEffect} from "react";
import "../css/empleado.css"

export const Empleado = () => {
    const salir = (e) => {
        e.preventDefault();
        localStorage.removeItem('empleado')
        window.open("/", "_self")
    }

    const verFatura = (e) => {
        e.preventDefault();
        console.log("Reportes");
        window.open("/verfactura", "_self")
    }

    const filtros = (e) => {
        e.preventDefault();
        window.open("/filtros", "_self")
    }

    const factura = (e) => {
        e.preventDefault();
        window.open("/factura", "_self")
    }

    return(
        <div>
            <div>
                <nav>
                    <label  className="logo">Empleado</label>
                    <label  className="empleado">{localStorage.getItem("empleado")}</label>
                </nav>
            </div>
            <div className="botones">
                <form>
                    <center><button onClick={filtros}>Aplicacion de Filtros</button></center>
                    <br/>
                    <center><button onClick={factura}>Generar Facturas</button></center>
                    <br/>
                    <center><button onClick={verFatura}>Ver Facturas</button></center>
                    <br/>
                    <center><button onClick={salir}>Cerrar Sesion</button></center>
                </form>
            </div>
        </div>
    );
}