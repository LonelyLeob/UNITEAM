import Answer from "../answer/Answers";
import "./Field.css"

function Field(props) {
    return (
        <div className="fieldsContainer">
        {props.fields.map((item, idx) => {
            return (
                <div key={idx}>
                    <p className="fieldsInput">Вопрос: {item.Quiz}</p>
                    <div className="fieldsAnsw">
                        <Answer answers={item.Answers}/>
                        <button className="fieldsBtn">+</button>
                    </div>

                </div>
        )
        })}
        </div>
    );
}

export default Field;
