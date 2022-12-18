import axios from "axios";

async function UpdateSection(id, content) {

    await axios.patch(`http://uni-team-inc.online:8000/api/v1/update/section?section=${id}`,
        JSON.stringify({
            content: content,
        }),
        {headers:{
                Authorization:`Bearer ${localStorage.getItem('access')}`
            }}
    )
}

export default UpdateSection
