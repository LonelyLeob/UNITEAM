import "./signInStyle.css"
import {useState} from "react";
import axios from 'axios'
import { useNavigate  } from "react-router-dom";

function SignIn(){

    const[name, setName] = useState('')
    const[password, setPassword] = useState('')
    const[error, setError] = useState('')

        let navigate = useNavigate()
        const handleClick = () => {
            navigate("/signUp") }


    let handleSubmit = async (e) => {
        e.preventDefault();
        setError('')

        if (name.trim().length === 0 || password.trim().length === 0){
            setError('Введены неверные данные')
            return}

        if (name.trim().length === 0 && password.trim().length === 0){
            setError('Введены неверные данные')
            return}

        try {
                let response = await axios.post("http://uni-team-inc.online:7000/api/v1/authorize",
                    JSON.stringify({
                        name: name,
                        password: password
                    })
                ).then((data) => {
                    setName("")
                    setPassword("")
                    let access = data.data.access
                    let refresh = data.data.refresh
                    localStorage.setItem("access", access)
                    localStorage.setItem("refresh", refresh)
                    navigate("/Form")
                })
        }
        catch {
            setError('Введены неверные данные')
        }
    }

    return(
        <div className="signInContain">
            <div className="signForm">
                <h1 className="formTitle">UNIVERCITY.Inc</h1>
                <input placeholder="Логин" value={name} onChange={(e) => setName(e.target.value)} name="name" type="text"/> <br/>
                <input placeholder="Пароль" value={password} onChange={(e) => setPassword(e.target.value)} name="password" type="password"/><br/>
                {error && <p className="validError">{error}</p>}
                <button onClick={(e) => handleSubmit(e) } type="submit">Войти</button>
                <button type="submit" onClick={handleClick}>Создать аккаунт</button>
            </div>
        </div>
    )
}
export default SignIn