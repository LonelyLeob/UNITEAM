import "./Form.css"
import {Link} from "react-router-dom";

function Form(props) {

            return (
                <div className="formContainer">
                    <div className="formContainerField">
                        <Link to="Change" onClick={(e) => {localStorage.setItem("data", JSON.stringify(props.item))}}>
                            <h1 className="formH">{props.item.Name}</h1>
                        </Link>
                    </div>
                    <p className="formP">{props.item.Description}</p>
                </div>
            )
}

export default Form;

