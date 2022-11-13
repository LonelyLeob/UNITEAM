import {useEffect, useState} from "react";
import axios from "axios";
import Form from "./Form"
import FalseForm from "../falseForm/falseForm";
import "./Form.css"

function AllForms(){

    const [forms, setForms] = useState([])
    let errText = "Создать форму"
    let notErrText = "Добавить форму"

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

            <div className="error">
                <FalseForm text = {errText}/>
                <p className="HasnotForm">У вас нет форм</p>
            </div>
        )
    }

    return(
        <div>
            <FalseForm text = {notErrText}/>
            {forms.map((item, idx) => {
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

