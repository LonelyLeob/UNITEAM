import "./edit.css"
import DeleteAnswer from "./requests/deleteAnswer";

function AnswersEdit(props){

    return(
        <>
            {props.answers.map((item, idx) => {
                return(
                    <div key={idx} >
                        <div className="changeAnswers">
                            <p>{item.Answer}</p>
                            <button className="answerDelBtn" onClick={() => DeleteAnswer(item.Id)}>X</button>
                        </div><br/>
                    </div>
                )
            })}
        </>
    )
}
export default AnswersEdit

