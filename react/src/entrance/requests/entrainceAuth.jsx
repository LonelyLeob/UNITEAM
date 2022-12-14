import axios from "axios";

async function Auth(url, userName, pass){
     await axios.post(url,
            JSON.stringify({
                name: userName,
                password: pass
            })
        )
            .then((data) => {
            localStorage.setItem("access", data.data.access)
            localStorage.setItem("refresh", data.data.refresh)  
        })
}

export default Auth

