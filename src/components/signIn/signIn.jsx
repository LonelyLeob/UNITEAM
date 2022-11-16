import "./signInStyle.css"
import {useState} from "react";
import axios from 'axios'


function SignIn(){
    const[name, setName] = useState('')
    const[password, setPassword] = useState('')

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
            })
        }   catch (err) {
            console.log("u vas err")
        }
    }



    return(
        <div className="signInContain">
            <form className="signForm">

                <h1 className="formTitle"> <b>Авторизация</b> </h1>

                <div className="socialNet">

                    <p>Vk</p>
                    <p>Telegram</p>

                </div>
                
                <label>Имя пользователя:</label>
                <br/>
                <input value={name} onChange={(e) => setName(e.target.value)} name="name" type="text" required className=""/> <br/>

                <label>Пароль:</label>
                <br/>
                <input value={password} onChange={(e) => setPassword(e.target.value)} name="password" type="password" required className=""/><br/>

                <button className="" onClick={(e) => handleSubmit(e) } type="submit">Войти</button>
            </form>


        </div>
    )
}
export default SignIn