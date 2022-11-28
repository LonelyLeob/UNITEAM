import "./signInStyle.css"
import {useState} from "react";
import axios from 'axios'
import { useNavigate  } from "react-router-dom";

function SignIn(){
    const[name, setName] = useState('')
    const[password, setPassword] = useState('')

        let navigate = useNavigate()
        const handleClick = () => {
            navigate("/signUp") }


    let handleSubmit = async (e) => {
            e.preventDefault();
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
        }   catch (err) {
            console.log("u vas err")
                alert("Введены некорректные данные")
        }
    }



    return(
        <div className="signInContain">
            <div className="signForm">
                <h1 className="formTitle">UNIVERCITY.Inc</h1>
                <input placeholder="Логин" value={name} onChange={(e) => setName(e.target.value)} name="name" type="text"/> <br/>
                <input placeholder="Пароль" value={password} onChange={(e) => setPassword(e.target.value)} name="password" type="password"/><br/>
                <button className="" onClick={(e) => handleSubmit(e) } type="submit">Войти</button>
                {/*<div className="othersSign">*/}
                {/*    <p>Войти с помощью:</p>*/}
                {/*    <div className="socialLink">*/}
                {/*        <a href="" className="vkLink">*/}
                {/*            <span className="icon"></span>*/}
                {/*        </a>*/}
                {/*        <a href="" className="fbLink">*/}
                {/*            <span className="icon"></span>*/}
                {/*        </a>*/}
                {/*        <a href="" className="googleLink">*/}
                {/*            <span className="icon"></span>*/}
                {/*        </a>*/}
                {/*        <a href="" className="qrLink">*/}
                {/*            <span className="icon"></span>*/}
                {/*        </a>*/}
                {/*    </div>*/}
                {/*</div>*/}
                {/*<Link to={}></Link>*/}
                <button className="" type="submit" onClick={handleClick}>Создать аккаунт</button>
            </div>
        </div>
    )
}
export default SignIn