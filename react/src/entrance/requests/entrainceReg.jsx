import axios from "axios"

async function Reg(url, userName, pass, email){

    await axios.post(url,
        JSON.stringify({
            name: userName,
            password: pass,
            email: email,
        })
    ).then((data) => {
        localStorage.setItem("access", data.data.access)
        localStorage.setItem("refresh", data.data.refresh)
    })
}

export default Reg
