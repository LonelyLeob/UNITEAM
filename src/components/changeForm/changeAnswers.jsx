import "./changeStyle.css"

function ChangeAnswers(props){
    return(
        <>
            {props.answers.map((item, idx) => {
                return(
                    <>
                    <div key={idx} className="changeAnswers">
                        <p className="queryAnsw">Ответ {idx +1}</p>
                        <p>{item.Answer}</p>
                    </div><br/>
                    </>
                );
            })}
        </>
    )
}
export default ChangeAnswers