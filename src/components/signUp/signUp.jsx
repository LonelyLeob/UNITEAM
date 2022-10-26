// import addUsr from "./addUsr"
import './signUpStyle.css'

function SignUp(){
    return(
        <div>
            <form method="POST" className="addUserForm" >
                <h1 className="formTitle"> <b>Sign-Up Form</b> </h1>

                <div className="socialNet">

                    <p>Vk</p>
                    <p>Telegram</p>

                </div>

                <label htmlFor="" className="">Name:</label> <br/>
                <input name="name" type="text" required className=""/> <br/>

                <label htmlFor="" className="">Password:</label><br/>
                <input  name="password" type="password" required className=""/><br/>

                <label htmlFor="" className="">Email:</label><br/>
                <input name="email" type="email" required className=""/><br/>

                <button className="" type="submit">Send</button>
            </form>
        </div>
    )
}

export default SignUp