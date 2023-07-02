import './App.css';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import { Login } from './components/login';
import { Administrador } from './components/administrador';
import { Reportes } from './components/reporte';
import { Empleado } from './components/empleado';
import { Filtros } from './components/filtos';
import { Factura } from './components/factura';
import { VerFactura } from './components/verfacturas';

function App() {
  return (
    <Router>
      <Routes>
        <Route exact path='/' element={<Login/>}/>
        <Route exact path='/admin' element={<Administrador/>}/>
        <Route exact path='/reportes' element={<Reportes/>}/>
        <Route exact path='/empleado' element={<Empleado/>}/>
        <Route exact path='/filtros' element={<Filtros/>}/>
        <Route exact path='/factura' element={<Factura/>} />
        <Route exact path='/verfactura' element={<VerFactura/>} />
      </Routes>
    </Router>
  )
}

export default App;
