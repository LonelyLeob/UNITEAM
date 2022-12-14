import axios from "axios";

function GetUser(setUser){

     axios.get("http://uni-team-inc.online:4000/api/v1/user",
        {headers:{
                Authorization:`Bearer ${localStorage.getItem('access')}`
            }}
    )
        .then(data => {
            setUser(data.data)
        })
}

export default GetUser