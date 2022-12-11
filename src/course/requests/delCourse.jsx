import axios from "axios";

async function DelCourse(id) {

    await axios.delete(`http://uni-team-inc.online:8000/api/v1/delete/course?course=${id}`,
        {headers:{
                Authorization:`Bearer ${localStorage.getItem('access')}`
            }}
    )
}

export default DelCourse
