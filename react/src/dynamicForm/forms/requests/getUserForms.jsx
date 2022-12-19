import axios from "axios";

async function GetUserForms(setForm){
         await axios.get("http://uni-team-inc.online:8080/api/v1/get/short",
            {headers:{
                    Authorization:`Bearer ${localStorage.getItem('access')}`
                }}
        )
            .then(data => {
                setForm(data.data)
            })
}

export default GetUserForms