import "./changeStyle.css"

function ChangeAnswers(props){
    return(
        <>
            {props.answers.map((item, idx) => {
                return(
                    <>
                    <div key={idx} className="changeAnswers">
                        <p>{item.Answer}</p>
                    </div><br/>
                    </>
                );
            })}
        </>
    )
}
export default ChangeAnswers