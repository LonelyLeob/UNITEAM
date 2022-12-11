import axios from "axios";

async function GetShort(setCourse) {

    await axios.get("http://uni-team-inc.online:8000/api/v1/get/courses/short",
        {headers:{
                Authorization:`Bearer ${localStorage.getItem('access')}`
            }}
    )
        .then(data => {
            setCourse(data.data)
        })
}

export default GetShort
