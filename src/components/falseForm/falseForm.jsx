import "./falseFormStyle.css"


function FalseForm(props){
    return(
      <div className="FasleForm">
          <p>{props.text}</p>
          <button>+</button>
      </div>
    )
}
export default FalseForm