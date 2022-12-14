import "./edit.css"
import AnswersEdit from "./answersEdit";
import CreateAnswer from "./requests/createAnswer";
import {useState} from "react";

function FieldsAnswerEdit(props){

    const [answer, setAnswer] = useState('')

        return(
            <>
                <div className="wrapperAnsw">
                    {props.fieldsAnswers !== null ? <AnswersEdit answers={props.fieldsAnswers} id={props.fieldsId}/> : <p>Нет ответов</p>}
                    <div className="wrapperBtn">
                        <input type="text" placeholder="Добавить ответ" maxLength="50" value={answer} onChange={event => setAnswer(event.target.value)} autoComplete="on"/>
                        <button className="btn" onClick={() => CreateAnswer(answer, props.fieldsId)}>+</button>
                    </div>
                </div><br/>
            </>
        )

}
export default FieldsAnswerEdit

