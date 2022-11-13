import "./addFormStyle.css"
import {useState, useEffect} from "react";
import axios from "axios";

function AddForm(){

    const[forms, setForms] = useState([])
    const[formName,setFormName] = useState('')
    const[formDescription,setFormDescription] = useState('')
    const[formAnon,setFormAnon] = useState(true)


    useEffect(() => {
        axios
            .get("http://localhost:8080/get/forms", {withCredentials: true})
            .then(data => {
                setForms(data.data)
            }).catch(err => {
            console.log(err)
        })
    }, [])


    let handleSubmit = async (e) => {
        e.preventDefault();
        try {
            let response = await axios.post("http://localhost:8080/create",
                JSON.stringify({
                    name: formName,
                    desc: formDescription,
                    anon: formAnon
                }),
                {withCredentials: true}
            ).then(data => {
                forms.unshift(data.data)
            })
        }   catch (err) {
            console.log("u vas err")
        }
    }

    return(
        <div className="AddForm">

            <form method="POST" className="addUserForm" >
                <h1 className=""> <b>Создание формы</b> </h1>


                <label>Название формы:</label>
                <br/>
                <input value={formName} onChange={event => setFormName(event.target.value)} name="formName" type="text" required/>
                <br/>

                <label>Описание:</label>
                <br/>
                <input value={formDescription} onChange={event => setFormDescription(event.target.value)} name="formDescription" type="text" required/>
                <br/>

                <div className="anon">
                    <label>Анонимная форма?</label>
                    <br/>
                    <input  type="checkbox" className="anonInp" checked={formAnon} onChange={(formAnon) => setFormAnon(!formAnon)} name="formAnon"/>
                    <br/>
                </div>

                <button className="" onClick={(e) => handleSubmit(e) } type="submit">Создать</button>
            </form>
        </div>
    )
}
export default AddForm