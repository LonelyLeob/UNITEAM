import "./edit.css"
import FieldsAnswerEdit from "./fieldsAnswerEdit";
import DeleteField from "./requests/deleteField";

function FieldsEdit(props){

    return(
        <>
            {props.fields.map((item, idx) => {
                return(
                    <div className="changeFields" key={idx}>
                        <div className="wrapper">
                            <p>{item.Quiz}</p>
                            <button className="fieldsDelBtn" onClick={async(e) => {
                                    e.preventDefault()
                                    await DeleteField(item.Id)
                                    props.setCount((prev) => prev + 1)
                            }}>X</button>
                        </div>
                        <br/>
                        <FieldsAnswerEdit fieldsId={item.Id} fieldsAnswers={item.Answers}/>
                    </div>
                )
            })}
        </>
    )
}
export default FieldsEdit

