import "./changeStyle.css"
import ChangeFieldsAnsw from "./changeFieldsAnsw";

function ChangeFields(props){


    return(
            <>
                {props.fields.map((item, idx) => {
                    return(
                        <div className="changeFields" key={idx}>
                                <div className="wrapper">
                                    <p className="queryAnsw">Поле {idx + 1}</p>
                                    <p>{item.Quiz}</p>
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