import "./signInStyle.css"
import {useState} from "react";
import axios from 'axios'

function SignIn(){
    const[name,setName] = useState('')
    const[password,setPassword] = useState('')

    let handleSubmit = async (e) => {
            e.preventDefault();
            try {
            let response = await axios.post("http://localhost:7000/auth/token", 
                JSON.stringify({
                name: name,
                password: password
            }), 
            {withCredentials: true}
            )
            let res = await axios.post("http://localhost:8080/forms/create", 
            JSON.stringify({
            name: "123",
            desc: "desc",
            anon: true
        }), 
        {withCredentials: true}
        ).then(data => console.log(data))
        }   catch (err) {
            console.log("u vas err")
        }
    }
      
    

    return(
        <div>
            <form className="signForm">

                <h1 className="formTitle"> <b>Sign-In Form</b> </h1>

                <div className="socialNet">

                    <p>Vk</p>
                    <p>Telegram</p>

                </div>
                
                <label>Name:</label> 
                <br/>
                <input value={name} onChange={(e) => setName(e.target.value)} name="name" type="text" required className=""/> <br/>

                <label>Password:</label>
                <br/>
                <input value={password} onChange={(e) => setPassword(e.target.value)} name="password" type="password" required className=""/><br/>

                <button className="" onClick={(e) => handleSubmit(e) } type="submit">Send</button>

            </form>


        </div>
    )
}
export default SignIn