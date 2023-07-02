import React,{useState,useEffect} from "react";
import "../css/filtros.css"

export const Filtros = () => {
    const [filtro, setFiltro] = useState(0);

    const handleChange = (e) => {
        var j = parseInt(e.target.value);
        setFiltro(j)
    }
    
    const salir = (e) => {
        e.preventDefault();
        window.open("/empleado", "_self")
    }

    const aplicarFiltro = (e) => {
        e.preventDefault();
        fetch('http://localhost:3001/aplicarfiltro',{
            method: 'POST',
            body: JSON.stringify({
                Tipo: filtro,
                Nombre_Imagen: ""
            }),
            headers:{
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => validar(data))
    }

    const validar = (e) => {
        if (e.status === 200){
            alert("Filtro Aplicado")
        }else if (e.status === 404){
            alert("Error al Aplicar Filtro")
        }
    }
    
    
    return(
        <div>
            <div className="text-center">
                <form className="card card-body">
                <h1 className="h3 mb-3 fw-normal">Empleado {localStorage.getItem("empleado")}</h1>
                <br/>
                <h3 className="h3 mb-3 fw-normal">Elige un Filtro</h3>
                <br/>
                <div className="selectdiv">
                    <select className="form-selectdiv" onChange={handleChange} aria-label=".form-select-lg example" >
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
                        <center><button onClick={aplicarFiltro}>Generar Imagen con Filtro</button></center>
                        <br/> <br/>
                        <center><button onClick={salir}>Regresar</button></center>
                    </form>
                </div>
            </div>
        </div>
    );
}