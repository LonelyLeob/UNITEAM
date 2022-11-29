import axios from 'axios'
import './signUpStyle.css'
import {useState} from "react";
import { useNavigate  } from "react-router-dom";

function SignUp(){

    const[name,setName] = useState('')
    const[password,setPassword] = useState('')
    const[email,setEmail] = useState('')
    const[error, setError] = useState('')

    let navigate = useNavigate()

    let handleSubmit = async (e) => {
            e.preventDefault();
            setError('')

            if (name.trim().length === 0 || password.trim().length === 0 || email.trim().length === 0){
                setError('Введены неверные данные')
                return}

            if (name.trim().length === 0 && password.trim().length === 0 && email.trim().length === 0){
                setError('Введены неверные данные')
                return}

            try {
                    let response = await axios.post("http://uni-team-inc.online:7000/api/v1/registration",
                        JSON.stringify({
                        name: name,
                        password: password,
                        email: email
                    }))
                        .then((data) => {
                        setName("")
                        setPassword("")
                        setEmail("")
                            let access = data.data.access
                            let refresh = data.data.refresh
                            localStorage.setItem("access", access)
                            localStorage.setItem("refresh", refresh)
                            navigate("/Form")
                    })
            }
            catch{
                setError('Введены неверные данные')
            }
    }

    return (
        <div className="signInContain">
            <form method="POST" className="addUserForm" >
                    <div className="btnBack">
                        <p className="backAuth" onClick={() => {  navigate("/")}}>⇐</p>
                        <h1 className="formTitle">UNIVERCITY.Inc</h1>
                    </div>
                <input value={name} onChange={event => setName(event.target.value)} name="name" type="text" placeholder="Логин" required/>
                <input value={password} onChange={event => setPassword(event.target.value)} name="password" type="password" placeholder="Пароль" required/>
                <input value={email} onChange={event => setEmail(event.target.value)} name="email" type="email" placeholder="E-mail" required/>
                {error && <p className="validError">{error}</p>}
                <button className="" onClick={(e) => handleSubmit(e) } type="submit">Создать аккаунт</button>
            </form>


        </div>
    )
}

export default SignUp