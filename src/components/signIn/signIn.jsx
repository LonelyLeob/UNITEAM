import "./signInStyle.css"

function SignIn(){
    return(
        <div>
            <form method="POST" className="signForm">

                <h1 className="formTitle"> <b>Sign-In Form</b> </h1>

                <div className="socialNet">

                    <p>Vk</p>
                    <p>Telegram</p>

                </div>

                <label htmlFor="" className="">Name:</label> <br/>
                <input name="name" type="text" required className=""/> <br/>

                <label htmlFor="" className="">Password:</label><br/>
                <input  name="password" type="password" required className=""/><br/>


                <button className="" type="submit">Send</button>

            </form>
        </div>
    )
}
export default SignIn