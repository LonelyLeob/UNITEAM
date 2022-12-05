import axios from "axios";

function DeleteForm(Uuid){

    let res = axios.delete(`http://uni-team-inc.online:8080/api/v1/delete?form=${Uuid}`)

}

export default DeleteForm