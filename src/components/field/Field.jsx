import Answer from "../answer/Answers";
import "./Field.css"
import axios from "axios";

function Field(props) {


    let Id = props.fieldsId

    let handleSubmit = async () => {
        let response = await axios.post("http://localhost:8080/create/field/answer?field=" + Id,
            JSON.stringify(
                {
                    answer:"Ответ"
                }),
            {withCredentials: true}
        )}


    if (props.fieldsAnswers !== null) {
                return (
                    <div className="fieldsContainer">
                        <div className="fieldsAnsw">
                            <Answer answers={props.fieldsAnswers}/>
                            <button className="fieldsBtn" onClick={(e) => {handleSubmit(e)}}>+</button>
                        </div>
                    </div>
                )}

                return (
                    <div className="fieldsContainer">
                        <div className="fieldsAnsw">
                            <button className="fieldsBtn" onClick={(e) => {handleSubmit(e)}}>+</button>
                        </div>
                    </div>
                )
}

export default Field;
