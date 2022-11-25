import "./changeStyle.css"
import axios from "axios"

function ChangeAnswers(props){

    let handleDelete = async (id) => {
        let result = window.confirm("Вы уверенны?");
            if (result === true){
                let res = await axios.delete(`http://localhost:8080/api/v1/delete/answer?id=${id}`)
            }
}

    return(
        <>
            {props.answers.map((item, idx) => {
                return(
                    <>
                    <div key={idx} className="changeAnswers">
                        <p>{item.Answer}</p>
                        <button className="answerDelBtn" onClick={() => {handleDelete(item.id)}}>X</button>
                    </div><br/>
                    </>
                );
            })}
        </>
    )
}
export default ChangeAnswers