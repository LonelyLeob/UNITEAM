import axios from "axios";
import GetUserForms from "../../forms/requests/getUserForms";

async function CreateField(Uuid, Field){

             await axios.post(`http://uni-team-inc.online:8080/api/v1/create/field?form=${Uuid}`,
        JSON.stringify(
            {
                quiz:Field
            }),{headers:{
                Authorization:`Bearer ${localStorage.getItem('access')}`
            }}
    )
    GetUserForms()

}

export default CreateField