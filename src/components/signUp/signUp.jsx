import axios from 'axios'
import './signUpStyle.css'
import {useState} from "react";



function SignUp(){

    const[name,setName] = useState('')
    const[password,setPassword] = useState('')
    const[email,setEmail] = useState('')

    let handleSubmit = async (e) => {
            e.preventDefault();
            try {
            let response = await axios.post("http://localhost:7000/registration", 
                JSON.stringify({
                name: name,
                password: password,
                email: email
            }), 
            {withCredentials: true}
            ).then(() => {
                setName("")
                setPassword("")
                setEmail("")
            })
        }   catch (err) {
            console.log("u vas err")
        }
    }


    return (
        <div className="signInContain">
            <form method="POST" className="addUserForm" >
                <h1 className="formTitle"> <b>Регистрация</b> </h1>

                <div className="socialNet">

                    <p>Vk</p>
                    <p>Telegram</p>

                </div>

                <label htmlFor="" className="">Имя пользователя:</label>
                <br/>
                <input value={name} onChange={event => setName(event.target.value)} name="name" type="text" required className=""/> 
                <br/>

                <label htmlFor="" className="">Пароль:</label>
                <br/>
                <input value={password} onChange={event => setPassword(event.target.value)} name="password" type="password" required className=""/>
                <br/>

                <label htmlFor="" className="">E-mail:</label><br/>
                <input value={email} onChange={event => setEmail(event.target.value)} name="email" type="text" required className=""/>
                <br/>

                <button className="" onClick={(e) => handleSubmit(e) } type="submit">Создать аккаунт</button>
            </form>


        </div>
    )
}

export default SignUp