import React,{useState,useEffect} from "react";
import "../css/filtros.css"

export const Filtros = () => {
    const salir = (e) => {
        e.preventDefault();
        localStorage.removeItem('empleado')
        window.open("/empleado", "_self")
    }
    return(
        <div>
            <div className="text-center">
                <form className="card card-body">
                <h1 className="h3 mb-3 fw-normal">Dashboard Empleado {localStorage.getItem("empleado")}</h1>
                <br/>
                <h3 className="h3 mb-3 fw-normal">Elige un Filtro</h3>
                <br/>
                <div className="selectdiv">
                    <select className="form-selectdiv" aria-label=".form-select-lg example" >
                        <option value={0}>Elegir</option>
                        <option value={1}>Negativo</option>
                        <option value={2}>Escala de Grises</option>
                        <option value={3}>Espejo X</option>
                        <option value={4}>Espejo Y</option>
                        <option value={5}>Ambos Espejos</option>
                    </select>
                </div>
                <br/>
                
                </form>
                <div className="mover-filtros">
                    <form>
                        <center><button >Generar Imagen con Filtro</button></center>
                        <br/> <br/>
                        <center><button onClick={salir}>Cerrar Sesion</button></center>
                    </form>
                </div>
            </div>
        </div>
    );
}