import {useEffect, useState} from "react";
import axios from "axios";
import Form from "./Form"


function AllForms(){

    const [data, setData] = useState([])


    useEffect(() => {
        axios
            .get("http://localhost:8080/get/forms", {withCredentials: true})
            .then(data => {
                setData(data.data)
            }).catch(err => {
            console.log(err)
        })
    }, [])

    return(
        <div>
            {data.map((item, idx) => {
                return(
                    <div key={idx}>
                        <Form item = {item}/>
                    </div>
                )
            })}
        </div>
    )
}

export default AllForms

