import Answer from "../answer/Answers";
import "./Field.css"
import axios from "axios";

function Field(props) {


    let id = props.fieldsId

    let handleSubmit = async () => {
        let response = await axios.post(`http://uni-team-inc.online:8080/api/v1/create/answer?field=${id}`,
            JSON.stringify(
                {
                    answer:"Ответ"
                }),{headers:{
                    Authorization:`Bearer ${localStorage.getItem('access')}`
                }}
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
