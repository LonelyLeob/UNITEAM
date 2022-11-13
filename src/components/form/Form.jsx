import "./Form.css"
import axios from "axios";
import AllFields from "../field/allFields"

function Form(props) {

    let Uuid = props.item.Uuid

    let handleSubmit = async () => {
        let res = await axios.post("http://localhost:8080/create/field?form=" + Uuid,
            JSON.stringify(
                {
                    quiz: "абщхда"
                }),
            {withCredentials: true}
    )}


        if (props.item.Fields !== null) {
            return (
                <div className="formContainer">
                    <div className="formContainerField">
                        <h1 className="formH">{props.item.Name}</h1>
                        <button className="fieldsBtn" onClick={(e) => {handleSubmit(e)}}>+</button>
                    </div>
                    <p className="formP">{props.item.Description}</p>
                    <AllFields fields={props.item.Fields}/>
                </div>
            );
        }

        return (
                <div key={props.key} className="formContainer">
                    <div className="formContainerField">
                        <h1 className="formH">{props.item.Name}</h1>
                        <button className="fieldsBtn" onClick={(e) => {handleSubmit(e)}}>+</button>
                    </div>
                    <p className="formP">{props.item.Description}</p>
                </div>

        );

}

export default Form;

