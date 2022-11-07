// import {useEffect, useState} from "react";
// import axios from "axios";
import Form from "./Form"
import data from "../../jj.json"


function AllForms(){

    // const [data, setData] = useState([])


    // useEffect(() => {
    //     axios
    //         .get("http://localhost:8080/get/forms", {withCredentials: true})
    //         .then(data => {
    //             setData(data.data)
    //             console.log(data)
    //         }).catch(err => {
    //         console.log(err)
    //     })
    // }, [])

    return (
        <div>
            {data.map((item, idx) => {return (<Form key={idx} item={item}/>)})}
        </div>
    )
}

export default AllForms

