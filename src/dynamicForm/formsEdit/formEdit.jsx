import "./edit.css"
import Header from "../../header/header";
import FieldsEdit from "./fieldsEdit";
import {useState} from "react";
import DeleteForm from "./requests/deleteForm";
import CreateField from "./requests/createField";

function FormEdit(){

    const[field, setField] = useState('')
    const data = JSON.parse(localStorage.getItem("data"));

    const createField = async (e) => {
        e.preventDefault()
        await CreateField(data.Uuid, field)
    }

    return(
                <>
                    <Header/>
                        <div className="mainChangeContainer">
                            <div className="btnDelContainer">
                                <label htmlFor="">
                                    Удалить форму
                                    <button className="btnDel" onClick={() => DeleteForm(data.Uuid)}>-</button>
                                </label>
                            </div>
                            <div className="changeContainer">
                                <div className="changeContainerContent">
                                    <h1>{data.Name}</h1>
                                    <p> {data.Description}</p><br/>
                                    <div className="wrapperBtn">
                                        <input type="text" placeholder="Добавить поле" maxLength="50" value={field} onChange={event => setField(event.target.value)} autoComplete="on"/>
                                        <button className="btn" onClick={(e) => {createField(e)}}>+</button>
                                    </div>
                                </div><br/>
                                {data.Fields && <FieldsEdit fields={data.Fields} uuid={data.Uuid}/>}
                            </div>
                        </div>
                </>
    )
}
export default FormEdit

