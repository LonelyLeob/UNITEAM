import axios from "axios";

async function AddCourse(title, desc) {

    await axios.post(`http://uni-team-inc.online:8000/api/v1/add/course`,
        JSON.stringify({
            title: title,
            desc: desc,
        }),
        {headers:{
                Authorization:`Bearer ${localStorage.getItem('access')}`
            }}
    )
}

export default AddCourse
