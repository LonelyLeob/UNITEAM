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
            })
        }   catch (err) {
            console.log("u vas err")
        }

    }

    return (
        <div className="signInContain">
            <form method="POST" className="addUserForm" >
                <h1 className="formTitle">UNIVERCITY.Inc</h1>

                <input value={name} onChange={event => setName(event.target.value)} name="name" type="text" placeholder="Логин" required/>
                <input value={password} onChange={event => setPassword(event.target.value)} name="password" type="password" placeholder="Пароль" required/>
                <input value={email} onChange={event => setEmail(event.target.value)} name="email" type="email" placeholder="E-mail" required/>

                <button className="" onClick={(e) => handleSubmit(e) } type="submit">Создать аккаунт</button>
            </form>


        </div>
    )
}

export default SignUp