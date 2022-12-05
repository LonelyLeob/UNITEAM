import "./formStyle.css"
import {useNavigate} from "react-router-dom";


function ViewForms(props){


    let navigate = useNavigate()
    let handlerSubmit = () => {
        JSON.stingyfy(localStorage.setItem("data", props.item))
        navigate('/edit')
    }
   
  
    return(
                <div className="formContainer" onClick={(e) => {handlerSubmit(e)}} >
                    {props.item ?
                        <>
                            <div className="formContainerField">
                                <h1 className="formH">{props.item.Name}</h1>
                            </div>
                            <p className="formP">{props.item.Description}</p>
                        </>
                        :    <h1>{props.err}</h1>
                    }
                </div>
    )
}
export default ViewForms

