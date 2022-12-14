import axios from "axios";

function DeleteAnswer(Id){

    let res = axios.delete(`http://uni-team-inc.online:8080/api/v1/delete/answer?id=${Id}`)

}

export default DeleteAnswer

