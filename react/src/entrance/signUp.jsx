import './entranceStyle.css'
import {useState} from "react";
import Reg from "./requests/entrainceReg";
import {useNavigate} from "react-router-dom";

function SignUp(){

    const[name,setName] = useState('')
    const[password,setPassword] = useState('')
    const[email,setEmail] = useState('')
    const url = "http://uni-team-inc.online:4000/api/v1/registration"
    const navigate = useNavigate()

    let handleSubmit = async(e) => {
        e.preventDefault()
        await Reg(url, name, password, email)
        navigate("/forms")
    }

    return(
        <div className="entrainceContain">
            <div className="entrainceForm">
                <div className="btnBack">
                    <p className="backAuth" onClick={() => navigate("/")}>⇐</p>
                    <h1 className="formTitle">UNIVERCITY.Inc</h1>
                </div>
                <form>
                    <input value={name} onChange={event => setName(event.target.value)} name="name" type="text" placeholder="Логин" autoComplete="on"/>
                    <input value={password} onChange={event => setPassword(event.target.value)} name="password" type="password" placeholder="Пароль" autoComplete="on"/>
                    <input value={email} onChange={event => setEmail(event.target.value)} name="email" type="email" placeholder="E-mail" autoComplete="on"/>
                    <button className="" onClick={(e) => handleSubmit(e)} type="submit">Создать аккаунт</button>
                </form>
                </div>
            </div>
    )
}

export default SignUp