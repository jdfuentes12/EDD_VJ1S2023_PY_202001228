import React,{useState,useEffect} from "react";
import "../css/factura.css"
import moment from 'moment'

export const Factura = () => {    
    const fecha = moment().format('DD-MM-yyyy -:: hh:mm:ss');
    const idEmpleado = localStorage.getItem("empleado");
    const [cliente, setCliente] = useState('');
    const [pago, setPago] = useState('');

    useEffect(() => {
        peticion();
    },[])

    const peticion = () => {
        fetch('http://localhost:3001/obtenerPedido', {})
        .then(response => response.json())
        .then(data => validar(data))
    }

    const generar = (e) => {
        e.preventDefault();
        console.log(fecha, " ",idEmpleado," ", cliente," ", pago);
        fetch('http://localhost:3001/generarFactura', {
            method: 'POST',
            body: JSON.stringify({
                Timestamp: fecha,
                Biller:    idEmpleado,
                Customer:  cliente,
                Payment:   pago
            }),
            headers:{
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => console.log(data))
    }
    

    const validar = (data) => {
        var idcliente = data.datos.IdCliente
        setCliente(idcliente.toString());
    }

    const salir = (e) => {
        e.preventDefault();
        window.open("/empleado", "_self")
    }

    return(
        <div>
            <div className="text-center">
                <form className="card card-body">
                    <div className="factura">
                        <h1 className="h3 mb-3 fw-normal">Factura de {localStorage.getItem("empleado")}</h1>
                        <br/>
                        <div className="contenido-factura">
                            <label>Fecha</label> <p></p>
                            <input value={fecha} disabled className="input" name="text" type="text"/>
                            <br/><br/>
                            <label>Empleado Cobrado</label> <p></p>
                            <input value={idEmpleado} disabled className="input" name="text" type="text"/>
                            <br/><br/>
                            <label>Usuario</label> <p></p>
                            <input value={cliente} disabled className="input" name="text" type="text"/>
                            <br/><br/>
                            <label>Pago</label> <p></p>
                            <input placeholder="Q 0.0" className="input" name="text" type="text"
                            onChange={(e) => setPago(e.target.value)}
                            value={pago}
                            />
                        </div>
                    </div>
                </form>
                <div className="mover-factura">
                    <form className="factura-botones">
                        <center><button onClick={generar}>Realizar pago</button></center>
                        <br/> 
                        <center><button onClick={salir}>Regresar</button></center>
                    </form>
                </div>
            </div>
        </div>
    );
}