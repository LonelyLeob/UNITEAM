import axios from "axios";

async function GetForm(setForm,Uuid ){

         await axios.get(`http://uni-team-inc.online:8080/api/v1/get/form?uid=${Uuid}`,
            {headers:{
                    Authorization:`Bearer ${localStorage.getItem('access')}`
                }}
        )
            .then(data => {
                setForm(data.data)
            })
}

export default GetForm