import './entranceStyle.css'
import {useState} from "react";
import Auth from "./requests/entrainceAuth";
import {useNavigate} from "react-router-dom";

function SignIn(){

    const[name, setName] = useState('')
    const[password, setPassword] = useState('')
    const url = "http://uni-team-inc.online:4000/api/v1/authorize"
    const navigate = useNavigate()

        let handleSubmit = async (e) => {
        e.preventDefault()
            await Auth(url, name, password)
            navigate("forms")
    }

    return(
        <div className="entrainceContain">
            <div className="entrainceForm">
                <h1 className="formTitle">UNIVERCITY.Inc</h1>
                <form>
                    <input placeholder="Логин" value={name} onChange={(e) => setName(e.target.value)} name="name" type="text" autoComplete="on"/>
                    <input placeholder="Пароль" value={password} onChange={(e) => setPassword(e.target.value)} name="password" type="password" autoComplete="on"/>
                    <p className="forgotPass" onClick={() => {navigate("/restorePass")}}>Забыли пароль?</p>
                    <button onClick={(e) => handleSubmit(e)} type="submit">Войти</button>
                    <button type="button" onClick={() => navigate("signUp")}>Создать аккаунт</button>
                </form>
            </div>
        </div>
    )
}
export default SignIn