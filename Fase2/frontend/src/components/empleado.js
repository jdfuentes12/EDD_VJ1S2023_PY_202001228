import React,{useState,useEffect} from "react";
import "../css/empleado.css"

export const Empleado = () => {
    const salir = (e) => {
        e.preventDefault();
        localStorage.removeItem('empleado')
        window.open("/", "_self")
    }

    return(
        <div>
            <div>
                <nav>
                    <label  className="logo">Dashboard Empleado</label>
                    <label  className="empleado">{localStorage.getItem("empleado")}</label>
                </nav>
            </div>
            <div className="botones">
                <form>
                    <center><button>Aplicacion de Filtros</button></center>
                    <br/>
                    <center><button>Generar Facturas</button></center>
                    <br/>
                    <center><button>Ver Facturas</button></center>
                    <br/>
                    <center><button onClick={salir}>Cerrar Sesion</button></center>
                </form>
            </div>
        </div>
    );
}