import Answer from "../answer/Answers";

function Field(props) {
    return (
        <div>
        {props.fields.map((item, idx) => {
            return (
                <div key={idx}>
                    <p className="pb-5 font-semibold">{item.Quiz}</p>
                    <Answer answers={item.Answers}/>
                    <button className="mt-5 ml-5 p-1.5 border-solid border border-cyan-400 rounded-lg bg-slate-100 hover:bg-slate-300">Отправить</button>
                </div>
        )
        })}
        </div>
    );
}

export default Field;
