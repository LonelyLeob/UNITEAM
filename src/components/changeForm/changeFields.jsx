import "./changeStyle.css"
import ChangeFieldsAnsw from "./changeFieldsAnsw";
import axios from "axios"

function ChangeFields(props){

    let handleDelete = async (id) => {
        let result = window.confirm("Вы уверенны?");
            if (result === true){
                let res = await axios.delete(`http://uni-team-inc.online:8080/api/v1/delete/field?id=${id}`)
            }
}

    return(
            <>
                {props.fields.map((item, idx) => {
                    return(
                        <div className="changeFields" key={idx}>
                                <div className="wrapper">
                                    <p>{item.Quiz}</p>
                                    <button className="fieldsDelBtn" onClick={() => {handleDelete(item.Id)}}>X</button>
                                </div>
                            <br/>
                            <ChangeFieldsAnsw fieldsId={item.Id} fieldsAnswers={item.Answers}/>
                        </div>
                    )
                })}
            </>
    )
}
export default ChangeFields