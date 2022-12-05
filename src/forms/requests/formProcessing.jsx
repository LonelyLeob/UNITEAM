import "../formStyle.css"
import {useEffect, useState} from "react";
import GetUserForms from "./getUserForms";
import ViewForms from "../form";
import Header from "../../header/header";
import Modal from "../modal/modalWin";
import AddNewForm from "./addNewForm";

function FormProcessing(){

    const [forms, setForms] = useState([])
    const [formsState, setFormsState] = useState()
    const [isModal, setModal] = useState(false);
    const[formName,setFormName] = useState('')
    const[formDescription,setFormDescription] = useState('')
    const[formAnon,setFormAnon] = useState(true)
    let i


    useEffect(() => {
        GetUserForms(setForms)
    }, [formsState])

    let content =
        <form action="" className="modalForm">
            <input placeholder="Название" value={formName} onChange={event => setFormName(event.target.value)} className="modalFormName" type="text"/>
            <input placeholder="Описание" value={formDescription} onChange={event => setFormDescription(event.target.value)} className="modalFormDesc" type="text"/>
            <label htmlFor="">Сделать форму анонимной?<input checked={formAnon} onChange={(event) => setFormAnon(event.target.checked)} className="modalFormCheck" type="checkbox"/></label><br/>
            <button type="submit" className="modalBtn" onClick={(e) =>  formHandler(e)}>Отправить</button>
        </form>

    let formHandler = async(e) => {
        e.preventDefault()
       await AddNewForm(formName, formDescription, formAnon)
        setFormsState(i + 1)
        setModal(false)
    }

    return(
            <div className="allFormContainer">
                    <Header/>
                    <div className="addForm">
                        {forms ? <p>Добавить форму</p> : <p>Создать форму</p>}
                        <button onClick={() => setModal(true)}>+</button>
                    </div>

                    <Modal
                        isVisible={isModal}
                        title="Создание формы"
                        content={content}
                        footer={<p></p>}
                        onClose={() => setModal(false)}/>

                <div className="forms">
                    {forms !== null ? forms.map((item, idx) => {
                        return(
                            <div key={idx}>
                                <ViewForms item = {item}/>
                            </div>

                        )}) : <ViewForms err = {"У вас нет форм"}/>
                    }
                </div>
            </div>

    )
}

export default FormProcessing