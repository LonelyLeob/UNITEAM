import axios from "axios";

async function Restore(url, name, pass){

    await axios.post(url,
        JSON.stringify({
            name: name,
            new: pass
        })
    )
}

export default Restore

