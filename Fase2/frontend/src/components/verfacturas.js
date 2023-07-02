import React, {useState, useEffect} from 'react';
import "../css/factura.css"

export const VerFactura = () => {
    const idEmpleado = localStorage.getItem("empleado")
    const [facturas, setFacturas] = useState([])
    const salir = (e) => {
        e.preventDefault();
        console.log("Listo")
        window.open("/empleado","_self")
    }

    useEffect(() => {
        peticion()
    },[])

    const peticion = () => {
        fetch('http://localhost:3001/facturaempleado',{
        })
        .then(response => response.json())
        .then(data => validar(data))
    }

    const validar = (data) =>{
        console.log(data.factura)
        setFacturas(data.factura) 
    }

    return(
        <div>
            <div className="text-center">
                  <form className="card card-body">
                    <div className="factura">
                    <h1 className="h3 mb-3 fw-normal">Facturas Generadas <br/> Empleado {localStorage.getItem("empleado")}</h1>
                    <br/>
                    <div className='table'>
                        <table className="rwd-table">
                            <thead>
                                <tr>
                                    <th scope="col">#</th>
                                    <th scope="col">ID Cliente</th>
                                    <th scope="col">ID Factura</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    facturas.map((element, j) => {
                                        if (element.Id_Cliente != '') {
                                            return <>
                                            <tr key={"fact"+j}>
                                                <th scope="row">{j+1}</th>
                                                <td>{element.Id_Cliente}</td>
                                                <td>{element.Id_Factura}</td>
                                            </tr>
                                        </>
                                        }
                                    })
                                }
                            </tbody>
                        </table>
                    </div>    
                    </div>
                  </form>
                  <div className="mover-factura">
                    <form className="factura-botones">
                        <center><button onClick={salir}>Regresar</button></center>
                    </form>
                </div>
            </div>
        </div>
    );
}