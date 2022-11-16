import "./Form.css"
import axios from "axios";
import AllFields from "../field/allFields"

function Form(props) {

    let uuid = props.item.Uuid

    let handleSubmit = async () => {
        let res = await axios.post(`http://uni-team-inc.online:8080/api/v1/create/field?form=${uuid}`,
            JSON.stringify(
                {
                    quiz: "абщхда"
                }),{headers:{
                    Authorization:`Bearer ${localStorage.getItem('access')}`
                }}
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

