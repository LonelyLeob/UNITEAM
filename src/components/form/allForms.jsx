import {useEffect, useState} from "react";
import axios from "axios";
import Form from "./Form"
import FalseForm from "../falseForm/falseForm";
import "./Form.css"
import Header from "../header/header";

function AllForms(){

    const [forms, setForms] = useState([])
    const [state, setState] = useState(false)

    // let addForm = (data) => {
    //     setForms(prev => [...prev, data])
    // }

    let errText = "Создать форму"
    let notErrText = "Добавить форму"


    useEffect(() => {
        axios
            .get("http://uni-team-inc.online:8080/api/v1/get/forms",
                {headers:{
                    Authorization:`Bearer ${localStorage.getItem('access')}`
                }})
            .then(data => {
                setForms(data.data)
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
            <div className="allFormContainer">
                {forms.map((item, idx) => {
                    return(
                        <div key={idx}>
                            <Form item = {item}/>
                        </div>
                    )})}
            </div>
        </div>
    )
}

export default AllForms

