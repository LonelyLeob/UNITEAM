import {useState} from "react";
import axios from "axios";
import "./modalStyle.css"


function ModalAdd(props){

    const[formName,setFormName] = useState('')
    const[formDescription,setFormDescription] = useState('')
    const[formAnon,setFormAnon] = useState(true)
    const[error, setError] = useState('')

    let handleSubmit = async (e) => {
        e.preventDefault();

        if (formName.trim().length === 0 || formDescription.trim().length === 0){
            setError('Введены неверные данные')
            return}

        if (formName.trim().length === 0 && formDescription.trim().length === 0){
            setError('Введены неверные данные')
            return}

            try{
                    let response = await axios.post("http://uni-team-inc.online:8080/api/v1/create",
                        JSON.stringify({
                            name: formName,
                            desc: formDescription,
                            anon: formAnon
                        }),{headers:{
                                Authorization:`Bearer ${localStorage.getItem('access')}`
                            }})
                // props.setNewForm(response.data)

            }
            catch{
             setError('Введены неверные данные')
            }
            props.setModal(false)
    }


return(
    <>
        <form action="" className="modalForm">
           <input placeholder="Название" value={formName} onChange={event => setFormName(event.target.value)} className="modalFormName" type="text"/>
            <input placeholder="Описание" value={formDescription} onChange={event => setFormDescription(event.target.value)} className="modalFormDesc" type="text"/>
            <label htmlFor="">Сделать форму анонимной?<input checked={formAnon} onChange={(event) => setFormAnon(event.target.checked)} className="modalFormCheck" type="checkbox"/></label><br/>
            <button type="submit" className="modalBtn" onClick={(e) => handleSubmit(e)}>Отправить</button>
            {error && <p className="validError">{error}</p>}
        </form>
    </>
)
}

export default ModalAdd