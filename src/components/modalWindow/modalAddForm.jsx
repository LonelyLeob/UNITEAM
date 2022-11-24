import {useEffect, useState} from "react";
import axios from "axios";
import "./modalStyle.css"

function ModalAdd(){

    const[forms, setForms] = useState([])
    const[formName,setFormName] = useState('')
    const[formDescription,setFormDescription] = useState('')
    const[formAnon,setFormAnon] = useState(true)

    useEffect(() => {
        axios
            .get("http://uni-team-inc.online:8080/api/v1/get/forms",
                {headers:{
                        Authorization:`Bearer ${localStorage.getItem('access')}`
                    }}
            )
            .then(data => {
                setForms(data.data)
            }).catch(err => {
            console.log(err)
        })
    }, [])


    let handleSubmit = async (e) => {
        e.preventDefault();
        try {
            let response = await axios.post("http://uni-team-inc.online:8080/api/v1/create",
                JSON.stringify({
                    name: formName,
                    desc: formDescription,
                    anon: formAnon
                }),{headers:{
                        Authorization:`Bearer ${localStorage.getItem('access')}`
                    }}
            ).then(data => {
                forms.unshift(data.data)
            })
        }   catch (err) {
            console.log("u vas err")
        }
    }


return(
    <>
        <form action="" className="modalForm">
           <input placeholder="Название" value={formName} onChange={event => setFormName(event.target.value)} className="modalFormName" type="text"/>
            <input placeholder="Описание" value={formDescription} onChange={event => setFormDescription(event.target.value)} className="modalFormDesc" type="text"/>
            <label htmlFor="">Сделать форму анонимной?<input checked={formAnon} onChange={(event) => setFormAnon(event.target.checked)} className="modalFormCheck" type="checkbox"/></label><br/>
            <button type="submit" className="modalBtn" onClick={(e) => handleSubmit(e)}>Отправить</button>
        </form>
    </>
)
}

export default ModalAdd