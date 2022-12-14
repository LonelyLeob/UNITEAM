 import axios from "axios";

async function AddNewForm(formName, formDescription, formAnon) {

    await axios.post("http://uni-team-inc.online:8080/api/v1/create",
        JSON.stringify({
            name: formName,
            desc: formDescription,
            anon: formAnon
        }),
        {headers:{
                Authorization:`Bearer ${localStorage.getItem('access')}`
            }})
}

export default AddNewForm