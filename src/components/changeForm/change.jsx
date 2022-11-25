import "./changeStyle.css"
import ChangeFields from "./changeFields";
import axios from "axios";
import {useState} from "react";


function ChangeForm() {

    const[field, setField] = useState('')

    // if (data != null) {
    //     localStorage.removeItem("data")}

    const data = JSON.parse(localStorage.getItem("data"));

    let uuid = data.Uuid

    let handleSubmit = async () => {
        let res = await axios.post(`http://uni-team-inc.online:8080/api/v1/create/field?form=${uuid}`,
            JSON.stringify(
                {
                    quiz:field
                }),{headers:{
                    Authorization:`Bearer ${localStorage.getItem('access')}`
                }}
        )}

        let handleDelete = async () => {
            let result = window.confirm("Вы уверенны?");
                if (result === true){
                    let res = await axios.delete(`http://uni-team-inc.online:8080/api/v1/delete?form=${uuid}`)
                }
    }

    if (data.Fields !== null) {
        return (
            <div className="mainChangeContainer">
                   <div className="btnDelContainer">
                        <label htmlFor="">
                            Удалить форму
                        <button className="btnDel" onClick={(e) => {handleDelete(e)}}>-</button>
                        </label>
                    </div>
                <div className="changeContainer">
                        <div className="changeContainerContent">
                            <h1>{data.Name}</h1>
                            <p> {data.Description}</p><br/>
                            <div className="wrapperBtn">
                                <input type="text" placeholder="Добавить поле" maxLength="50" value={field} onChange={event => setField(event.target.value)}/>
                                <button className="btn" onClick={(e) => {handleSubmit(e)}}>+</button>
                            </div>
                        </div><br/>
                    <ChangeFields fields={data.Fields} uuid={uuid}/>
                </div>
            </div>
        )}

    return (
        <div className="mainChangeContainer">
            <div className="btnDelContainer">
                <label htmlFor="">
                    Удалить форму
                    <button className="btnDel" onClick={(e) => {handleDelete(e)}}>-</button>
                </label>
            </div>
            <div className="changeContainer">
                <div className="changeContainerContent">
                    <h1>{data.Name}</h1>
                    <p> {data.Description}</p><br/>
                    <div className="wrapperBtn">
                        <input type="text" placeholder="Добавить поле" maxLength="50" value={field} onChange={event => setField(event.target.value)}/>
                        <button className="btn" onClick={(e) => {handleSubmit(e)}}>+</button>
                    </div>
                </div><br/>
            </div>
        </div>
    )}

export default ChangeForm