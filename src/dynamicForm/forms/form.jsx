import "./formStyle.css"
import {Link} from "react-router-dom";

function ViewForms(props){

    return(
        <div>
            {props.item ?
                <Link to={`/edit/${props.item.Uuid}`} className="courseTitleLink">
                    <div className="formContainer">
                            <div className="formContainerField">
                                <h1 className="formH">{props.item.Name}</h1>
                            </div>
                            <p className="formP">{props.item.Description}</p>
                    </div>
                    </Link>
                : <h1 className="errText">{props.err}</h1> }
        </div>
    )
}
export default ViewForms

