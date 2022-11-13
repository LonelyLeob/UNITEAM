import Field from "./Field";
import "./Field.css"

function AllFields(props){

    return(
            <div className="fieldsContainer">
                {props.fields.map((item, idx) => {
                        return (
                            <div key={idx}>
                                <p className="fieldsInput">Вопрос: {item.Quiz}</p>
                                    <Field fieldsId={item.Id} fieldsAnswers={item.Answers} />
                            </div>
                        )
                })}
            </div>
    )
}
export default AllFields