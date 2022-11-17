import "./falseFormStyle.css"
import ModalWin from "../modalWindow/modalWin";

function FalseForm(props){
    return(
      <div className="FasleForm">
          <p>{props.text}</p>
          <ModalWin/>
      </div>
    )
}
export default FalseForm