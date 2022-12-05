import axios from "axios";

function DeleteUser(name, pass){

     axios.delete(`http://uni-team-inc.online:4000/api/v1/delete?n=${name}&p=${pass}`)
     
}

export default DeleteUser