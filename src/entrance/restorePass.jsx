import './entranceStyle.css'
import {useState} from "react";
import Restore from "./requests/restorePass";
import {useNavigate} from "react-router-dom";

function RestorePass(){


    const[password, setPassword] = useState('')
    const[name, setName] = useState('')
    const navigate = useNavigate()
    const url = "http://uni-team-inc.online:4000/api/v1/forget/pwd"

    const handleSubmit = async (e) => {
        e.preventDefault()
        await Restore(url, name, password)
        // navigate("/")
    }

    return(
        <div className="entrainceContain">
            <div className="entrainceForm">
                <div className="btnBack">
                    <p className="backAuth" onClick={() => navigate("/")}>⇐</p>
                    <h1 className="formTitle">UNIVERCITY.Inc</h1>
                </div>
                <form>
                    <input placeholder="Логин" value={name} onChange={(e) => setName(e.target.value)} name="name" type="text" autoComplete="on"/>
                    <input placeholder="Новый пароль" value={password} onChange={(e) => setPassword(e.target.value)} name="password" type="password" autoComplete="on"/>
                    <button onClick={(e) => handleSubmit(e)} type="submit">Отправить</button>
                    <p>Вам на почту будет отправлено письмо, перейдите по ссылке в письме чтобы обновить пароль</p>
                </form>
            </div>
        </div>
    )
}
export default RestorePass