import axios from "axios";

async function GetShort(setForm) {

    await axios.get("http://uni-team-inc.online:8000/api/v1/get/courses/short",
        {headers:{
                Authorization:`Bearer ${localStorage.getItem('access')}`
            }}
    )
        .then(data => {
            setForm(data.data)
        })
}

export default GetShort
