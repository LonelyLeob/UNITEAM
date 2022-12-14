import axios from "axios";

function CreateAnswer(Answer, Id){

        let res = axios.post(`http://uni-team-inc.online:8080/api/v1/create/answer?field=${Id}`,
            JSON.stringify(
                {
                    answer:Answer
                }),{headers:{
                    Authorization:`Bearer ${localStorage.getItem('access')}`
                }}
        )
}

export default CreateAnswer