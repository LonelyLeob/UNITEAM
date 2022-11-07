import "./Answer.css"

function Answer(props) {

    return(
        <>
            {props.answers.map((item, idx) => {
                return(
                    <div key={idx} className='answerContainer'>
                        <p className="textAnsw">{item.Answer}</p>
                    </div>
                );
            })}
        </>
    )


   }

export default Answer;
