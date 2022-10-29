// import addUsr from "./addUsr"
import './signUpStyle.css'
import {useState} from "react";



function SignUp(){

    const[name,setName] = useState('')
    const[password,setPassword] = useState('')
    const[email,setEmail] = useState('')




    return(
        <div>
            <form method="POST" className="addUserForm" >
                <h1 className="formTitle"> <b>Sign-Up Form</b> </h1>

                <div className="socialNet">

                    <p>Vk</p>
                    <p>Telegram</p>

                </div>

                <label htmlFor="" className="">Name:</label> <br/>
                <input value={name} onChange={event =>  setName(event.target.value)} name="name" type="text" required className=""/> <br/>

                <label htmlFor="" className="">Password:</label><br/>
                <input value={password} onChange={event =>  setPassword(event.target.value)} name="password" type="password" required className=""/><br/>

                <label htmlFor="" className="">Email:</label><br/>
                <input value={email} onChange={event =>  setEmail(event.target.value)} name="email" type="email" required className=""/><br/>

                <button className=""  type="submit">Send</button>
                {/*<button className="" onClick={ () => addUsr(name, password, email)} type="submit">Send</button>*/}
            </form>


        </div>
    )
}

export default SignUp