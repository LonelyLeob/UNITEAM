import "./formStyle.css"
import {useNavigate} from "react-router-dom";
// import {useEffect, useState} from "react";

function ViewForms(props){

    let navigate = useNavigate()

    let handlerSubmit = () => {
        localStorage.setItem("data", JSON.stringify(props.item))
        navigate('/edit')
    }

    return(
        <div>
            {props.item ?
                    <div className="formContainer" onClick={(e) => {handlerSubmit(e)}} >
                            <div className="formContainerField">
                                <h1 className="formH">{props.item.Name}</h1>
                            </div>
                            <p className="formP">{props.item.Description}</p>
                    </div>
                : <h1 className="errText">{props.err}</h1> }
        </div>
    )
}
export default ViewForms

