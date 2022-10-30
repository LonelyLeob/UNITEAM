import "./signInStyle.css"
import {useState} from "react";
import axios from "axios";
// import authUsr from "./authUsr";


function SignIn(){
    const[name,setName] = useState('')
    const[password,setPassword] = useState('')
    const url = 'http://localhost:7000/auth/token'

                                            //          AXIOS FUNC
    async function authUsr(name,password,url){
        const data = {
                       name: name,
                       password: password,
                   }
        await axios.post(url, JSON.stringify(data), {withCredentials: true})
    }


                                                //          FETCH FUNC
 // async function authUsr(name,password,url){
 //        const data = {
 //            name: name,
 //            password: password,
 //        }
 //        let response = await fetch(url,{
 //            method:'POST',
 //            body: JSON.stringify(data),
 //        })
 //    }


    return(
        <div>
            <form method="POST"  className="signForm">

                <h1 className="formTitle"> <b>Sign-In Form</b> </h1>

                <div className="socialNet">

                    <p>Vk</p>
                    <p>Telegram</p>

                </div>

                <label htmlFor="" className="">Name:</label> <br/>
                <input value={name} onChange={event => setName(event.target.value)} name="name" type="text" required className=""/> <br/>

                <label htmlFor="" className="">Password:</label><br/>
                <input value={password} onChange={event => setPassword(event.target.value)} name="password" type="password" required className=""/><br/>

                <button className="" onClick={ () => authUsr(name, password, url) } type="submit">Send</button>

            </form>


        </div>
    )
}
export default SignIn