import axios from "axios";

async function GetCourse(setCourse, id) {

    await axios.get(`http://uni-team-inc.online:8000/api/v1/get/course?course=${id}`,
        {headers:{
                Authorization:`Bearer ${localStorage.getItem('access')}`
            }}
    )
        .then(data => {
            setCourse(data.data)
        })
}

export default GetCourse
