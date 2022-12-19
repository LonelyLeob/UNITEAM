import "./edit.css"
import Header from "../../header/header";
import FieldsEdit from "./fieldsEdit";
import {useState} from "react";
import DeleteForm from "./requests/deleteForm";
import CreateField from "./requests/createField";
import {useParams, useNavigate} from "react-router-dom";
import { useEffect } from "react";
import GetForm from "./requests/getForm";

function FormEdit(){

    const[field, setField] = useState('')
    const[form, setForm] = useState([])
    const[count, setCount] = useState(0)
    const navigate = useNavigate()
    const params = useParams()
    const formUuid = params.uuid


    const createField = async (e) => {
        e.preventDefault()
        await CreateField(formUuid, field)
        setCount((prev) => prev + 1)
        setField('')
    }

    useEffect(() => {
        req()
    }, [count])

    const req = async() => {
        await GetForm(setForm, formUuid)
    }

    return(
                <>
                    <Header/>
                        <div className="mainChangeContainer">
                            <div className="btnDelContainer">
                                <label htmlFor="">
                                    Удалить форму
                                    <button className="btnDel" onClick={async(e) => {
                                        e.preventDefault()
                                        await DeleteForm(formUuid) 
                                        navigate(-1)
                                        }}>-</button>
                                </label>
                            </div>
                            <div className="changeContainer">
                                <div className="changeContainerContent">
                                    <h1>{form.Name}</h1>
                                    <p> {form.Description}</p><br/>
                                    <div className="wrapperBtn">
                                        <input type="text" placeholder="Добавить поле" maxLength="50" value={field} onChange={event => setField(event.target.value)} autoComplete="on"/>
                                        <button className="btn" onClick={(e) => {createField(e)}}>+</button>
                                    </div>
                                </div><br/>
                                {form.Fields && <FieldsEdit fields={form.Fields} uuid={form.Uuid} setCount={setCount} />}
                            </div>
                        </div>
                </>
    )
}
export default FormEdit

