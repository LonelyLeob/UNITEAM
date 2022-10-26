// import addUsr from "./addUsr"

function SignUp(){
    return(
        <div>
            <form method="POST" className=" mx-40  inline-block p-10 border-double border-4 border-cyan-400 max-w-3xl bg-slate-100 text-l" >
                <h1 className="mb-3 text-center">Sign-Up Form.</h1>
                <label htmlFor="" className="mr-2">Name:</label> <br/>
                <input name="name" type="text" required className=" mb-3 bg-slate-400 border-double border-2 rounded-lg w-64"/> <br/>

                <label htmlFor="" className="mr-2">Password:</label><br/>
                <input  name="password" type="password" required className=" mb-3 bg-slate-400 border-double border-2 rounded-lg w-64"/><br/>

                <label htmlFor="" className="mr-2">Email:</label><br/>
                <input name="email" type="email" required className=" mb-3 bg-slate-400 border-double border-2 rounded-lg w-64"/><br/>

                <button className="p-1 w-48 border-double border-4 border-cyan-400 max-w-3xl bg-slate-100 rounded-full text-center" type="submit">Send</button>
            </form>
        </div>
    )
}

export default SignUp