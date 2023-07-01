import React, {useState} from 'react';
import "../css/login.css"

export const Login = () => {
    const [user, setUser] = useState('')
    const [pass, setPass] = useState('')
    
    
    const handleSubmit = (e) => {
        e.preventDefault();
        console.log(user + " " + pass)
        fetch('http://localhost:3001/login',{
            method: 'POST',
            body: JSON.stringify({
                Username: user,
                Password: pass
            }),
            headers:{
                'Content-Type': 'application/json'
            }
        })
        .then(response => response.json())
        .then(data => validar(data))
    }

    const validar = (data) => {
        if (data.status === 400){
            window.open("/admin", "_self")        
        }else if (data.status === 200){
            localStorage.setItem('empleado', user)
            console.log("Bienvedo usario")
            window.open("/empleado", "_self")
        }else if (data.status === 404){
            alert("Usuario o contraseña incorrecta")
            window.open("/", "_self")
        }
    }
    return (
        <div className="login-box">
            <h2>Login</h2>
            <div className="text-center">
                <form onSubmit={handleSubmit} className="login">
                    <div className='user-box'>
                        <input type="text" id="user" required=""
                        onChange={e => setUser(e.target.value)} 
                        value={user}  
                        />
                        <label className="visually-hidden">Usuario</label>
                    </div>

                    <div className='user-box'>
                        <input type="password" id="password" name="" required="" 
                        onChange={e => setPass(e.target.value)}
                        value={pass} 
                        autoFocus/>
                        <label className="visually-hidden">Password</label>
                    </div>
                    <button type="submit">
                        <span></span>
                        <span></span>
                        <span></span>
                        <span></span>
                        Iniciar Sesion
                    </button>
                </form>
            </div>
        </div>
    )
};