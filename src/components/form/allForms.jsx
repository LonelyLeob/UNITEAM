import {useEffect, useState} from "react";
import axios from "axios";
import Form from "./Form"
import "./Forms.css"

function AllForms(){

    const [forms, setForms] = useState([])

    let handleSubmit = async (e) => {
        e.preventDefault();
        try {
        let response = await axios.post("http://localhost:8080/create", 
            JSON.stringify({
            name: "1234",
            desc: "3214"
        }), 
        {withCredentials: true}
        ).then(data => {
            forms.unshift(data.data)
        })
    }   catch (err) {
        console.log("u vas err")
    }
}


    useEffect(() => {
        axios
            .get("http://localhost:8080/get/forms", {withCredentials: true})
            .then(data => {
                setForms(data.data)
            }).catch(err => {
            console.log(err)
        })
    }, [])

    if (forms == null) {
        return (
            <div className="error">Нет форм</div>
        ) 
    }

    return (
        <div>
            <button onClick={(e) => handleSubmit(e) } className="addForm">+</button>
            {forms.map((item) => {
                return(
                    <Form item = {item}/>
                )
            })}
        </div>
    )
}

export default AllForms

