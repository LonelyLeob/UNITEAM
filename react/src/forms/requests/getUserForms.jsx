import axios from "axios";

function GetUserForms(setForms){
           axios.get("http://uni-team-inc.online:8080/api/v1/get/forms",
            {headers:{
                    Authorization:`Bearer ${localStorage.getItem('access')}`
                }}
        )
            .then(data => {
                return setForms(data.data)
            })
}

export default GetUserForms