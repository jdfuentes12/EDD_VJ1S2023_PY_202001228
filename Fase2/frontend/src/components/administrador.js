import React,{useState,useEffect} from "react";
import "../css/admin.css"

export const Administrador = () => {

    const onChange = (e) => {
        e.preventDefault();
        var reader = new FileReader();
        reader.onload = (e) => {
            var obj = JSON.parse(e.target.result);
            console.log(obj.pedidos);
            fetch('http://localhost:3001/cargarpedidos', {
                method: 'POST',
                body: JSON.stringify({
                    Pedidos: obj.pedidos
                }),
                headers:{
                    'Content-Type': 'application/json'
                }
            })
            .then(response => response.json())
            .then(data => validar(data))
        };
        reader.readAsText(e.target.files[0]);
    }

    const onChange1 = (e,file1) => {
        var file = file1 || e.target.files[0];
        console.log(file);
        console.log(file.name);
        fetch('http://localhost:3001/cargarempleados',{
            method: 'POST',
            body: JSON.stringify({
                Nombre: file.name
            }),
            headers:{
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => validar(data))
    }
    
    const validar = (data) => {
        console.log(data);
        console.log(data.status);
        if (data.status === 200){
            alert("Carga de pedidos exitosamente");
        } else if (data.status === 201){
            alert("Carga de empleados exitosamente");
        } else if (data.status === 404){
            alert("Error al generar la imagen");
        }
    }

    const reportes = (e) => {
        e.preventDefault();
        console.log("Reportes");
        window.open("/reportes", "_self")
    }

    const salir = (e) => {
        e.preventDefault();
        console.log("Salir");
        window.open("/", "_self")
    }

    return (
        <div className="card">
            <div className="tools">
                <div className="circle">
                    <span className="red box"></span>
                </div>
                <div className="circle">
                    <span className="yellow box"></span>
                </div>
                <div className="circle">
                    <span className="green box"></span>
                </div>
            </div>
            <div className="card__content">
                <h1>Bienvenido Administrador</h1>
                <div className="pedidos">
                    <h2>Cargar pedidos</h2>
                    <div className="container-input">
                        <input type="file" name="file-1" id="file-1" className="inputfile inputfile-1" data-multiple-caption="{count} archivos seleccionados" multiple 
                        accept="application/json"
                        onChange={onChange}
                        />
                        <label htmlFor="file-1">
                        <svg xmlns="http://www.w3.org/2000/svg" className="iborrainputfile" width="20" height="17" viewBox="0 0 20 17"><path d="M10 0l-5.2 4.9h3.3v5.1h3.8v-5.1h3.3l-5.2-4.9zm9.3 11.5l-3.2-2.1h-2l3.4 2.6h-3.5c-.1 0-.2.1-.2.1l-.8 2.3h-6l-.8-2.2c-.1-.1-.1-.2-.2-.2h-3.6l3.4-2.6h-2l-3.2 2.1c-.4.3-.7 1-.6 1.5l.6 3.1c.1.5.7.9 1.2.9h16.3c.6 0 1.1-.4 1.3-.9l.6-3.1c.1-.5-.2-1.2-.7-1.5z"></path></svg>
                        <span className="iborrainputfile">Seleccionar archivo</span>
                        </label>
                    </div>
                </div>
                <div className="empleados">
                    <h2>Cargar empleados</h2>
                    <div className="container-input">
                        <input type="file" name="file-2" id="file-2" className="inputfile inputfile-2" data-multiple-caption="{count} archivos seleccionados" multiple 
                        accept="application/csv"
                        onChange={onChange1}
                        />
                        <label htmlFor="file-2">
                        <svg xmlns="http://www.w3.org/2000/svg" className="iborrainputfile" width="20" height="17" viewBox="0 0 20 17"><path d="M10 0l-5.2 4.9h3.3v5.1h3.8v-5.1h3.3l-5.2-4.9zm9.3 11.5l-3.2-2.1h-2l3.4 2.6h-3.5c-.1 0-.2.1-.2.1l-.8 2.3h-6l-.8-2.2c-.1-.1-.1-.2-.2-.2h-3.6l3.4-2.6h-2l-3.2 2.1c-.4.3-.7 1-.6 1.5l.6 3.1c.1.5.7.9 1.2.9h16.3c.6 0 1.1-.4 1.3-.9l.6-3.1c.1-.5-.2-1.2-.7-1.5z"></path></svg>
                        <span className="iborrainputfile">Seleccionar archivo</span>
                        </label>
                    </div>
                </div>
            </div>
            <button className="buttonReporte" onClick={reportes}>Reportes</button>
            <button className="button" onClick={salir}>Cerrar Sesion</button>
        </div>
    )
}