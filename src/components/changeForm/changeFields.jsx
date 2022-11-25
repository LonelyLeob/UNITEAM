import "./changeStyle.css"
import ChangeFieldsAnsw from "./changeFieldsAnsw";
import axios from "axios"

function ChangeFields(props){

    let handleDelete = async () => {
        let result = window.confirm("Вы уверенны?");
            if (result === true){
                let res = await axios.delete(`http://localhost:8080/api/v1/delete/field?id=${props.uuid}`)
            }
}

    return(
            <>
                {props.fields.map((item, idx) => {
                    return(
                        <div className="changeFields" key={idx}>
                                <div className="wrapper">
                                    <p>{item.Quiz}</p>
                                    <button className="fieldsDelBtn" onClick={() => {handleDelete()}}>X</button>
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