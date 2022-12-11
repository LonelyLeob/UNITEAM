import axios from "axios";

async function AddSection(id, content) {

    await axios.post(`http://uni-team-inc.online:8000/api/v1/add/section`,
        JSON.stringify({
            course_id: id,
            content: content,
        }),
        {headers:{
                Authorization:`Bearer ${localStorage.getItem('access')}`
            }}
    )
}

export default AddSection
