import axios from "axios";

function DeleteField(Id){

    let res = axios.delete(`http://uni-team-inc.online:8080/api/v1/delete/field?id=${Id}`)

}

export default DeleteField

