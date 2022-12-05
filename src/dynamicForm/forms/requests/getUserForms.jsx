import axios from "axios";

async function GetUserForms(){
         await axios.get("http://uni-team-inc.online:8080/api/v1/get/forms",
            {headers:{
                    Authorization:`Bearer ${localStorage.getItem('access')}`
                }}
        )
            .then(data => {
                localStorage.removeItem("forms")
                localStorage.setItem("forms", JSON.stringify(data.data))
            })
}

export default GetUserForms