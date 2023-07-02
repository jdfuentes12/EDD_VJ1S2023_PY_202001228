import React,{useState,useEffect} from "react";
import "../css/reporte.css"

export const Reportes = () => {
    const [imagen, setImagen] = useState("https://camo.githubusercontent.com/62da68eb62b1e5f175f7d1f0191dd89a653d7908feb22d37d4a0ab07365d6791/68747470733a2f2f6d656469612e67697068792e636f6d2f6d656469612f4d3967624264396e6244724f5475314d71782f67697068792e676966")

    const reporteArbol = async (e) => {
        e.preventDefault();
        fetch('http://localhost:3001/reporteavl', {})
        .then(response => response.json())
        .then(data => validar(data))
    }

    const validar = (data) => {
        console.log(data)
        if (data.status === 404){
            alert("El arbol no contiene ningun dato");
        }else if (data.status === 200){
            setImagen(data.imagen.ImagenBase64)
        }
        
    }

    const salir = (e) => {
        e.preventDefault();
        console.log("Salir");
        window.open("/admin", "_self")
    }

    return(
        <div>
            <div className="content">
                <h1>Reportes</h1>
                <div className="botones">
                    <button>Grafo</button>
                    <br></br><br></br><br></br>
                    <button onClick={reporteArbol}>Arbol AVL</button>
                    <br></br><br></br><br></br>
                    <button>Facturas</button>
                    <br></br><br></br><br></br>
                    <button onClick={salir}>Regresar</button>
                </div> 
            </div>
            <div className="imagen">
                <br></br>
                <h1>Imagen</h1>
                <br></br>
                <center><img src={imagen} width="350" height="400" alt="some value"></img></center>
            </div>
            
        </div>
    );
}