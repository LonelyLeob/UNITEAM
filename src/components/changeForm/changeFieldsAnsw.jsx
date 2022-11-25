import "./changeStyle.css"
import ChangeAnswers from "./changeAnswers";
import axios from "axios";
import {useState} from "react";

function ChangeFieldsAnsw(props){

    const [answer, setAnswer] = useState('')

    let id = props.fieldsId

    let handleSubmit = async () => {
        let response = await axios.post(`http://uni-team-inc.online:8080/api/v1/create/answer?field=${id}`,
            JSON.stringify(
                {
                    answer:answer
                }),{headers:{
                    Authorization:`Bearer ${localStorage.getItem('access')}`
                }}
        )}

    if (props.fieldsAnswers !== null) {
        return(
            <>
            <div className="wrapperAnsw">
                <ChangeAnswers answers={props.fieldsAnswers} id={id}/>
                <div className="wrapperBtn">
                    <input type="text" placeholder="Добавить ответ" maxLength="50" value={answer} onChange={event => setAnswer(event.target.value)}/>
                <button className="btn" onClick={(e) => {handleSubmit(e)}}>+</button>
                </div>
            </div><br/>
            </>
        )
    }

    return(
        <>
            <p>Нет ответов</p><br/>
            <div className="wrapperBtn">
                <input type="text" placeholder="Добавить ответ" maxLength="50" value={answer} onChange={event => setAnswer(event.target.value)}/>
                <button className="btn" onClick={(e) => {handleSubmit(e)}}>+</button>
            </div>
        </>
    )
}
export default ChangeFieldsAnsw