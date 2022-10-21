function Answer(props) {
    return (
        <div>
            {props.answers.map((item, idx) => {
                return (
                    <div key={idx} className='flex p-5 border-solid border-b border-cyan-400 items-center space-x-60 '>
                        <p className="min-w-[24%]">{item.Answer}</p>
                        <p className='p-3 border-solid border-2 border-cyan-400 rounded-full '></p>
                    </div>
                )
            })}
        </div>
    );}

export default Answer;
